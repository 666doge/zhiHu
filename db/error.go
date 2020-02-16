package db

import (
	"errors"
)

var (
	ErrUserExist = errors.New("username exists")
	ErrUserNotExist = errors.New("user not exists")
	ErrUserPasswordWrong = errors.New("password is wrong")
	ErrNoRecord = errors.New("no record")
)