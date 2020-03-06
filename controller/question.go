package controller

import (
	"fmt"
	"strconv"
	"zhiHu/db"
	"zhiHu/filter"
	"zhiHu/util"
	"zhiHu/model"
	"zhiHu/id_gen"
	"zhiHu/middlewares/account"
	"zhiHu/kafka"

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
	kafka.SendMessage("zhihu_question", question)
	return
}

func GetQuestionList (c *gin.Context) {
	var qList []*model.Question
	qList, err := db.GetQuestionList()
	if err != nil {
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}

	authorMap := map[int64]string{}
	authorIdList := []int64{}
	categoryMap := map[int64]string{}
	categoryIdList := []int64{}
	for _, q := range qList {
		if _, ok := authorMap[q.AuthorId]; ok == false {
			authorMap[q.AuthorId] = ""
			authorIdList = append(authorIdList, q.AuthorId)
		}
		if _, ok := categoryMap[q.CategoryId]; ok == false {
			categoryMap[q.CategoryId] = ""
			categoryIdList = append(categoryIdList, q.CategoryId)
		}
	}

	authorList, err := db.GetUserList(authorIdList)
	categoryList, err := db.GetCategoryListById(categoryIdList)

	qDetailList := []*model.QuestionDetail{}
	for _, q := range qList {
		qDetail := &model.QuestionDetail{
			Question: *q,
		}
		for _, a := range authorList {
			if q.AuthorId == a.UserId {
				qDetail.AuthorName = a.UserName
			}
		}
		for _, c := range categoryList {
			if q.CategoryId == c.CategoryId {
				qDetail.CategoryName = c.CategoryName
			}
		}
		qDetailList = append(qDetailList, qDetail)
	}
	util.RespSuccess(c, qDetailList)
	return
}

func GetQuestionDetail (c *gin.Context) {
	quesIdStr := c.Query("question_id")
	quesId, err := strconv.ParseInt(quesIdStr, 10, 64)

	question, err := db.GetQuestion(quesId)
	if err == db.ErrNoRecord {
		util.RespError(c, util.ErrCodeNoRecord)
		return
	}
	if err != nil {
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}

	cateName, err := db.GetCategoryName(int64(question.CategoryId))
	if err != nil {
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}

	authorName, err := db.GetUserName(int64(question.AuthorId))
	if err != nil {
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}

	questionDetail := &model.QuestionDetail{
		Question: *question,
		CategoryName: cateName,
		AuthorName: authorName,
	}
	util.RespSuccess(c, questionDetail)

}