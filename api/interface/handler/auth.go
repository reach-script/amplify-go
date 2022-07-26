package handler

import (
	"backend/domain/entity"
	"backend/packages/context"
	"backend/packages/utils"
	"backend/usecase"

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

	token := utils.GetJwt(c)
	jwt := utils.NewJwt(token)

	auth := entity.Auth{
		Key1:    ctx.Claim().Sub,
		Key2:    jwt.Payload,
		Payload: jwt.Payload,
	}

	handler.authUseCase.Logout(db, &auth)
	return nil
}
