package server

import (
	"github.com/LeMinh0706/ChatApp/internal/service"
	"github.com/LeMinh0706/ChatApp/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Config         util.Config
	Router         *gin.Engine
	UserService    *service.UserService
	MessageService *service.MessageService
}

func NewServer(config util.Config) (*Server, error) {

	server := &Server{
		Config: config,
		Router: gin.Default(),
	}
	err := server.InitService()
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
