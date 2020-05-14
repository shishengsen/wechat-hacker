package cmd

import (
	"fmt"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/e421083458/gin_scaffold/util"
	wsClient "github.com/e421083458/gin_scaffold/websocket/client"
	"github.com/e421083458/golang_common/log"
	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"
	"reflect"
)

type Commands map[string]interface{}

var AllCommand Commands

func LoadCommands() {
	AllCommand = Commands{}
	mergeCommands(AllCommand, ContactCmd)
	mergeCommands(AllCommand, GroupCmd)
	mergeCommands(AllCommand, InfoCmd)
	mergeCommands(AllCommand, MsgCmd)
}

func mergeCommands(cmdMap1 Commands, cmdMap2 Commands) {
	for k, v := range cmdMap2 {
		if _, ok := cmdMap2[k]; ok {
			cmdMap1[k] = v
		}
	}
}

func CallCmd(msgData *wsClient.ClientMessage) error {
	fmt.Printf("ready to run cmd: %v\n", msgData)
	cmd, ok := AllCommand[msgData.Cname]
	if !ok {
		return errors.New("invalid command name:" + msgData.Cname)
	}
	fv := reflect.ValueOf(cmd)
	args := []reflect.Value{reflect.ValueOf(msgData)}
	binaryData, err := util.EncodeCmd(msgData.Cmd);
	if err != nil {
		return nil
	}
	// 区分连接者角色-客户消息将cid存储至redis客户队列，终端消息根据cid匹配是否需要推送给客户
	if msgData.Cmd.Cid == "" {
		log.Info("empty cid: %v", msgData)
		return nil
	}
	customerCid := public.CustomerCmdKey + msgData.Cmd.Cid
	if msgData.Role == public.RoleCustomer {
		redC, err := public.NewRedisConn()
		defer redC.Close()
		_, err = redC.Do("set", customerCid, msgData.Wxid)
		if err != nil {
			log.Error(" [redis] set customer cmd fail:clientId:%s customerId: %d error: %v", msgData.ClientConnId, msgData.Wxid, err)
			return nil
		}
		public.SingleMsg(uint64(msgData.Cmd.Wid), binaryData)
	}
	if msgData.Role == public.RoleClient {
		// 根据cid匹配是否为客户指定回调
		redC, err := public.NewRedisConn()
		defer redC.Close()
		customerId, err := redis.Uint64(redC.Do("get", customerCid))
		if err == nil && customerId > 0 {
			public.SingleMsg(customerId, binaryData)
			_, err = redC.Do("del", customerCid)
			if err != nil {
				log.Info("[redis] del customer cmd fail: %s", customerCid)
				return nil
			}
		}
		fv.Call(args)
	}
	return nil
}
