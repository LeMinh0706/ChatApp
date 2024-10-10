package service

import (
	"context"

	"github.com/LeMinh0706/ChatApp/db"
	"github.com/LeMinh0706/ChatApp/internal/repo"
)

type MessageService struct {
	messageRepo *repo.MessageRepo
}

func NewMessageService(messageRepo *repo.MessageRepo) *MessageService {
	return &MessageService{
		messageRepo: messageRepo,
	}
}

func (ms *MessageService) SendMessage(ctx context.Context, arg db.CreateMessageParams) (db.Message, error) {
	var res db.Message
	message, err := ms.messageRepo.CreateMessage(ctx, arg)
	if err != nil {
		return res, err
	}
	res = message
	return res, nil
}
