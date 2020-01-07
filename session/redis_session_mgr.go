package session

import (
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
)

type RedisSessionManager struct {
	addr string
	passwd string
	pool *redis.Pool
	rwLock sync.RWMutex
	sessionMap map[string]Session
}

func newPool(addr string, passwd string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			/*
			   if _, err := c.Do("AUTH", password); err != nil {
			   c.Close()
			   return nil, err
			   }*/
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		}
	}
}

func (rsm *RedisSessionManager) Init(addr string, options ...string) {
	if len(options) > 0 {
		rsm.passwd = options[0]
	}
	rsm.pool = newPool(addr, rsm.passwd)
	rsm.addr = addr
	return
}

func (rsm *RedisSessionManager) CreateSession() (session Session, err error) {
	rsm.rwLock.Lock()
	defer rsm.rwLock.Unlock()

	uuid, err := uuid.NewV4()
	if err != nil {
		return
	}
	sessionId := uuid.String()
	session := NewRedisSession()
	rsm.sessionMap[sessionId] = session

	return
}

func (rsm *RedisSessionManager) Get(sessionId string) (session Session, err error) {
	rsm.rwLock.RLock()
	defer rsm.rwLock.RUnlock()

	session, ok := rsm.sessionMap[sessionId]
	if !ok {
		err = ErrSessionNotExists
		return
	}
	return
}
