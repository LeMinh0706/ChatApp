package server

import "github.com/LeMinh0706/ChatApp/internal/router"

func NewRouter(s *Server) {
	v1 := s.Router.Group("/api")
	{
		router.NewUserRouter(s.Router, v1, s.Config, s.Token, s.UserService)
		router.NewMessageRouter(s.Router, v1, s.MessageService)
	}
}
