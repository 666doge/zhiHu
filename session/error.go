package session

import (
	"errors"
)

var (
	ErrSessionNotExists = errors.New("session not exists")
	ErrKeyNotExistInSession = errors.New("key not exists in session")
)