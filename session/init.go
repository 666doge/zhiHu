package session

import (
	"fmt"
)

var sessionMgr SessionManager

func Init(sessionType string, addr string, options ...string) (err error) {
	switch(sessionType) {
	case "redis":
		sessionMgr = NewRedisSessionManager()
	case "memory":
		sessionMgr = NewMemorySessionManager()
	default:
		err = fmt.Errorf("not support")
	}
	sessionMgr.Init(addr, options...)
	return
}

func CreateSession() (session Session, err error) {
	return sessionMgr.CreateSession()
}

func Get(sessionId string) (session Session, err error) {
	return sessionMgr.Get(sessionId)
}

