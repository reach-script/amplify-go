package api

import (
	"backend/infrastructure/database"
	"backend/infrastructure/persistence"
	"backend/interface/handler"
	"backend/interface/middleware"
	"backend/packages/context"
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

type HandlerFunc func(ctx context.Context, c *gin.Context) error

func wrapperFunc(handlerFunc HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := context.New(c, database.GetRDB, database.GetDynamoDB)

		handlerFunc(ctx, c)

		// if err != nil {
		// 	switch v := err.(type) {
		// 	case *errors.Error:
		// 		v.Response().Do(c, ctx.RequestID())
		// 	default:
		// 		errors.NewUnexpected(v).Response().Do(c, ctx.RequestID())
		// 	}

		// 	_ = c.Error(err)
		// }
	}
}
