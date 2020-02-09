package controller

import (
	"fmt"
	"zhiHu/db"
	"zhiHu/filter"
	"zhiHu/util"
	"zhiHu/model"
	"zhiHu/id_gen"
	"zhiHu/middlewares/account"
	"github.com/gin-gonic/gin"
)

func CreateQuestion (c *gin.Context) {
	var question model.Question
	err := c.BindJSON(&question)
	if err != nil {
		util.RespError(c, util.ErrCodeParameter)
		return
	}

	question.Status = 1
	question.Title = filter.Filter(question.Title)
	question.Content = filter.Filter(question.Content)

	id, err := id_gen.GetId()
	if err != nil {
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}
	question.QuestionId = int64(id)

	// 获取作者id
	authorId, err := account.GetUserId(c)
	fmt.Println(authorId, err)

	if err != nil || authorId <= 0{
		util.RespError(c, util.ErrCodeNotLogin)
		return
	}
	question.AuthorId = authorId

	err = db.CreateQuestion(&question)
	if err != nil {
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}

	util.RespSuccess(c, nil)
	return
}