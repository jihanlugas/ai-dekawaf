package websocket

import (
	"ai-dekawaf/wsredis"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/websocket"
)

type WebsocketHandler struct {
	wsSub *wsredis.Subscriber
	// wsPub        wsredis.Publisher
	wsAdminSub *wsredis.Subscriber
	// wsAdminPub   wsredis.Publisher
}

type DefaultListenData struct {
	ProjectID string `json:"project_id"`
}

type DefaultMessageData struct {
	ProjectID string      `json:"project_id"`
	Data      interface{} `json:"data"`
}

type AdminListenData struct{}

type AdminMessageData struct {
	Data interface{} `json:"data"`
}

var (
	wsRedisPool    *redis.Pool
	defaultChannel string

	adminChannel string
)

func InitHandler() *WebsocketHandler {
	newHandler := WebsocketHandler{}
	// Init Redis Websocket
	wsRedisAddr := fmt.Sprintf("%s:%s", os.Getenv("WS_REDIS_HOST"), os.Getenv("WS_REDIS_PORT"))
	wsRedisPool = wsredis.NewPool(wsRedisAddr, os.Getenv("WS_REDIS_PASSWORD"))
	defaultChannel = os.Getenv("WS_REDIS_CHANNEL")

	// Initialize default channel
	// wsPub = wsredis.NewPublisher(wsRedisPool, defaultChannel)
	newHandler.wsSub = wsredis.NewSubscriber(wsRedisPool, defaultChannel, func(conn *wsredis.ConnectionData, msg []byte) error {
		if conn == nil {
			return errors.New("no socket connection found")
		}

		data := conn.Data.(DefaultListenData)
		var message DefaultMessageData
		if err := json.Unmarshal(msg, &message); err != nil {
			return err
		}

		if data.ProjectID == message.ProjectID {
			if err := conn.Socket.WriteJSON(message.Data); err != nil {
				return err
			}
		}

		return nil
	})

	// Initialize admin channel
	adminChannel = os.Getenv("WS_REDIS_CHANNEL") + "-admin"
	// wsAdminPub = wsredis.NewPublisher(wsRedisPool, adminChannel)
	newHandler.wsAdminSub = wsredis.NewSubscriber(wsRedisPool, adminChannel, func(conn *wsredis.ConnectionData, msg []byte) error {
		if conn == nil {
			return errors.New("no socket connection found")
		}

		var message AdminMessageData
		if err := json.Unmarshal(msg, &message); err != nil {
			return err
		}

		if err := conn.Socket.WriteJSON(message.Data); err != nil {
			return err
		}

		return nil
	})

	go func() {
		for {
			err := newHandler.wsSub.Run()
			if err == nil {
				break
			}
			log.Println(err)
		}
	}()

	go func() {
		for {
			err := newHandler.wsAdminSub.Run()
			if err == nil {
				break
			}
			log.Println(err)
		}
	}()

	return &newHandler
}

func (h *WebsocketHandler) WsHandler(c *gin.Context) {
	var upgrader = websocket.Upgrader{}

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("unable to upgrade to websockets"))
		return
	}

	projectID := c.Query("project_id")
	registeredData := wsredis.ConnectionData{
		Socket: socket,
		Data: DefaultListenData{
			ProjectID: projectID,
		},
	}

	h.wsSub.Register(&registeredData)

	for {
		msgType, _, err := socket.ReadMessage()
		if err != nil {
			// if websocket.IsCloseError(err, websocket.CloseGoingAway) || err == io.EOF {
			// 	log.Println("Websocket closed!")
			// 	break
			// }
			log.Println("Websocket closed!", err)
			break
			// log.Println("Error reading websocket message")
		}
		switch msgType {
		// If client able to send message, do here
		// case websocket.TextMessage:
		// 	msg, err := validateMessage(p)
		// 	if err != nil {
		// 		log.Println("Invalid Message")
		// 		break
		// 	}
		// 	wsPub.Publish(p)
		default:
			// log.Println("Unknown Message!")
		}
	}

	h.wsSub.Unregister(&registeredData)

	socket.WriteMessage(websocket.CloseMessage, []byte{})
}

func (h *WebsocketHandler) WsHandlerAdmin(c *gin.Context) {
	var upgrader = websocket.Upgrader{}

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("unable to upgrade to websockets"))
		return
	}

	registeredData := wsredis.ConnectionData{
		Socket: socket,
		Data:   nil,
	}

	h.wsAdminSub.Register(&registeredData)

	for {
		msgType, _, err := socket.ReadMessage()
		if err != nil {
			// if websocket.IsCloseError(err, websocket.CloseGoingAway) || err == io.EOF {
			// 	log.Println("Websocket closed!")
			// 	break
			// }
			log.Println("Websocket closed!", err)
			break
			// log.Println("Error reading websocket message")
		}
		switch msgType {
		// If client able to send message, do here
		// case websocket.TextMessage:
		// 	msg, err := validateMessage(p)
		// 	if err != nil {
		// 		log.Println("Invalid Message")
		// 		break
		// 	}
		// 	wsPub.Publish(p)
		default:
			// log.Println("Unknown Message!")
		}
	}

	h.wsAdminSub.Unregister(&registeredData)

	socket.WriteMessage(websocket.CloseMessage, []byte{})
}

func SendWebsocketMessage(projectID string, data interface{}) {
	if wsRedisPool == nil {
		log.Println("failed to broadcast websocket: wsRedisPool is not initialized")
		return
	}
	message := DefaultMessageData{
		ProjectID: projectID,
		Data:      data,
	}
	b, err := json.Marshal(message)
	if err != nil {
		log.Println("failed to broadcast websocket", err)
		return
	}

	if err := wsredis.PublishWithRetry(wsRedisPool, defaultChannel, b); err != nil {
		log.Println("failed to broadcast websocket", err)
	}

	dataJSON, _ := json.Marshal(data)
	log.Println("success broadcast websocket", projectID, string(dataJSON))
}

func SendWebsocketMessageAdmin(data interface{}) {
	if wsRedisPool == nil {
		log.Println("failed to broadcast websocket admin: wsRedisPool is not initialized")
		return
	}
	message := AdminMessageData{
		Data: data,
	}
	b, err := json.Marshal(message)
	if err != nil {
		log.Println("failed to broadcast websocket admin", err)
		return
	}

	if err := wsredis.PublishWithRetry(wsRedisPool, adminChannel, b); err != nil {
		log.Println("failed to broadcast websocket admin", err)
	}

	dataJSON, _ := json.Marshal(data)
	log.Println("success broadcast websocket admin", string(dataJSON))
}
