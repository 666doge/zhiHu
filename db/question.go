package db

import (
	"database/sql"
	"zhiHu/model"
	"fmt"
)

func CreateQuestion(q *model.Question) (err error) {
	sqlStr := ` insert question
		(title, content, status, author_id, question_id, category_id)
		values (?,?,?,?,?,?)
	`
	_, err = DB.Exec(sqlStr, q.Title, q.Content, q.Status, q.AuthorId, q.QuestionId, q.CategoryId)
	if err != nil {
		fmt.Printf("create question failed, question:%#v, err:%v", q, err)
		return
	}
	return
}

func GetQuestionList() (qList []*model.Question, err error) {
	sqlStr := `select 
			title, content, author_id, category_id, question_id
		from question
		limit 5
	`
	err = DB.Select(&qList, sqlStr)
	return
}

func GetQuestion(quesId int64) (question *model.Question, err error) {
	question = &model.Question{}
	sqlStr := `select
			question_id, title, content, category_id, author_id, status, create_time
		from question 
		where question_id = ? `
	err = DB.Get(question, sqlStr, quesId)
	if err == sql.ErrNoRows {
		err = ErrNoRecord
	}
	return
}

