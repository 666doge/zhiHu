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
	err = DB.Get(&userId, sqlStr, user.UserName)
	if (err != nil && err != sql.ErrNoRows) {
		return
	}

	if userId > 0 {
		err = ErrUserExist
		return
	}

	sqlStr = `insert into user 
		(user_id, username, nickname, sex, email, phone, password)
		values (?,?,?,?,?,?,?)
	`
	dbPassword := util.Md5([]byte(user.Password + PasswordSalt))
	_, err = DB.Exec(sqlStr, user.UserId, user.UserName, user.NickName, user.Sex, user.Email, user.Phone, dbPassword)
	return
}

func UserLogin(user *model.User) (err error) {
	originPassword := user.Password
	
	sqlStr := "select password, username, user_id from user where username = ?"
	err = DB.Get(user, sqlStr, user.UserName)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	if err == sql.ErrNoRows{
		err = ErrUserNotExist
		return
	}

	ps := util.Md5([]byte(originPassword + PasswordSalt))
	if user.Password != ps {
		err = ErrUserPasswordWrong
		return
	}

	return
}

func GetUserName(userId int64) (userName string, err error) {
	sqlStr := `select username from user where user_id = ?`
	err = DB.Get(&userName, sqlStr, userId)
	return
}