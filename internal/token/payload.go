package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Payload struct {
	Id        uuid.UUID `json:"id"`
	UserId    int64     `json:"user_id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
	jwt.RegisteredClaims
}

func NewPayLoad(user_id int64, username string, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		Id:        tokenId,
		UserId:    user_id,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}
