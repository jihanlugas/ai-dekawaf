package wsredis

import (
	"encoding/json"
	"log"

	// "time"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

// Message will be used in term websocket need to wait for initialization of redis adapter
// type message struct {
// 	Handle string `json:"handle"`
// 	Text   string `json:"text"`
// }

// ConnectionData is connected socket data type
type ConnectionData struct {
	Socket *websocket.Conn // socket connection
	Data   interface{}     // listening params (e.g. project_id)
}

// var (
// 	waitingMessage, availableMessage []byte
// 	waitSleep                        = time.Second * 10
// )

// func init() {
// 	var err error
// 	waitingMessage, err = json.Marshal(message{
// 		Handle: "system",
// 		Text:   "Waiting for redis to be available. Messaging won't work until redis is available",
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	availableMessage, err = json.Marshal(message{
// 		Handle: "system",
// 		Text:   "Redis is now available & messaging is now possible",
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// }

type MsgHandler func(data *ConnectionData, msg []byte) error

// Subscriber receives messages from Redis and broadcasts them to all
// registered websocket connections that are Registered.
type Subscriber struct {
	pool    *redis.Pool
	channel string

	messages       chan []byte
	newConnections chan *ConnectionData
	rmConnections  chan *ConnectionData

	MsgHandlers []MsgHandler
}

// NewSubscriber creates a receiver that will use the provided
// redis.Pool.
func NewSubscriber(pool *redis.Pool, channel string, msgHandlers ...MsgHandler) *Subscriber {
	return &Subscriber{
		pool:           pool,
		channel:        channel,
		messages:       make(chan []byte, 1000), // 1000 is arbitrary
		newConnections: make(chan *ConnectionData),
		rmConnections:  make(chan *ConnectionData),
		MsgHandlers:    msgHandlers,
	}
}

// func (r *Subscriber) Wait(_ time.Time) error {
// 	r.broadcast(waitingMessage)
// 	time.Sleep(waitSleep)
// 	return nil
// }

// Run receives pubsub messages from Redis after establishing a connection.
// When a valid message is received it is broadcast to all connected websockets
func (r *Subscriber) Run() error {
	conn := r.pool.Get()
	defer conn.Close()
	psc := redis.PubSubConn{Conn: conn}
	psc.Subscribe(r.channel)
	go r.connHandler()
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			var result interface{}
			if err := json.Unmarshal(v.Data, &result); err != nil {
				log.Println("Error unmarshalling message from redis")
				continue
			}
			r.broadcast(v.Data)
		case redis.Subscription:
			log.Println("New subscription received to channel:", v.Channel, "subscription count:", v.Count)
		case error:
			return errors.Wrap(v, "Error while subscribed to redis channel")
		default:
			log.Println("Unknown redis receive during subscription")
		}
	}
}

// broadcast the provided message to all connected websocket connections.
// If an error occurs while writting a message to a websocket connection it is
// closed and deregistered.
func (r *Subscriber) broadcast(msg []byte) {
	r.messages <- msg
}

// Register the websocket connection with the receiver.
func (r *Subscriber) Register(conn *ConnectionData) {
	r.newConnections <- conn
}

// Unregister the connection by closing it and removing it from our list.
func (r *Subscriber) Unregister(conn *ConnectionData) {
	r.rmConnections <- conn
}

// connHandler will handle received data from channel
func (r *Subscriber) connHandler() {
	conns := make([]*ConnectionData, 0)
	for {
		select {
		case msg := <-r.messages:
			for _, conn := range conns {
				for _, handler := range r.MsgHandlers {
					if err := handler(conn, msg); err != nil {
						log.Println("Error when handling connection! Closing and removing Connection")
						conns = removeConn(conns, conn)
					}
				}
				// if err := conn.Socket.WriteMessage(websocket.TextMessage, msg); err != nil {
				// 	log.Println("Error writting data to connection! Closing and removing Connection")
				// 	conns = removeConn(conns, conn)
				// }
			}
		case conn := <-r.newConnections:
			conns = append(conns, conn)
		case conn := <-r.rmConnections:
			conns = removeConn(conns, conn)
		}
	}
}

// removeConn is a function to remove connection from list
func removeConn(conns []*ConnectionData, remove *ConnectionData) []*ConnectionData {
	var i int
	var found bool
	for i = 0; i < len(conns); i++ {
		if conns[i] == remove {
			found = true
			break
		}
	}
	if !found {
		log.Printf("conn: %#v cannot be found\n", remove)
		return conns
	}
	copy(conns[i:], conns[i+1:]) // shift down
	conns[len(conns)-1] = nil    // nil last element
	return conns[:len(conns)-1]  // truncate slice
}
