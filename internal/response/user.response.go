package response

import "github.com/LeMinh0706/ChatApp/db"

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	UrlAvatar string `json:"url_avatar"`
}

func UserRes(user db.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		UrlAvatar: user.UrlAvatar,
	}
}
