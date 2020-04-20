package dao

import (
	"github.com/e421083458/gin_scaffold/public"
	"github.com/gin-gonic/gin"
	"time"
)

type WechatGroupList struct{
	Id         uint32 `json:"id" orm:"column(id);auto"`
	GroupId uint64 `json:"groupId" orm:"column(group_id)"`
	CustomGroupId string `json:"customGroupId" orm:"column(custom_group_id)"`
	OwnerId uint64 `json:"ownerId" orm:"column(owner_id)"`
	Name string `json:"name" orm:"column(name)"`
	QrCode string `json:"qrCode" orm:"column(qr_code)"`
	MemberIds string `json:"memberIds" orm:"column(member_ids)"`
	CreateAt   time.Time `json:"created_at" orm:"column(create_at);type(datetime)" description:"创建时间"`
	UpdateAt   time.Time `json:"updated_at" orm:"column(update_at);type(datetime)" description:"更新时间"`
	DeletedAt  time.Time `json:"deleted_at" orm:"column(deleted_at);type(datetime)" description:"删除时间"`
}

func (f *WechatGroupList) TableName() string {
	return "wechat_group_list"
}

func (f *WechatGroupList) Del(c *gin.Context, idSlice []string) error {
	err := public.GormPool.SetCtx(public.GetGinTraceContext(c)).Where("id in (?)", idSlice).Delete(&WechatGroupList{}).Error
	if err != nil {
		return err
	}
	return nil
}

