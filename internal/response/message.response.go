package response

import "github.com/LeMinh0706/ChatApp/db"

func MesRes(arg db.CreateMessageParams) db.CreateMessageParams {
	return db.CreateMessageParams{
		FromUser: arg.FromUser,
		ToUser:   arg.ToUser,
		Content:  arg.Content,
	}
}
