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

type HandlerFunc func(ctx context.Context, c *gin.Context) error

func registerRoutes(r *gin.Engine) {
	userPersistance := persistence.NewUserPersistance()

	userUseCase := usecase.NewUserUseCase(userPersistance)

	userHandler := handler.NewUserHandler(userUseCase)

	userApi := r.Group("users")
	userApi.Use(middleware.WithAuth)
	userApi.GET("/:id", wrapperFunc(userHandler.GetByID))
	userApi.POST("/", wrapperFunc(userHandler.Create))
	userApi.PATCH("/", wrapperFunc(userHandler.Update))
	userApi.DELETE("/:id", wrapperFunc(userHandler.Delete))
}

func wrapperFunc(handlerFunc HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := context.New(c, database.Get)

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
