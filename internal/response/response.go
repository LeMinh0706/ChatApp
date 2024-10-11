package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code, status int) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[status],
		Data:    nil,
	})
}

func ErrorNonKnow(c *gin.Context, code int, massage string) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: massage,
		Data:    nil,
	})
}

func SuccessSocket(conn *websocket.Conn, code int, data interface{}) {
	err := conn.WriteJSON(ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
	if err != nil {
		ErrorSocket(conn, 500, "Interval")
	}
}

func ErrorSocket(conn *websocket.Conn, code int, message string) {
	conn.WriteJSON(ResponseData{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
