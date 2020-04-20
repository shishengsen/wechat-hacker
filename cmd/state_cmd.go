package cmd

import (
	"fmt"
	"github.com/e421083458/gin_scaffold/public"
	wsClient "github.com/e421083458/gin_scaffold/websocket/client"
	"log"
)

var StateCmd = Commands{
	10001: StateConfirmConnect,
	10002: StateDisconnect,
}

func StateConfirmConnect(data *wsClient.ClientMessage) {
	redC, err := public.NewRedisConn()
	defer redC.Close()
	_, err = redC.Do("sadd", public.RegisterKey, data.Wxid)
	if err != nil {
		log.Printf(" [redis] set register fail:%s %d %v", data.ClientConnId, data.Wxid, err)
		return
	}
	// 未连接池移除
	_, _ = redC.Do("srem", public.UnRegisterKey, data.Wxid)
	// ws连接映射
	_, _ = redC.Do("set", data.Wxid, data.ClientConnId)
	_, _ = redC.Do("set", data.ClientConnId, data.Wxid)
}

func StateDisconnect(data *wsClient.ClientMessage) {
	redC, err :=  public.NewRedisConn()
	defer redC.Close()
	if err != nil {
		fmt.Println("Connect to redis error", err)
	}
	// 从连接池中移除
	_, _ = redC.Do("srem", public.RegisterKey, data.Wxid)
	// 进入未连接池
	_, _ = redC.Do("sadd", public.UnRegisterKey, data.Wxid)
	// 移除连接映射
	_, _ = redC.Do("del", data.Wxid)
	_, _ = redC.Do("del", data.ClientConnId)
}

//

//func MyMissionMethod(a string){
//	fmt.Println("hello, world, this is my mission.")
//	fmt.Printf("and this is my params: %s \n", a)
//}

//
//func main() {
//	mission := MyMissionMethod
//	CallMethod(mission)
//}
