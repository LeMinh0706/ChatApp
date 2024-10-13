package router

import (
	"github.com/LeMinh0706/ChatApp/internal/controller"
	"github.com/LeMinh0706/ChatApp/internal/middleware"
	"github.com/LeMinh0706/ChatApp/internal/service"
	"github.com/LeMinh0706/ChatApp/internal/token"
	"github.com/LeMinh0706/ChatApp/util"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(r *gin.Engine, router *gin.RouterGroup, config util.Config, token token.Maker, userService *service.UserService) {
	uc := controller.NewUserController(config, token, userService)
	userGroup := r.Group(router.BasePath() + "/users")
	auth := userGroup.Group("", middleware.AuthMiddleware(token))
	{
		userGroup.POST("register", uc.Register)
		userGroup.POST("login", uc.Login)
		auth.GET("yourself", uc.Yourself)
	}
}
