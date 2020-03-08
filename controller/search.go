package controller

import (
	"zhiHu/util"
	"zhiHu/model"
	"zhiHu/es"
	"encoding/json"

	"github.com/gin-gonic/gin"
)


func Search(c *gin.Context){
	searchResult := &model.SearchResult{
		QuestionList: []*model.Question{},
		AnswerList: []*model.Answer{},
	}
	q := c.Query("query")

	// 取问题
	qHits, err := es.SearchByMatchQuery("zhihu_question", "title", q)
	if err != nil {
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}
	for _, hit := range qHits {
		var ques model.Question
		err = json.Unmarshal(hit.Source, &ques)
		if err != nil {
			util.RespError(c, util.ErrCodeServerBusy)
			return
		}
		searchResult.QuestionList = append(searchResult.QuestionList, &ques)
	}

	// 取答案
	aHits, err := es.SearchByMatchQuery("zhihu_answer", "content", q)
	if err != nil {
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}
	for _, hit := range aHits {
		var a model.Answer
		err = json.Unmarshal(hit.Source, &a)
		if err != nil {
			util.RespError(c, util.ErrCodeServerBusy)
			return
		}
		searchResult.AnswerList = append(searchResult.AnswerList, &a)
	}
	util.RespSuccess(c, searchResult)
	return
}