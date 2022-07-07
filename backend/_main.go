package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	router := gin.Default()

	router.Use(checkLoggedIn())

	// 200 -> /user/john, 301 -> /user/john/, 404 -> /user/john/get
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 200 -> /user/john/get, /user/john/get/ok
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.POST("/user/:name/*action", func(c *gin.Context) {
		// /user/john/get/ok も /user/:name/*actionと等価になる
		fmt.Println(c.FullPath() == "/user/:name/*action")
	})

	// welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		fmt.Printf("%s", lastname)

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.GET("/person", func(c *gin.Context) {
		var person Person
		err := c.ShouldBindQuery(&person)
		if err != nil {
			fmt.Print(("invalid query parameter"))
			c.String(http.StatusBadRequest, "invalid params")
			return
		}
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	})

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	v1 := router.Group("/v1")
	v1.GET("/get", func(c *gin.Context) {
		c.String(200, "/v1/GET")
	})

	v2 := router.Group("/v2")
	{
		v2.GET("/get", func(c *gin.Context) {
			c.String(200, "/v2/GET")
		})
	}

	router.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		c.JSON(http.StatusOK, msg)
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
