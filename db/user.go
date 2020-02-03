package db

import (
	"zhiHu/model"
	"zhiHu/util"
	"database/sql"
	// "github.com/jmoiron/sqlx"
)

const (
	PasswordSalt = "HBZciU2SiSDr4uPeJ1e7qlIlMbyusQ0v"
)

func Register(user *model.User) (err error) {
	var userId int64
	sqlStr := `select user_id from user where username=?`
	err = DB.Get(&userId, sqlStr, user.Username)
	if (err != nil && err != sql.ErrNoRows) {
		return
	}

	if userId > 0 {
		err = ErrCodeUserExist
		return
	}

	sqlStr = `insert into user 
		(user_id, username, nickname, sex, email, phone, password)
		values (?,?,?,?,?,?,?)
	`
	dbPassword := util.Md5([]byte(user.Password + PasswordSalt))
	_, err = DB.Exec(sqlStr, user.UserId, user.Username, user.Nickname, user.Sex, user.Email, user.Phone, dbPassword)
	return
}