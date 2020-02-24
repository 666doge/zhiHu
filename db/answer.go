package db

import (
	"zhiHu/model"
	"zhiHu/logger"

	"github.com/jmoiron/sqlx"
)

func CreateAnswer(answer *model.Answer, questionId int64) (err error) {
	sqlStr := `insert into answer 
		(answer_id, content, author_id)
		values (?,?,?)
	`
	tx, err := DB.Beginx()
	if err != nil {
		return
	}
	_, err = tx.Exec(sqlStr, answer.AnswerId, answer.Content, answer.AuthorId)
	if err != nil {
		tx.Rollback()
		logger.Error("create answer failed, answer:%#v, err:%v", answer, err)
		return
	}

	sqlStr = `insert into question_answer_rel
		(question_id, answer_id)
		values (?,?)
	`
	_, err = tx.Exec(sqlStr, questionId, answer.AnswerId)
	if err != nil {
		tx.Rollback()
		logger.Error("create answer failed, answer:%#v, err:%v", answer, err)
		return
	}
	tx.Commit()
	return
}

func GetAnswerIdList(questionId int64, pageSize int64, pageNo int64) (answerIds []int64, err error) {
	offset := pageSize * (pageNo - 1)
	sqlStr := `select answer_id from question_answer_rel where question_id = ? order by id desc limit ?, ?`
	err = DB.Select(&answerIds, sqlStr, questionId, offset, pageSize)
	if err != nil {
		logger.Error("get answer id list failed, err:%v", err)
		return
	}

	return
}

func GetAnswerList(answerIds []int64) (answerList []*model.Answer, err error) {
	sqlStr := `select 
			answer_id, content, author_id, comment_count, voteup_count, status, create_time, update_time, can_comment
		from answer 
		where answer_id in (?)
	`
	query, args, err := sqlx.In(sqlStr, answerIds)
	if err != nil {
		logger.Error("sqlx.in failed, sqlstr:%v, err:%v", sqlStr, err)
		return
	}

	err = DB.Select(&answerList, query, args...)
	if err != nil {
		logger.Error("get answer list failed, inSqlStr:%v, answerIds: %v,  err:%v",query, answerIds, err)
		return
	}

	return
}

func GetAnswerCount(questionId int64) (answerCount int64, err error) {
	sqlStr := `select count(answer_id) from question_answer_rel where question_id = ?`
	err = DB.Get(&answerCount, sqlStr, questionId)
	if err != nil {
		logger.Error("get answer count failed, err:%v", err)
		return
	}
	return
}

func LikeAnswer(answerId int64) (err error) {
	sqlStr := `update answer set voteup_count=voteup_count + 1 where answer_id = ?`
	_, err = DB.Exec(sqlStr, answerId)
	return
}