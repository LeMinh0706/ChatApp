package server

import "errors"

func (server *Server) InitService() error {
	err := errors.New("Haha")
	if err != nil {
		return err
	}
	return nil
}
