package router

import (
	"github.com/LeMinh0706/ChatApp/internal/controller"
	"github.com/LeMinh0706/ChatApp/internal/service"
	"github.com/gin-gonic/gin"
)

func NewMessageRouter(r *gin.Engine, messageService *service.MessageService) {
	c := controller.NewMessageController(messageService)
	r.GET("/ws", c.MessageSocket)
}
