package session

import (
	"sync"
	"encoding/json"

	"github.com/garyburd/redigo/redis"
)

const (
	SessionFlagNone = iota
	SessionFlagModify
	SessionFlagLoad
)

type RedisSession struct {
	data map[string]interface{}
	id string
	rwLock sync.RWMutex
	pool *redis.Pool
	flag int
}

func NewRedisSession (id string, pool *redis.Pool) *RedisSession {
	s := &RedisSession{
		data: make(map[string]interface{}, 8),
		id: id,
		pool: pool,
		flag: SessionFlagNone,
	}
	return s
}

func (rs *RedisSession) Set (key string, value interface{}) (err error) {
	rs.rwLock.Lock()
	defer rs.rwLock.Unlock()

	rs.data[key] = value
	rs.flag = SessionFlagModify
	return nil
}

func (rs *RedisSession) LoadFromRedis() (err error) {
	conn := rs.pool.Get()
	reply, err := conn.Do("GET", rs.id)
	if err != nil {
		return
	}

	data, err := redis.String(reply, err)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(data), &rs.data)
	if err != nil {
		return
	}
	return
}

func (rs *RedisSession) Get (key string) (result interface{}, err error) {
	rs.rwLock.Lock()
	defer rs.rwLock.Unlock()

	if rs.flag == SessionFlagNone {
		err = rs.LoadFromRedis()
		if err != nil {
			return
		}
	}

	result, ok := rs.data[key]
	if !ok {
		err = ErrKeyNotExistInSession
	}
	return
}

func (rs *RedisSession) GetId() string{
	return rs.id
}

func (rs *RedisSession) Del(key string) (err error) {
	rs.rwLock.Lock()
	defer rs.rwLock.Unlock()

	rs.flag = SessionFlagModify
	delete(rs.data, key)
	return nil
}

func (rs *RedisSession) Save()(err error) {
	rs.rwLock.Lock()
	defer rs.rwLock.Unlock()

	if (rs.flag != SessionFlagModify) {
		return
	}

	data, err := json.Marshal(rs.data)
	if err != nil {
		return
	}

	conn := rs.pool.Get()
	_, err = conn.Do("SET", rs.id, string(data))
	if err != nil {
		return
	}
	return
}

func (rs *RedisSession) IsModify() bool {
	return rs.flag == SessionFlagModify
}