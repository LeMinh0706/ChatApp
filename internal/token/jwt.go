package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTMaker struct {
	secretKey string
}

var ErrUnexpected = errors.New("unexpected signing method")
var ErrParse = errors.New("could not parse token")
var ErrOutDate = errors.New("token out of date")
var ErrLenKey = errors.New("secret key must be at least 32")

// CreateToken implements Maker.
func (j *JWTMaker) CreateToken(user_id int64, username string, duration time.Duration) (string, error) {
	payload, err := NewPayLoad(user_id, username, duration)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         payload.Id,
		"user_id":    payload.UserId,
		"username":   payload.Username,
		"issued_at":  payload.IssuedAt,
		"expired_at": payload.ExpiredAt,
	})
	return token.SignedString([]byte(j.secretKey))
}

// VerifyToken implements Maker.
func (j *JWTMaker) VerifyToken(token string) (*Payload, error) {
	claims := &Payload{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpected
		}
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return nil, ErrParse
	}
	if time.Now().After(claims.ExpiredAt) {
		return nil, ErrOutDate
	}
	return claims, nil
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < 32 {
		return nil, ErrLenKey
	}
	return &JWTMaker{secretKey}, nil
}
