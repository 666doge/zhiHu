package model

type Category struct {
	CategoryId int64 `json:"categoryId" db:"category_id"`
	CategoryName string `json:"categoryName" db:"category_name"`
}