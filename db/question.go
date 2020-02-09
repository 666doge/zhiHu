package db

import (
	// "database/sql"
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