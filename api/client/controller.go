package client

import (
	"github.com/e421083458/gin_scaffold/middleware"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/e421083458/gin_scaffold/util"
	"github.com/e421083458/gin_scaffold/websocket"
	"github.com/gin-gonic/gin"
)

/**
客户端设备控制器
*/
type Controller struct {
}

// 获取token
func (c *Controller) GetToken(ctx *gin.Context) {
	token, err := util.MakeToken(public.RoleClient)
	if err != nil {
		middleware.ResponseError(ctx, 2001, err)
	}
	middleware.ResponseSuccess(ctx, token)
	return
}

// 推送消息
func (c *Controller) Push(ctx *gin.Context) {
	msg := websocket.Message{}
	middleware.ResponseSuccess(ctx, msg)
	return
}

// 同步数据
func (c *Controller) SyncData(ctx *gin.Context) {
	return
}

/*--------------------------------------------用户信息操作-----------------------------------------------*/

// 修改头像
func (c *Controller) AvatarModify(ctx *gin.Context) {
	return
}

// 获取当前用户登录账号二维码
func (c *Controller) GetQrCode(ctx *gin.Context) {
	return
}

// 获取当前用户信息
func (c *Controller) GetUserInfo(ctx *gin.Context) {
	return
}

// 获取所有部门信息
func (c *Controller) GetDepartments(ctx *gin.Context) {
	return
}

// 获取指定部门信息
func (c *Controller) GetDepartment(ctx *gin.Context) {
	return
}

// 获取所有外部客户
func (c *Controller) GetAllCustomer(ctx *gin.Context) {
	return
}

// 获取所有会话列表
func (c *Controller) GetAllChat(ctx *gin.Context) {
	return
}

// 获取所有群会话列表
func (c *Controller) GetAllGroup(ctx *gin.Context) {
	return
}

// 获取所有保存到通讯录的群会话列表
func (c *Controller) GetAllSavedGroup(ctx *gin.Context) {
	return
}


/*--------------------------------------------联系人相关-----------------------------------------------*/
// 修改内部联系人备注
func (c *Controller) InsideContactRemark(ctx *gin.Context){
	return
}

// 修改外部联系人备注
func (c *Controller) OutsideContactRemark(ctx *gin.Context){
	return
}

// 根据手机号搜索联系人
func (c *Controller) SearchContactByPhone(ctx *gin.Context){
	return
}

// 添加联系人
func (c *Controller) AddContact(ctx *gin.Context){
	return
}

// 删除外部联系人
func (c *Controller) DeleteContactOutside(ctx *gin.Context){
	return
}

/*--------------------------------------------群操作相关-----------------------------------------------*/

// 创建群聊
func (c *Controller) CreateGroup(ctx *gin.Context){
	return
}

// 退出群聊
func (c *Controller) QuitGroup(ctx *gin.Context){
	return
}

// 解散群聊
func (c *Controller) DismissGroup(ctx *gin.Context){
	return
}

// 修改群名称
func (c *Controller) RenameGroup(ctx *gin.Context){
	return
}

// 修改当前账号群名片
func (c *Controller) RenameGroupRemark(ctx *gin.Context){
	return
}

// 获取群内成员信息
func (c *Controller) GroupMemberList(ctx *gin.Context){
	return
}

// 邀请他人入群
func (c *Controller) InviteJoinGroup(ctx *gin.Context){
	return
}

// 添加他人入群
func (c *Controller) AddJoinGroup(ctx *gin.Context){
	return
}

// 设置群主
func (c *Controller) SetGroupOwner(ctx *gin.Context){
	return
}

// 设置群公告
func (c *Controller) SetGroupNotice(ctx *gin.Context){
	return
}


// 将群保存或移出通讯录
func (c *Controller) GroupContactChange(ctx *gin.Context){
	return
}

// 设置/取消置顶
func (c *Controller) SetTopGroup(ctx *gin.Context){
	return
}

// 设置/取消免打扰
func (c *Controller) SetGroupDisturb(ctx *gin.Context){
	return
}

// 设置/取消禁言
func (c *Controller) SetGroupForbidden(ctx *gin.Context){
	return
}

// 设置/取消入群确认
func (c *Controller) SetGroupConfirm(ctx *gin.Context){
	return
}


/*--------------------------------------------消息发送-----------------------------------------------*/
// 发送消息[文本/图片/语音/视频/文件/链接/地理位置/小程序]
func (c *Controller) MsgSend(ctx *gin.Context) {
	return
}


