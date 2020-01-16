package main

import (
	// "fmt"
	"zhiHu/controller"
	"zhiHu/middlewares/account"

	"github.com/gin-gonic/gin"
)

func main(){
	initService()
}

func initService() {
	r := gin.Default()
	r.Use(account.Auth())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user", controller.GetUserList)
	r.Run()
}