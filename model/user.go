package model

const (
	UserSexMan = 1
	UserSexWoman = 2
)

type User struct {
	UserId int64 `json:"userId" db:"user_id"`
	Username string `json:"userName" db:"username"`
	Nickname string `json:"nickName" db:"nick_name"`
	Phone string `json:"phone" db:"phone"`
	Password string `json:"password" db:"password"`
	Email string `json:"email" db:"email"`
	Sex int `json:"sex" db:"sex"`
}