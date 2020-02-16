package model

import(
	"time"
)

type Question struct {
	QuestionId int64 `json:"questionId" db:"question_id"`
	Title string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
	AuthorId int64 `json:"authorId" db:"author_id"`
	CategoryId int64 `json:"categoryId" db:"category_id"`
	Status int `json:"status" db:"status"`
	CreateTime time.Time `json:"createTime" db:"create_time"`
}

type QuestionDetail struct {
	Question *Question `json:"question"`
	AuthorName string `json:"authorName"`
	CategoryName string `json:"categoryName"`
}