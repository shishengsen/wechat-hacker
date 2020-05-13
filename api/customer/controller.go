package customer

import (
	"encoding/json"
	"errors"
	"github.com/e421083458/gin_scaffold/middleware"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/e421083458/gin_scaffold/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

/**
	客户控制器
 */
type Controller struct{}

var svr *Service

// 登录接口-每次登录会重置token[有效期7天，需用户自行登录获取]
func (c *Controller) Login(ctx *gin.Context) {
	params := &LoginInput{}
	if err := public.BindWithValidate(ctx, params); err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}
	// 密码校验
	user, err := svr.VerifyPassword(ctx, params)
	if err != nil {
		middleware.ResponseError(ctx, 2001, errors.New("账号或密码错误"))
		return
	}
	token, err := util.MakeToken(user.Role)
	if err != nil {
		middleware.ResponseError(ctx, 2001, err)
	}
	user.ApiToken = token.Token
	err = user.Save(ctx); if err !=nil {
		middleware.ResponseError(ctx, 2002, errors.New("更新token失败"))
	}
	resp := &LoginResp{
		UserId:   user.Id,
		UserName: user.UserName,
		Phone:    user.Phone,
		ApiToken: token.Token,
		ExpireTime: token.ExpireTime,
		Role: user.Role,
	}
	session := sessions.Default(ctx)
	sessData, err := json.Marshal( &resp); if err != nil {
		middleware.ResponseError(ctx, 2002, err)
	}
	session.Set("user", string(sessData))
	if err = session.Save(); err != nil {
		middleware.ResponseError(ctx, 2002, err)
	}
	middleware.ResponseSuccess(ctx, resp)
	return
}

// 修改密码
func (c * Controller) ModifyPassword(ctx *gin.Context) {
	params := &ModifyPasswordInput{}
	if err := public.BindWithValidate(ctx, params); err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}
	// 密码校验
	userInput := &LoginInput{
		Username:params.Username,
		Password:params.Password,
	}
	user, err := svr.VerifyPassword(ctx, userInput)
	if err != nil {
		middleware.ResponseError(ctx, 2001, errors.New("账号或密码错误"))
		return
	}
	// 更新密码
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(params.NewPassword), bcrypt.DefaultCost); if err != nil{
		middleware.ResponseError(ctx, 2007, err)
	}
	token, err := util.MakeToken(user.Role)
	if err != nil {
		middleware.ResponseError(ctx, 2001, err)
	}
	user.ApiToken = token.Token
	user.Password = string(hashPwd)
	err = user.Save(ctx); if err !=nil {
		middleware.ResponseError(ctx, 2002, errors.New("修改密码失败"))
	}
	resp := &LoginResp{
		UserId:   user.Id,
		UserName: user.UserName,
		Phone:    user.Phone,
		ApiToken: token.Token,
		ExpireTime: token.ExpireTime,
		Role: user.Role,
	}
	session := sessions.Default(ctx)
	sessData, err := json.Marshal( &resp); if err != nil {
		middleware.ResponseError(ctx, 2002, err)
	}
	session.Set("user", string(sessData))
	if err = session.Save(); err != nil {
		middleware.ResponseError(ctx, 2002, err)
	}
	middleware.ResponseSuccess(ctx, resp)
	return
}


// 注销登录
func(c Controller) Logout(ctx *gin.Context){
	session := sessions.Default(ctx)
	session.Delete("user")
	middleware.ResponseSuccess(ctx, "success")
	return
}

// 获取token
func (c *Controller) GetToken(ctx *gin.Context) {
	token, err := util.MakeToken(public.RoleCustomer)
	if err != nil {
		middleware.ResponseError(ctx, 2001, err)
	}
	middleware.ResponseSuccess(ctx, token)
	return
}