package api

import (
	"backend/config"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Execute() {
	router := gin.Default()

	registerRoutes(router)

	router.Run(fmt.Sprintf(":%s", config.Env.API.Port))
}
