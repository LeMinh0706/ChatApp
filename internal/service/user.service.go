package service

import (
	"context"

	"github.com/LeMinh0706/ChatApp/db"
	"github.com/LeMinh0706/ChatApp/internal/repo"
	"github.com/LeMinh0706/ChatApp/internal/response"
	"github.com/LeMinh0706/ChatApp/util"
)

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService(repo *repo.UserRepo) *UserService {
	return &UserService{
		userRepo: repo,
	}
}

func (us *UserService) Register(ctx context.Context, username, password string) (response.UserResponse, error) {
	var res response.UserResponse
	hashPassword, err := util.HashPashword(password)
	if err != nil {
		return res, err
	}
	url := util.RandomURL()
	user, err := us.userRepo.CreateUser(ctx, db.CreateUserParams{
		Username:  username,
		Password:  hashPassword,
		UrlAvatar: url,
	})

	if err != nil {
		return res, err
	}

	res = response.UserRes(user)

	return res, nil
}

func (us *UserService) Login(ctx context.Context, username, password string) (response.UserResponse, error) {
	var res response.UserResponse
	user, err := us.userRepo.GetUser(ctx, username)
	if err != nil {
		return res, err
	}
	if err = util.CheckPassword(password, user.Password); err != nil {
		return res, err
	}
	res = response.UserRes(user)
	return res, nil
}
