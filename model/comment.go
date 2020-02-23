package model

import (
	"time"
)

type Comment struct {
	CommentId int64 `json:"commentId" db:"comment_id"`
	Content string `json:"content" db:"content"`
	AnswerId int64 `json:"answerId" db:"answer_id"`
	ParentId int64 `json:"parentId" db:"parent_id"`
	ToUserId int64 `json:"toUserId" db:"to_user_id"`
	FromUserId int64 `json:"fromUserId" db:"from_user_id"`
	IsDel int `json:"isDel" db:"is_del"`
	CreateTime time.Time `json:"createTime" db:"create_time"`
}