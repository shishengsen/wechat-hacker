package cmd

import (
	"github.com/e421083458/gin_scaffold/proto"
	"github.com/e421083458/gin_scaffold/util"
	"github.com/e421083458/gin_scaffold/websocket/client"
	"log"
)

var MsgCmd = Commands{
	"CmdWw361":  contactMsgSync,
	"CmdWw362":  systemMsgSync,
	"CmdWw363":  sendTextMsg,
	"CmdWw364":  sendImageMsg,
	"CmdWw365":  sendVoiceMsg,
	"CmdWw366":  sendVideoMsg,
	"CmdWw367":  sendFileMsg,
	"CmdWw368":  sendLinkMsg,
	"CmdWw369":  sendLocationMsg,
	"CmdWw3610": sendLiteAppMsg,
}

// 新消息同步
func contactMsgSync(msg *wsClient.ClientMessage) {
	data := &proto.CmdWw361{}
	if err := util.DecodeCmd(msg.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 系统消息同步
func systemMsgSync(msg *wsClient.ClientMessage) {
	data := &proto.CmdWw362{}
	if err := util.DecodeCmd(msg.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 发送文本消息
func sendTextMsg(msg *wsClient.ClientMessage) {
	data := &proto.CmdWw363{}
	if err := util.DecodeCmd(msg.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 发送图片消息
func sendImageMsg(msg *wsClient.ClientMessage) {
	data := &proto.CmdWw364{}
	if err := util.DecodeCmd(msg.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 发送语音消息(不能超过60s)
func sendVoiceMsg(msg *wsClient.ClientMessage) {
	data := &proto.CmdWw365{}
	if err := util.DecodeCmd(msg.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 发送视频消息
func sendVideoMsg(msg *wsClient.ClientMessage) {
	data := &proto.CmdWw366{}
	if err := util.DecodeCmd(msg.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 发送文件消息
func sendFileMsg(msg *wsClient.ClientMessage) {
	data := &proto.CmdWw367{}
	if err := util.DecodeCmd(msg.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 发送链接消息
func sendLinkMsg(msg *wsClient.ClientMessage) {
	data := &proto.CmdWw368{}
	if err := util.DecodeCmd(msg.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 发送地理位置消息
func sendLocationMsg(msg *wsClient.ClientMessage) {
	data := &proto.CmdWw369{}
	if err := util.DecodeCmd(msg.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 发送小程序消息
func sendLiteAppMsg(msg *wsClient.ClientMessage) {
	data := &proto.CmdWw3610{}
	if err := util.DecodeCmd(msg.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}
