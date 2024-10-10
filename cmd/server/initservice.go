package server

import initialize "github.com/LeMinh0706/ChatApp/internal/Initialize"

func (server *Server) InitService() error {
	factory, err := initialize.NewFactory()
	if err != nil {
		return err
	}

	server.UserService = factory.UserService
	server.MessageService = factory.MessageService
	return nil
}
