package dao

import (
	"github.com/e421083458/gin_scaffold/public"
	"github.com/gin-gonic/gin"
	"time"
)

type WechatMember struct{
	Id         uint32 `json:"id" orm:"column(id);auto"`
	Wxid int64 `json:"wxid" orm:"column(wxid)"`
	Name string `json:"name" orm:"column(name)"`
	NickName string `json:"nickName" orm:"column(nick_name)"`
	Alias string `json:"alias" orm:"column(alias)"`
	Avatar string `json:"avatar" orm:"column(avatar)"`
	Gender int32 `json:"gender" orm:"column(gender)"`
	Phone string `json:"phone" orm:"column(phone)"`
	Mobile string `json:"mobile" orm:"column(mobile)"`
	Acctid string `json:"acctid" orm:"column(acctid)"`
	CorpId int64 `json:"corpId" orm:"column(corp_id)"`
	CorpName string `json:"corpName" orm:"column(corp_name)"`
	CorpFname string `json:"corpFname" orm:"column(corp_fname)"`
	CorpAddress string `json:"corpAddress" orm:"column(corp_address)"`
	Job string `json:"job" orm:"column(job)"`
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

func (f *WechatMember) FindByWxid(c *gin.Context, wxid int64) (*WechatMember, error) {
	var member WechatMember
	err := public.GormPool.SetCtx(public.GetGinTraceContext(c)).Where("wxid = ?", wxid).First(&member).Error
	if err != nil {
		return nil, err
	}
	return &member, nil
}


func (f *WechatMember) Save(c *gin.Context) error {
	if err := public.GormPool.SetCtx(public.GetGinTraceContext(c)).Save(f).Error; err != nil {
		return err
	}
	return nil
}