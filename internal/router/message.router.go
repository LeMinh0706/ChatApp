package router

import (
	"github.com/LeMinh0706/ChatApp/internal/controller"
	"github.com/LeMinh0706/ChatApp/internal/service"
	"github.com/gin-gonic/gin"
)

func NewMessageRouter(r *gin.Engine, router *gin.RouterGroup, messageService *service.MessageService) {
	c := controller.NewMessageController(messageService)
	messageGroup := r.Group(router.BasePath() + "/messages")
	{
		messageGroup.GET("", c.HistoryMessage)
	}
	r.GET("/ws", c.MessageSocket)
}
