package response

import (
	"time"

	"github.com/LeMinh0706/ChatApp/db"
)

type MessageResponse struct {
	FromID      int64             `json:"from_id"`
	ToID        int64             `json:"to_id"`
	Content     string            `json:"content"`
	FromUser    db.GetUserByIdRow `json:"from_user"`
	ToUser      db.GetUserByIdRow `json:"to_user"`
	DateCreated time.Time         `json:"date_created"`
}

func MesRes(arg db.Message, user1, user2 db.GetUserByIdRow) MessageResponse {
	return MessageResponse{
		FromID:      arg.FromID,
		ToID:        arg.ToID,
		Content:     arg.Content,
		FromUser:    user1,
		ToUser:      user2,
		DateCreated: arg.DateCreated,
	}
}
