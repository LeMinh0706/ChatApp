package controller

import (
	"github.com/LeMinh0706/ChatApp/internal/middleware"
	"github.com/LeMinh0706/ChatApp/internal/response"
	"github.com/LeMinh0706/ChatApp/internal/service"
	"github.com/LeMinh0706/ChatApp/internal/token"
	"github.com/LeMinh0706/ChatApp/util"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
	config      util.Config
	token       token.Maker
}

func NewUserController(config util.Config, token token.Maker, userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
		config:      config,
		token:       token,
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
	token, err := uc.token.CreateToken(user.ID, user.Username, uc.config.TimeDuration)
	if err != nil {
		response.ErrorNonKnow(g, 401, err.Error())
		return
	}
	res := response.LoginRes(user, token)
	response.SuccessResponse(g, 200, res)
}

func (uc *UserController) Yourself(g *gin.Context) {
	auth := g.MustGet(middleware.AuthorizationPayload).(*token.Payload)

	user, err := uc.userService.GetUserById(g, auth.UserId)
	if err != nil {
		response.ErrorNonKnow(g, 401, err.Error())
		return
	}

	response.SuccessResponse(g, 200, user)
}
