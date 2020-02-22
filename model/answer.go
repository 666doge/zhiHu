package model

import (
	"time"
)

type Answer struct {
	AnswerId int64 `json:"answerId" db:"answer_id"`
	Content string `json:"content" db:"content"`
	CommentCount int32     `json:"commentCount" db:"comment_count"`
	VoteupCount  int32     `json:"voteup_count" db:"voteup_count"`
	AuthorId     int64     `json:"authorId" db:"author_id"`
	Status       int32     `json:"status" db:"status"`
	CanComment   int32     `json:"canComment" db:"can_comment"`
	CreateTime   time.Time `json:"createTime" db:"create_time"`
	UpdateTime   time.Time `json:"updateTime" db:"update_time"`
	QuestionId   int64    `json:"questionId"`
}

type AnswerDetail struct {
	Answer
	AuthorName string `json:"authorName"`
}

type AnswerDetailList struct {
	AnswerList []*AnswerDetail `json:"answerList"`
	TotalCount int64 `json:"totalCount"`
}