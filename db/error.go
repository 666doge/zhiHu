package db

import (
	"errors"
)

var (
	ErrCodeUserExist = errors.New("username exists")
)