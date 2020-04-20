package dao

import (
	"github.com/e421083458/gin_scaffold/public"
	"github.com/gin-gonic/gin"
	"time"
)

type CustomerCompany struct{
	Id         uint32 `json:"id" orm:"column(id);auto"`
	Name string `json:"name" orm:"column(name)"`
	CorpId string `json:"corpId" orm:"column(corp_id)"`
	AppSecret string `json:"appSecret" orm:"column(app_secret)"`
	AppToken string `json:"appToken" orm:"column(app_token)"`
	CustomerId int32 `json:"customerId" orm:"column(customer_id)"`
	Status int8 `json:"status" orm:"column(status)"`
	UpdateAt   time.Time `json:"updated_at" orm:"column(update_at);type(datetime)" description:"更新时间"`
	CreateAt   time.Time `json:"created_at" orm:"column(create_at);type(datetime)" description:"创建时间"`
	DeletedAt  time.Time `json:"deleted_at" orm:"column(deleted_at);type(datetime)" description:"删除时间"`
}

func (f *CustomerCompany) TableName() string {
	return "CustomerCompany"
}

func (f *CustomerCompany) Del(c *gin.Context, idSlice []string) error {
	err := public.GormPool.SetCtx(public.GetGinTraceContext(c)).Where("id in (?)", idSlice).Delete(&CustomerCompany{}).Error
	if err != nil {
		return err
	}
	return nil
}

