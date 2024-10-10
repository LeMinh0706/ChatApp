package controller

import (
	"log"
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
		log.Println("Error upgrade connect:", err)
		return
	}
	defer conn.Close()

	for {
		var msg db.CreateMessageParams
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error read message:", err)
			break
		}

		res, err := mc.messageService.SendMessage(g, msg)
		if err != nil {
			log.Println("Error save message:", err)
			break
		}
		response.SuccessResponse(g, 201, res)
		err = conn.WriteJSON(msg)
		if err != nil {
			log.Println("Error send message:", err)
			break
		}
	}
}
