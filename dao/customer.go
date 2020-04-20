package dao

import (
	"github.com/e421083458/gin_scaffold/public"
	"github.com/gin-gonic/gin"
	"time"
)

type Customer struct{
	Id         uint32 `json:"id" orm:"column(id);auto"`
	Name string `json:"name" orm:"column(name)"`
	Status int8 `json:"status" orm:"column(status)"`
	UpdateAt   time.Time `json:"updated_at" orm:"column(update_at);type(datetime)" description:"更新时间"`
	CreateAt   time.Time `json:"created_at" orm:"column(create_at);type(datetime)" description:"创建时间"`
	DeletedAt  time.Time `json:"deleted_at" orm:"column(deleted_at);type(datetime)" description:"删除时间"`
}

func (f *Customer) TableName() string {
	return "customer"
}

func (f *Customer) Del(c *gin.Context, idSlice []string) error {
	err := public.GormPool.SetCtx(public.GetGinTraceContext(c)).Where("id in (?)", idSlice).Delete(&Customer{}).Error
	if err != nil {
		return err
	}
	return nil
}

