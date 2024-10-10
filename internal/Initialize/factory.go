package initialize

import (
	"github.com/LeMinh0706/ChatApp/db"
	"github.com/LeMinh0706/ChatApp/internal/repo"
	"github.com/LeMinh0706/ChatApp/internal/service"
)

type Factory struct {
	UserRepo    *repo.UserRepo
	UserService *service.UserService
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
	userService := service.NewUserService(userRepo)
	return &Factory{
		UserRepo:    userRepo,
		UserService: userService,
	}, nil
}
