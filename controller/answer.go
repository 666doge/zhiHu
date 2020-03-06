package controller

import (
	"html"
	"strconv"
	"zhiHu/model"
	"zhiHu/util"
	"zhiHu/logger"
	"zhiHu/id_gen"
	"zhiHu/middlewares/account"
	"zhiHu/db"
	"zhiHu/kafka"

	"github.com/gin-gonic/gin"
)

func CreateAnswer(c *gin.Context) {
	var answer model.Answer
	err := c.BindJSON(&answer)
	if err != nil {
		logger.Error("bind json failed, err:%v", err)
		util.RespError(c, util.ErrCodeParameter)
		return
	}

	userId, err := account.GetUserId(c)
	if err != nil || userId == 0 {
		logger.Error("get user id failed %v : %v", userId, err)
		util.RespError(c, util.ErrCodeNotLogin)
		return
	}
	answer.AuthorId = userId
	answer.Content = html.EscapeString(answer.Content) // 转义，防止xss攻击

	// 生成 答案id
	cid, err := id_gen.GetId()
	if err != nil {
		logger.Error("generate answer id failed %v : %v", cid, err)
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}
	answer.AnswerId = int64(cid)

	err = db.CreateAnswer(&answer, answer.QuestionId)
	if err != nil {
		logger.Error("insert answer failed,  %#v : %v", answer, err)
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}
	util.RespSuccess(c, nil)
	kafka.SendMessage("zhihu_answer", answer)
}

func GetAnswerList(c *gin.Context) {
	questionIdStr := c.Query("questionId")
	questionId, _ := strconv.ParseInt(questionIdStr, 10, 64)

	pageSizeStr := c.Query("pageSize")
	pageSize, _ := strconv.ParseInt(pageSizeStr, 10, 64)

	pageNoStr := c.Query("pageNo")
	pageNo, _ := strconv.ParseInt(pageNoStr, 10, 64)

	answerIds, err := db.GetAnswerIdList(questionId, pageSize, pageNo)
	if err != nil {
		logger.Error("get answer ids failed, questionId: %v, err: %v", questionId, err)
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}

	if len(answerIds) == 0 {
		util.RespSuccess(c, "")
		return
	}

	answerList, err := db.GetAnswerList(answerIds)
	if err != nil {
		logger.Error("get answer list failed, answerIds: %v, err: %v", answerIds, err)
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}

	var userIds []int64
	for _, a := range answerList {
		userIds = append(userIds, a.AuthorId)
	}

	userList, err := db.GetUserList(userIds)
	if err != nil {
		logger.Error("get user list failed, userIds: %v, err: %v", userIds, err)
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}

	answerDetailList := &model.AnswerDetailList{}
	for _, a := range answerList {
		answerDetail := &model.AnswerDetail{
			Answer: *a,
		}

		for _, u := range userList {
			if u.UserId == a.AuthorId {
				answerDetail.AuthorName = u.UserName
				answerDetail.QuestionId = questionId
				break
			}
		}
		answerDetailList.AnswerList = append(answerDetailList.AnswerList, answerDetail)
	}

	// 获取totalCount
	totalCount, err := db.GetAnswerCount(questionId)
	if err != nil {
		logger.Error("get totalcount failed, questionId: %v, err: %v", questionId, err)
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}
	answerDetailList.TotalCount = totalCount

	util.RespSuccess(c, answerDetailList)
	return
}

func LikeAnswer(c *gin.Context) {
	answerIdStr := c.Query("answerId")
	answerId, _ := strconv.ParseInt(answerIdStr, 10, 64)
	err := db.LikeAnswer(answerId)
	if err != nil {
		logger.Error("update like number failed, answerId: %v, err: %v", answerId, err)
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}
	util.RespSuccess(c, nil)
	return
}