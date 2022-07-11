package handler

import (
	"backend/domain/entity"
	"backend/packages/context"
	"backend/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createUserParams struct {
	Name       string `form:"name" json:"name"`
	CognitoSub string `form:"cognito_sub" json:"cognito_sub"`
}
type updateUserParams struct {
	ID         uint   `form:"id" json:"id"`
	Name       string `form:"name" json:"name"`
	CognitoSub string `form:"cognito_sub" json:"cognito_sub"`
}

type UserHandler struct {
	userUseCase usecase.User
}

func NewUserHandler(userUseCase usecase.User) UserHandler {
	return UserHandler{
		userUseCase: userUseCase,
	}
}

func (handler *UserHandler) Create(ctx context.Context, c *gin.Context) error {
	db := ctx.DB()
	params := createUserParams{}
	if err := c.BindJSON(&params); err != nil {
		panic(err)
	}

	userEntity := entity.User{
		Name:       params.Name,
		CognitoSub: params.CognitoSub,
	}
	user, err := handler.userUseCase.Create(db, &userEntity)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, user)
	return nil
}

func (handler *UserHandler) Update(ctx context.Context, c *gin.Context) error {
	db := ctx.DB()
	params := updateUserParams{}
	if err := c.BindJSON(&params); err != nil {
		panic(err)
	}

	user := entity.User{
		Name:       params.Name,
		CognitoSub: params.CognitoSub,
	}
	user.ID = params.ID
	_, err := handler.userUseCase.Update(db, &user)
	if err != nil {
		log.Println("%v", err)
		panic(err)
	}

	c.Status(http.StatusOK)
	return nil
}

func (handler *UserHandler) Delete(ctx context.Context, c *gin.Context) error {
	db := ctx.DB()

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if err := handler.userUseCase.Delete(db, uint(id)); err != nil {
		panic(err)
	}

	c.Status(http.StatusOK)
	return nil
}

func (handler *UserHandler) GetByID(ctx context.Context, c *gin.Context) error {
	db := ctx.DB()
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	user, err := handler.userUseCase.GetByID(db, uint(id))

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, user)
	return nil
}
