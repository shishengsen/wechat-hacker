package cmd

import (
	"github.com/e421083458/gin_scaffold/proto"
	"github.com/e421083458/gin_scaffold/util"
	"github.com/e421083458/gin_scaffold/websocket/client"
	"log"
)

var GroupCmd = Commands{
	"CmdWw355":  createGroup,
	"CmdWw356":  quitGroup,
	"CmdWw357":  dismissGroup,
	"CmdWw358":  renameGroup,
	"CmdWw359":  renameGroupRemark,
	"CmdWw3510": groupMemberList,
	"CmdWw3511": inviteJoinGroup,
	"CmdWw3512": addJoinGroup,
	"CmdWw3515": setGroupOwner,
	"CmdWw3516": setGroupNotice,
	"CmdWw3517": groupContactChange,
	"CmdWw3518": setTopGroup,
	"CmdWw3519": setGroupDisturb,
	"CmdWw3520": setGroupForbidden,
	"CmdWw3521": setGroupConfirm,
	"CmdWw3522": setHideConversation,
	"CmdWw3523": setHasReadConversation,
}

// 创建群聊
func createGroup(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw355{}
	if err := util.DecodeCmd(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 退出群聊
func quitGroup(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw356{}
	if err := util.DecodeCmd(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 解散群聊
func dismissGroup(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw357{}
	if err := util.DecodeCmd(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 修改群名称
func renameGroup(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw358{}
	if err := util.DecodeCmd(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 修改当前账号群名片
func renameGroupRemark(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw359{}
	if err := util.DecodeCmd(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 获取群内成员信息
func groupMemberList(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw3510{}
	if err := util.DecodeCmd(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 邀请他人入群
func inviteJoinGroup(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw3511{}
	if err := util.DecodeCmd(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 添加他人入群
func addJoinGroup(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw3512{}
	if err := util.DecodeCmd(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 设置群主
func setGroupOwner(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw3515{}
	if err := util.DecodeCmd(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 设置群公告
func setGroupNotice(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw3516{}
	if err := util.DecodeCmd(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 将群保存或移出通讯录
func groupContactChange(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw3517{}
	if err := util.DecodeCmd(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 设置/取消置顶
func setTopGroup(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw3518{}
	if err := util.DecodeCmd(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 设置/取消免打扰
func setGroupDisturb(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw3519{}
	if err := util.DecodeCmd(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 设置/取消禁言
func setGroupForbidden(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw3520{}
	if err := util.DecodeCmd(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 设置/取消入群确认
func setGroupConfirm(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw3521{}
	if err := util.DecodeCmd(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 隐藏会话
func setHideConversation(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw3522{}
	if err := util.DecodeCmd(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 标记会话为已读
func setHasReadConversation(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw3523{}
	if err := util.DecodeCmd(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}
