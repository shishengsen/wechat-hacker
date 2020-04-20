package public

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func NewRedisConn() (redis.Conn, error) {
	redC, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return nil, err
	}
	return redC, nil
}

const (
	RegisterKey   = "wechat-hacker-clients-register"
	UnRegisterKey = "wechat-hacker-clients-unregister"
)