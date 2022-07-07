package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(checkLoggedIn())

	// 200 -> /user/john, 301 -> /user/john/, 404 -> /user/john/get
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.Run(":8080")
}

func checkLoggedIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("ログイン確認処理 前")
		ctx.Next()
		fmt.Println("ログイン確認処理 後")
	}
}
