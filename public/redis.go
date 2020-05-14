package public

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"strconv"
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
	CustomerCmdKey = "wechat-hacker-customer-send-cmd"
)

func SingleMsg(wxid uint64, msg []byte) {
	redC, err := NewRedisConn()
	if err != nil {
		log.Printf("[redis error] new redis connect err: %v", err)
		return
	}
	defer redC.Close()
	_, err = redC.Do("lpush", strconv.Itoa(int(wxid))+"-msg-queue", msg)
	if err != nil {
		fmt.Printf("push err: %v", err)
	}
}

func BatchMsg(msg []byte) {
	redC, err := NewRedisConn()
	if err != nil {
		log.Printf("[redis error] new redis connect err: %v", err)
		return
	}
	clientWxids, err := redis.Int64s(redC.Do("smembers", RegisterKey))
	if err != nil {
		log.Printf("[redis error] new redis connect err: %v", err)
		return
	}
	for _, wxid := range clientWxids {
		SingleMsg(uint64(wxid), msg)
		fmt.Printf("client wxid: %v\n", wxid)
	}
}