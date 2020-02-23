package db

import (
	"zhiHu/model"
)

func CreateComment (comment *model.Comment) (err error) {
	tx, err := DB.Beginx()
	if err != nil {
		return
	}
	sqlx := `insert into comment 
		(comment_id, content, answer_id, to_user_id, from_user_id, parent_id)
		values (?, ?, ?, ?, ?, ?)
	`
	_, err = tx.Exec(
		sqlx,
		comment.CommentId,
		comment.Content,
		comment.AnswerId,
		comment.ToUserId,
		comment.FromUserId,
		comment.ParentId,
	)
	if err != nil {
		tx.Rollback()
		return
	}
	sqlx = `update answer
			set comment_count = comment_count + 1 
			where answer_id = ?
		`
	_, err = tx.Exec(sqlx, comment.AnswerId)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func GetCommentList(answerId int64) (commentList []*model.Comment, err error) {
	sqlStr := `select 
			content, comment_id, answer_id, parent_id, to_user_id, from_user_id, create_time
		from comment 
		where answer_id = ? and is_del = 0
	`
	err = DB.Select(&commentList, sqlStr, answerId)
	return
}