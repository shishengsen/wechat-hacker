package dao

import (
	"github.com/e421083458/gin_scaffold/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"time"
)

type User struct {
	Id         uint32 `json:"id" orm:"column(id);auto"`
	UserName   string `json:"userName" orm:"column(user_name);"`
	Phone      string `json:"phone" orm:"column(phone);"`
	Password   string `json:"password" orm:"column(password);"`
	Role       public.RoleType    `json:"role" orm:"column(role);"`
	CustomerId int    `json:"customerId" orm:"column(customer_id);"`
	CompanyIds string `json:"companyIds" orm:"column(company_ids);"`
	ApiToken   string `json:"apiToken" orm:"column(api_token);"`
	UpdateAt   time.Time `json:"updated_at" orm:"column(update_at);type(datetime)" description:"更新时间"`
	CreateAt   time.Time `json:"created_at" orm:"column(create_at);type(datetime)" description:"创建时间"`
	DeletedAt  time.Time `json:"deleted_at" orm:"column(deleted_at);type(datetime)" description:"删除时间"`
}



func (f *User) TableName() string {
	return "user"
}

func (f *User) Del(c *gin.Context, idSlice []string) error {
	err := public.GormPool.SetCtx(public.GetGinTraceContext(c)).Where("id in (?)", idSlice).Delete(&User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (f *User) FindByPhone(c *gin.Context, phone string) (*User, error) {
	var user User
	err := public.GormPool.SetCtx(public.GetGinTraceContext(c)).Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (f *User) PageList(c *gin.Context, phone string, pageNo int, pageSize int) ([]*User, int64, error) {
	var user []*User
	var userCount int64
	//limit offset,pagesize
	offset := (pageNo - 1) * pageSize
	query := public.GormPool.SetCtx(public.GetGinTraceContext(c))
	if phone != "" {
		query = query.Where("phone = ?", phone)
	}

	err := query.Limit(pageSize).Offset(offset).Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	errCount := query.Table("user").Count(&userCount).Error
	if errCount != nil {
		return nil, 0, err
	}
	return user, userCount, nil
}

func (f *User) Save(c *gin.Context) error {
	if err := public.GormPool.SetCtx(public.GetGinTraceContext(c)).Save(f).Error; err != nil {
		return err
	}
	return nil
}
