package client

import (
	wsClient "github.com/e421083458/gin_scaffold/websocket/client"
)

type Service struct {

}


// 给单个客户端发送消息
func SendSingleMsg(wxid uint64, msg *wsClient.ClientMessage) {

}

// 给多个客户端发送消息
func SendCommonMsg(wxids []uint64,msg *wsClient.ClientMessage){

}