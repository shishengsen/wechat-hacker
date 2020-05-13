package api

import "github.com/e421083458/gin_scaffold/public"

type Request struct {
}

type RegisterInput struct {
	UserName   string `form:"UserName" json:"userName" validate:"required"`
	Phone      string `form:"Phone" json:"phone" validate:"required"`
	Password   string `form:"Password" json:"password" validate:"required"`
	Role       public.RoleType    `form:"Role" json:"role" validate:"required"`
	CustomerId int    `form:"CustomerId" json:"customerId" validate:"required"`
	CompanyIds string `form:"CompanyIds" json:"companyIds" validate:"required"`
	ApiToken   string `form:"ApiToken" json:"apiToken"`
}