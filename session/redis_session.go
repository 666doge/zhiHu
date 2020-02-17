package session

import (
	"fmt"
)

func getReidsKey(key string) string{
	return fmt.Sprintf("user_session_%v", key)
}

func Set(key string, filed string, value interface{}) error {
	key = getReidsKey(key)
	_, err := RedisClient.Do("hset", key, filed, value).Result()
	return err
}

func Get(key string, filed string) (value interface{}, err error) {
	key = getReidsKey(key)
	value, err = RedisClient.Do("hget", key, filed).Result()
	return
}
