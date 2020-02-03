package main

import (
	// "fmt"
	"zhiHu/controller"
	"zhiHu/middlewares/account"
	"zhiHu/id_gen"
	"zhiHu/db"

	"github.com/gin-gonic/gin"
)

func main(){
	initService()
}

func initService() {
	err := id_gen.Init(1)
	if err != nil {
		panic(err)
	}

	err = db.Init("root:xsN231564@tcp(localhost:3306)/zhihu?parseTime=true")
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(account.Auth())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/user/register", controller.UserRegister)
	// r.GET("/user/list", controller.GetUserList)
	r.Run()
}