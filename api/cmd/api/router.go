package api

import (
	"backend/infrastructure/database"
	"backend/infrastructure/persistence"
	"backend/interface/handler"
	"backend/interface/middleware"
	"backend/packages/context"
	"backend/packages/errors"
	"backend/usecase"

	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine) {
	userPersistance := persistence.NewUserPersistance()
	authPersistance := persistence.NewAuthPersistance()

	userUseCase := usecase.NewUserUseCase(userPersistance)
	authUseCase := usecase.NewAuthUseCase(authPersistance)

	userHandler := handler.NewUserHandler(userUseCase)
	userApi := r.Group("users")
	userApi.Use(middleware.Auth)
	userApi.GET("/:id", wrapperFunc(userHandler.GetByID))
	userApi.POST("/", wrapperFunc(userHandler.Create))
	userApi.PATCH("/", wrapperFunc(userHandler.Update))
	userApi.DELETE("/:id", wrapperFunc(userHandler.Delete))

	authHandler := handler.NewAuthHandler(authUseCase)
	authApi := r.Group("auth")
	authApi.Use(middleware.Auth)
	authApi.POST("/logout", wrapperFunc(authHandler.Logout))

}

type HandlerFunc func(ctx context.Context, c *gin.Context) errors.IError

func wrapperFunc(handlerFunc HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := context.New(c, database.GetRDB, database.GetDynamoDB)

		if err := handlerFunc(ctx, c); err != nil {
			// TODO: request headerからidを取得する仕組みを実装
			requestID := ""
			err.Response().Do(c, requestID)
			c.Error(err)
		}
	}
}
