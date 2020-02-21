package main

import (
	"fmt"
	"zhiHu/controller"
	"zhiHu/middlewares/account"
	"zhiHu/middlewares"
	"zhiHu/id_gen"
	"zhiHu/db"
	"zhiHu/session"
	"zhiHu/filter"
	"zhiHu/util"
	"zhiHu/logger"

	"github.com/gin-gonic/gin"
)

func main(){
	initService()
}

func initService() {
	// init logger
	logger.InitLogger("file", map[string]string {
		"log_level": "file",
		"log_path": util.GetWorkDirectory() + "/logs",
		"log_name": "main",
		"log_split_type": "size",
		"log_split_size": "1024",
	})
	fmt.Println("init logger ok")

	// init id generator
	err := id_gen.Init(1)
	if err != nil {
		panic(err)
	}

	// init mysql
	err = db.Init("root:xsN231564@tcp(localhost:3306)/zhihu?parseTime=true")
	if err != nil {
		panic(err)
	}

	// init reids
	err = session.Init("localhost:6379", "")
	if err != nil {
		panic(err)
	}

	// init sensitive word filter
	err = filter.Init("./filter/sensitiveWords.txt")
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(middlewares.Logger())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/user/register", controller.UserRegister)
	r.POST("/user/login", controller.UserLogin)

	r.GET("/category/list", controller.GetCategoryList)

	// qRouter := r.Group("/question")
	qRouter := r.Group("/question", account.Auth())
	{
		qRouter.POST("/add", controller.CreateQuestion)
		qRouter.GET("/detail", controller.GetQuestionDetail)
		qRouter.GET("/list", controller.GetQuestionList)
	}

	r.Run()
}