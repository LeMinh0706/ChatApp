package controller

import (
	"net/http"

	"github.com/LeMinh0706/ChatApp/db"
	"github.com/LeMinh0706/ChatApp/internal/response"
	"github.com/LeMinh0706/ChatApp/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type MessageController struct {
	messageService *service.MessageService
	upgrader       websocket.Upgrader
}

func NewMessageController(messageService *service.MessageService) *MessageController {
	return &MessageController{
		messageService: messageService,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (mc *MessageController) MessageSocket(g *gin.Context) {
	conn, err := mc.upgrader.Upgrade(g.Writer, g.Request, nil)
	if err != nil {
		// log.Println("Error upgrade connect:", err)
		response.ErrorSocket(conn, 400, err.Error())
		return
	}
	defer conn.Close()

	for {
		var msg db.CreateMessageParams
		err := conn.ReadJSON(&msg)
		if err != nil {
			// log.Println("Error read message:", err)
			response.ErrorSocket(conn, 401, err.Error())
			break
		}

		res, err := mc.messageService.SendMessage(g, msg)
		if err != nil {
			// log.Println("Error save message:", err)
			response.ErrorSocket(conn, 404, err.Error())
			break
		}
		response.SuccessSocket(conn, 201, res)
	}
}

func (mc *MessageController) HistoryMessage(g *gin.Context) {
	var req db.GetMessagesParams
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 400, 40000)
		return
	}
	messages, err := mc.messageService.HistoryMessage(g, req)
	if err != nil {
		response.ErrorNonKnow(g, 404, err.Error())
		return
	}
	response.SuccessResponse(g, 200, messages)
}
