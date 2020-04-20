package dao

import (
	"github.com/e421083458/gin_scaffold/public"
	"github.com/gin-gonic/gin"
	"time"
)

type WechatGroupStaticMessage struct{
	Id         uint32 `json:"id" orm:"column(id);auto"`
	Msgid uint64 `json:"msgid" orm:"column(msgid)"`
	Data string `json:"data" orm:"column(data)"`
	OriginUrl string `json:"originUrl" orm:"column(origin_url)"`
	StorageUrl string `json:"storageUrl" orm:"column(storage_url)"`
	FromWxid uint64 `json:"fromWxid" orm:"column(from_wxid)"`
	ToWxid uint64 `json:"toWxid" orm:"column(to_wxid)"`
	Status int8 `json:"statue" orm:"column(status)"`
	CreateAt   time.Time `json:"created_at" orm:"column(create_at);type(datetime)" description:"创建时间"`
	UpdateAt   time.Time `json:"updated_at" orm:"column(update_at);type(datetime)" description:"更新时间"`
	DeletedAt  time.Time `json:"deleted_at" orm:"column(deleted_at);type(datetime)" description:"删除时间"`
}

func (f *WechatGroupStaticMessage) TableName() string {
	return "wechat_group_static_message"
}

func (f *WechatGroupStaticMessage) Del(c *gin.Context, idSlice []string) error {
	err := public.GormPool.SetCtx(public.GetGinTraceContext(c)).Where("id in (?)", idSlice).Delete(&WechatGroupStaticMessage{}).Error
	if err != nil {
		return err
	}
	return nil
}

