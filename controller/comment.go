package controller

import (
	"zhiHu/model"
	"zhiHu/util"
	"zhiHu/logger"
	"zhiHu/id_gen"
	"zhiHu/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var comment model.Comment
	err := c.BindJSON(&comment)
	if err != nil {
		logger.Error("bind json failed, err: %v", err)
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}

	if comment.FromUserId == 0 || comment.AnswerId == 0 {
		logger.Error("invalid params, fromuserid: %v, answerId: %v", comment.FromUserId, comment.AnswerId)
		util.RespError(c, util.ErrCodeParameter)
		return
	}

	cid, err := id_gen.GetId()
	comment.CommentId = int64(cid)
	err = db.CreateComment(&comment)
	if err != nil {
		logger.Error("insert comment failed, comment: %#v,  err: %v", comment, err)
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}
	util.RespSuccess(c, nil)
}

func GetCommentList (c *gin.Context) {
	answerIdStr := c.Query("answerId")
	answerId, err := strconv.ParseInt(answerIdStr, 10, 64)
	if err != nil {
		logger.Error("convert answer id failed, answerIdStr: %v", answerIdStr)
		util.RespError(c, util.ErrCodeParameter)
		return
	}

	commentList, err := db.GetCommentList(answerId)
	if err != nil {
		logger.Error("get comment list failed, answerId: %v", answerId)
		util.RespError(c, util.ErrCodeServerBusy)
		return
	}

	util.RespSuccess(c, commentList)
	return;
}