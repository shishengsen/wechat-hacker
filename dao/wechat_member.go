package dao

import (
	"github.com/e421083458/gin_scaffold/public"
	"github.com/gin-gonic/gin"
	"time"
)

type WechatMember struct{
	Id         uint32 `json:"id" orm:"column(id);auto"`
	Wxid uint64 `json:"wxid" orm:"column(wxid)"`
	NickName string `json:"nickName" orm:"column(nick_name)"`
	RealName string `json:"realName" orm:"column(real_name)"`
	Phone string `json:"phone" orm:"column(phone)"`
	Avatar string `json:"avatar" orm:"column(avatar)"`
	Deptname string `json:"deptname" orm:"column(deptname)"`
	Corpname string `json:"corpname" orm:"column(corpname)"`

	CreateAt   time.Time `json:"created_at" orm:"column(create_at);type(datetime)" description:"创建时间"`
	UpdateAt   time.Time `json:"updated_at" orm:"column(update_at);type(datetime)" description:"更新时间"`
	DeletedAt  time.Time `json:"deleted_at" orm:"column(deleted_at);type(datetime)" description:"删除时间"`
}

func (f *WechatMember) TableName() string {
	return "wechat_member"
}

func (f *WechatMember) Del(c *gin.Context, idSlice []string) error {
	err := public.GormPool.SetCtx(public.GetGinTraceContext(c)).Where("id in (?)", idSlice).Delete(&WechatMember{}).Error
	if err != nil {
		return err
	}
	return nil
}

