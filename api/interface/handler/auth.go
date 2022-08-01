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
	db := ctx.DynamoDB()

	authorizationValue := c.Request.Header.Get("Authorization")
	token := strings.Split(authorizationValue, " ")[1]

	jwt := entity.NewJwt(token)

	auth := entity.Auth{
		Sub:     ctx.Claim().Sub,
		Payload: jwt.Payload,
	}

	handler.authUseCase.Logout(db, &auth)
	return nil
}
