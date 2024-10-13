package server

import (
	"fmt"

	"github.com/LeMinh0706/ChatApp/internal/service"
	"github.com/LeMinh0706/ChatApp/internal/token"
	"github.com/LeMinh0706/ChatApp/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Config         util.Config
	Token          token.Maker
	Router         *gin.Engine
	UserService    *service.UserService
	MessageService *service.MessageService
}

func NewServer(config util.Config) (*Server, error) {
	token, err := token.NewJWTMaker(config.SecretKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token: %w", err)
	}
	server := &Server{
		Config: config,
		Token:  token,
		Router: gin.Default(),
	}
	EnableCors(server.Router)
	err = server.InitService()
	if err != nil {
		return nil, err
	}
	Static(server.Router)
	NewRouter(server)

	return server, nil
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}
