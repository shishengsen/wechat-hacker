package cmd

import (
	"fmt"
	"github.com/e421083458/gin_scaffold/dao"
	"github.com/e421083458/gin_scaffold/proto"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/e421083458/gin_scaffold/websocket/client"
	proto2 "github.com/gogo/protobuf/proto"
	"log"
)

var InfoCmd = Commands{
	"CmdWw331":  getUserInfo,
	"CmdWw332":  avatarModify,
	"CmdWw333":  getQrCode,
	"CmdWw348":  getDepartments,
	"CmdWw349":  getDepartment,
	"CmdWw3411": getAllCustomer,
	"CmdWw352":  getAllChat,
	"CmdWw353":  getAllGroup,
	"CmdWw354":  getAllSavedGroup,
	"CmdWw322":  disconnect,
}

// 用户信息上报-获取当前用户信息[连接确认]
func getUserInfo(cmd *wsClient.ClientMessage) {
	redC, err := public.NewRedisConn()
	defer redC.Close()
	_, err = redC.Do("sadd", public.RegisterKey, cmd.Wxid)
	if err != nil {
		log.Printf(" [redis] set register fail:%s %d %v", cmd.ClientConnId, cmd.Wxid, err)
		return
	}
	// 未连接池移除
	_, _ = redC.Do("srem", public.UnRegisterKey, cmd.Wxid)
	// ws连接映射
	_, _ = redC.Do("set", cmd.Wxid, cmd.ClientConnId)
	_, _ = redC.Do("set", cmd.ClientConnId, cmd.Wxid)
	// 用户数据
	data := &proto.CmdWw331{}
	if err := proto2.Unmarshal(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	memberDao := &dao.WechatMember{}
	_, err = memberDao.FindByWxid(nil, data.Data.Id)
	if err != nil {
		// 无用户则新增用户
		memberDao.Wxid = data.Data.Id
		memberDao.Phone = data.Data.Phone
		memberDao.Acctid = data.Data.Acctid
		memberDao.Avatar = data.Data.AvatarUrl
		memberDao.Gender = data.Data.Gender
		memberDao.Name = data.Data.Name
		memberDao.NickName = data.Data.Nickname
		memberDao.Alias = data.Data.Alias
		memberDao.Gender = data.Data.Gender
		memberDao.Phone = data.Data.Phone
		memberDao.Mobile = data.Data.Mobile
		memberDao.Acctid = data.Data.Acctid
		memberDao.CorpId = data.Data.CorpId
		memberDao.CorpName = data.Data.CorpName
		memberDao.CorpFname = data.Data.CorpFName
		memberDao.CorpAddress = data.Data.CorpAddress
		memberDao.Job = data.Data.Job
		err = memberDao.Save(nil)
		if err != nil {
			log.Printf(" [redis] save member data error: %v", err)
		}
	}
	return
}

func disconnect(cmd *wsClient.ClientMessage) {
	redC, err := public.NewRedisConn()
	defer redC.Close()
	if err != nil {
		fmt.Println("Connect to redis error", err)
	}
	// 从连接池中移除
	_, _ = redC.Do("srem", public.RegisterKey, cmd.Wxid)
	// 进入未连接池
	_, _ = redC.Do("sadd", public.UnRegisterKey, cmd.Wxid)
	// 移除连接映射
	_, _ = redC.Do("del", cmd.Wxid)
	_, _ = redC.Do("del", cmd.ClientConnId)
}

// 修改头像
func avatarModify(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw332{}
	if err := proto2.Unmarshal(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 获取当前用户登录账号二维码
func getQrCode(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw333{}
	if err := proto2.Unmarshal(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 获取所有部门信息
func getDepartments(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw348{}
	if err := proto2.Unmarshal(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 获取指定部门信息
func getDepartment(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw349{}
	if err := proto2.Unmarshal(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 获取所有外部客户
func getAllCustomer(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw3411{}
	if err := proto2.Unmarshal(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 获取所有会话列表
func getAllChat(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw352{}
	if err := proto2.Unmarshal(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 获取所有群会话列表
func getAllGroup(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw353{}
	if err := proto2.Unmarshal(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}

// 获取所有保存到通讯录的群会话列表
func getAllSavedGroup(cmd *wsClient.ClientMessage) {
	data := &proto.CmdWw354{}
	if err := proto2.Unmarshal(cmd.Cmd.Data, data); err != nil {
		log.Printf(" [redis] decode cmd failed: %v", err)
	}
	return
}
