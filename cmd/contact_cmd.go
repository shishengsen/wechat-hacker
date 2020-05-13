package cmd

import (
	"github.com/e421083458/gin_scaffold/proto"
	"github.com/e421083458/gin_scaffold/util"
	"github.com/e421083458/gin_scaffold/websocket/client"
	"log"
)

var ContactCmd = Commands{
	"CmdWw342": insideContactRemark,
	"CmdWw341": outsideContactRemark,
	"CmdWw343": searchContactByPhone,
	"CmdWw344": addContact,
	"CmdWw347": deleteContactOutside,
}
// 修改内部联系人备注
func insideContactRemark(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw342{}
	if err := util.DecodeCmd(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 修改外部联系人备注
func outsideContactRemark(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw341{}
	if err := util.DecodeCmd(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 根据手机号搜索联系人
func searchContactByPhone(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw343{}
	if err := util.DecodeCmd(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 添加联系人
func addContact(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw344{}
	if err := util.DecodeCmd(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 删除外部联系人
func deleteContactOutside(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw347{}
	if err := util.DecodeCmd(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}
