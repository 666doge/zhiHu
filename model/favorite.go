package model

import (
	// "time"
)

type FavoriteDir struct {
	DirId int64 `json:"dirId" db:"dir_id"`
	DirName string `json:"dirName" db:"dir_name"`
	UserId int64 `json:"userId" db:"user_id"`
}

type Favorite struct {
	AnswerId int64 `json:"answerId" db:"answer_id"`
	UserId int64 `json:"userId" db:"user_id"`
	DirId int64 `json:"dirId" db:"dir_id"`
}