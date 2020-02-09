package main

import (
	// "fmt"
	"zhiHu/controller"
	"zhiHu/middlewares/account"
	"zhiHu/id_gen"
	"zhiHu/db"
	"zhiHu/session"
	"zhiHu/filter"

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

	// init mysql
	err = db.Init("root:xsN231564@tcp(localhost:3306)/zhihu?parseTime=true")
	if err != nil {
		panic(err)
	}

	err = session.Init("memory", "")
	if err != nil {
		panic(err)
	}

	// init sensitive word filter
	err = filter.Init("./filter/sensitiveWords.txt")
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/user/register", controller.UserRegister)
	r.POST("/user/login", controller.UserLogin)

	// qRouter := r.Group("/question")
	qRouter := r.Group("/question", account.Auth())
	{
		qRouter.POST("/detail", controller.CreateQuestion)
	}

	r.Run()
}