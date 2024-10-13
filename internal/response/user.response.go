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

type LoginResponse struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	UrlAvatar   string `json:"url_avatar"`
	AccessToken string `json:"access_token"`
}

func UserRes(user db.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		UrlAvatar: user.UrlAvatar,
	}
}

func LoginRes(user UserResponse, token string) LoginResponse {
	return LoginResponse{
		ID:          user.ID,
		Username:    user.Username,
		UrlAvatar:   user.UrlAvatar,
		AccessToken: token,
	}

}
