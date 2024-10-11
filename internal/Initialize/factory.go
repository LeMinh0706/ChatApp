package initialize

import (
	"github.com/LeMinh0706/ChatApp/db"
	"github.com/LeMinh0706/ChatApp/internal/repo"
	"github.com/LeMinh0706/ChatApp/internal/service"
)

type Factory struct {
	UserRepo       *repo.UserRepo
	MessageRepo    *repo.MessageRepo
	UserService    *service.UserService
	MessageService *service.MessageService
}

func NewFactory() (*Factory, error) {
	pg, err := InitPostgres()
	if err != nil {
		return nil, err
	}
	queries := db.New(pg)

	userRepo, err := repo.NewUserRepo(queries)
	if err != nil {
		return nil, err
	}
	messageRepo, err := repo.NewMessageRepo(queries)
	if err != nil {
		return nil, err
	}

	userService := service.NewUserService(userRepo)
	messageService := service.NewMessageService(messageRepo, userService)

	return &Factory{
		UserRepo:       userRepo,
		MessageRepo:    messageRepo,
		UserService:    userService,
		MessageService: messageService,
	}, nil
}
