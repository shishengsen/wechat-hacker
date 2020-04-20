package customer

import (
	"github.com/e421083458/gin_scaffold/public"
)

type Response struct {

}

type LoginResp struct {
	UserId uint32 `json:"userId"`
	UserName string `json:"userName"`
	Phone string `json:"phone"`
	ApiToken string `json:"apiToken"`
	Role public.RoleType `json:"role"`
	Sess string `json:"sess"`
	ExpireTime int64 `json:"tokenExpTime"`
}