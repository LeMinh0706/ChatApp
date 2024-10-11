package service

import (
	"context"

	"github.com/LeMinh0706/ChatApp/db"
	"github.com/LeMinh0706/ChatApp/internal/repo"
	"github.com/LeMinh0706/ChatApp/internal/response"
)

type MessageService struct {
	messageRepo *repo.MessageRepo
	userService *UserService
}

func NewMessageService(messageRepo *repo.MessageRepo, userService *UserService) *MessageService {
	return &MessageService{
		messageRepo: messageRepo,
		userService: userService,
	}
}

func (ms *MessageService) SendMessage(ctx context.Context, arg db.CreateMessageParams) (response.MessageResponse, error) {
	var res response.MessageResponse
	user1, err := ms.userService.GetUserById(ctx, arg.FromID)
	if err != nil {
		return response.MessageResponse{}, err
	}
	user2, err := ms.userService.GetUserById(ctx, arg.ToID)
	if err != nil {
		return response.MessageResponse{}, err
	}
	message, err := ms.messageRepo.CreateMessage(ctx, arg)
	if err != nil {
		return res, err
	}
	res = response.MesRes(message, user1, user2)
	return res, nil
}

func (ms *MessageService) HistoryMessage(ctx context.Context, arg db.GetMessagesParams) ([]response.MessageResponse, error) {
	var res []response.MessageResponse
	user1, err := ms.userService.GetUserById(ctx, arg.FromID)
	if err != nil {
		return res, err
	}
	user2, err := ms.userService.GetUserById(ctx, arg.ToID)
	if err != nil {
		return res, err
	}

	messages, err := ms.messageRepo.ListMessage(ctx, arg)
	if err != nil {
		return nil, err
	}
	for _, message := range messages {
		if message.FromID == arg.FromID {
			mesRes := response.MesRes(message, user1, user2)
			res = append(res, mesRes)
		} else {
			mesRes := response.MesRes(message, user2, user1)
			res = append(res, mesRes)
		}
	}

	if len(res) == 0 {
		return []response.MessageResponse{}, err
	}
	return res, nil
}
