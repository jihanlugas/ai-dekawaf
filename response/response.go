package response

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, code int, msg string, payload interface{}) {
	response := Response{
		Status:  true,
		Message: msg,
		Data:    payload,
	}

	c.JSON(code, response)
}

func Error(c *gin.Context, code int, msg string, err error, payload interface{}) {
	errorMessage := ""

	if err != nil {
		errorMessage = err.Error()
	}

	response := Response{
		Status:  false,
		Error:   errorMessage,
		Message: msg,
		Data:    payload,
	}

	c.JSON(code, response)
}

func ErrorMiddleware(c *gin.Context, code int, msg string, err error, payload interface{}) {
	errorMessage := ""

	if err != nil {
		errorMessage = err.Error()
	}

	response := Response{
		Status:  false,
		Error:   errorMessage,
		Message: msg,
		Data:    payload,
	}

	c.AbortWithStatusJSON(code, response)
}
