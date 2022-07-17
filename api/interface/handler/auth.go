package handler

import (
	"backend/domain/entity"
	"backend/packages/context"
	"backend/usecase"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUseCase usecase.Auth
}

func NewAuthHandler(authUseCase usecase.Auth) AuthHandler {
	return AuthHandler{
		authUseCase: authUseCase,
	}
}

func (handler *AuthHandler) Logout(ctx context.Context, c *gin.Context) error {
	ddb := ctx.DynamoDB()

	authorization := c.Request.Header.Get("Authorization")
	value := strings.Split(authorization, " ")[1]
	payload := strings.Split(value, ".")[1]

	auth := entity.Auth{
		Payload: payload,
	}

	handler.authUseCase.Logout(ddb, &auth)
	return nil
}
