package dao

import (
	"github.com/e421083458/gin_scaffold/public"
	"github.com/gin-gonic/gin"
	"time"
)

type WechatCustomerList struct{
	Id         uint32 `json:"id" orm:"column(id);auto"`
	OwnerId uint64 `json:"ownerId" orm:"column(owner_id)"`
	Wxid uint64 `json:"wxid" orm:"column(wxid)"`
	NickName string `json:"nickName" orm:"column(nick_name)"`
	Avatar string `json:"avatar" orm:"column(avatar)"`
	Type int8 `json:"type" orm:"column(type)"`
	Gender int8 `json:"gender" orm:"column(gender)"`
	State int8 `json:"state" orm:"column(state)"`
	AddTime   time.Time `json:"addTime" orm:"column(add_time);type(datetime)" description:"添加时间"`
	UpdateAt   time.Time `json:"updated_at" orm:"column(update_at);type(datetime)" description:"更新时间"`
	CreateAt   time.Time `json:"created_at" orm:"column(create_at);type(datetime)" description:"创建时间"`
	DeletedAt  time.Time `json:"deleted_at" orm:"column(deleted_at);type(datetime)" description:"删除时间"`
}

func (f *WechatCustomerList) TableName() string {
	return "wechat_customer_list"
}

func (f *WechatCustomerList) Del(c *gin.Context, idSlice []string) error {
	err := public.GormPool.SetCtx(public.GetGinTraceContext(c)).Where("id in (?)", idSlice).Delete(&WechatCustomerList{}).Error
	if err != nil {
		return err
	}
	return nil
}

