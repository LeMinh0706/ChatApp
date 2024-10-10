package router

import (
	"github.com/LeMinh0706/ChatApp/internal/controller"
	"github.com/LeMinh0706/ChatApp/internal/service"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(r *gin.Engine, router *gin.RouterGroup, userService *service.UserService) {
	uc := controller.NewUserController(userService)
	userGroup := r.Group(router.BasePath() + "/users")
	{
		userGroup.POST("register", uc.Register)
		userGroup.POST("login", uc.Login)
	}
}
