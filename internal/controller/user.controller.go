package controller

import (
	"github.com/LeMinh0706/ChatApp/internal/response"
	"github.com/LeMinh0706/ChatApp/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(g *gin.Context) {
	var req response.UserRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 400, 40000)
		return
	}
	user, err := uc.userService.Register(g, req.Username, req.Password)
	if err != nil {
		response.ErrorNonKnow(g, 404, err.Error())
		return
	}
	response.SuccessResponse(g, 201, user)
}

func (uc *UserController) Login(g *gin.Context) {
	var req response.UserRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 400, 40000)
		return
	}
	user, err := uc.userService.Login(g, req.Username, req.Password)
	if err != nil {
		response.ErrorNonKnow(g, 404, err.Error())
		return
	}
	response.SuccessResponse(g, 200, user)
}
