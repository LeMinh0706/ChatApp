package middleware

import (
	"fmt"
	"strings"

	"github.com/LeMinh0706/ChatApp/internal/response"
	"github.com/LeMinh0706/ChatApp/internal/token"
	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeaderKey  = "authorization"
	AuthorizationTypeBearer = "bearer"
	AuthorizationPayload    = "authorization_payload"
)

func AuthMiddleware(token token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(AuthorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			response.ErrorResponse(ctx, 401, 40101)
			ctx.Abort()
			return
		}
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			response.ErrorResponse(ctx, 401, 40102)
			ctx.Abort()
			return
		}
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != AuthorizationTypeBearer {
			err := fmt.Errorf("cannot use: %s type", authorizationType)
			response.ErrorNonKnow(ctx, 401, err.Error())
			ctx.Abort()
			return
		}

		accessToken := fields[1]
		payload, err := token.VerifyToken(accessToken)
		if err != nil {
			response.ErrorNonKnow(ctx, 401, err.Error())
			ctx.Abort()
			return
		}
		ctx.Set(AuthorizationPayload, payload)
		ctx.Next()
	}
}
