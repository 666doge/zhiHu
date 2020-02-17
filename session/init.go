package session

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

var RedisClient *redis.Client

func Init(addr string, options ...string) (err error){
	RedisClient = redis.NewClient(&redis.Options{
		Addr: addr,
		Password: options[0],
		DB: 0,  // use default DB
	})

	_, err = RedisClient.Ping().Result()
	if err == nil {
		fmt.Println("Initialize redis OK!!!")
	}
	return
}
