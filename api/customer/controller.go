package customer

import (
	"encoding/json"
	"errors"
	"github.com/e421083458/gin_scaffold/middleware"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/e421083458/gin_scaffold/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"reflect"
)

var (
	CookieSession    = "sess"           // session名称
	CookieMaxAge     = 3600 * 24 * 30   // cookie有效时间,一个月
	CookieDomain     = "127.0.0.1" // cookie domain
)

type Controller struct{}

var svr *Service

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
		return
	}
	middleware.ResponseSuccess(ctx, resp)
	return
}

func (c Controller) Test(ctx *gin.Context){
	middleware.ResponseSuccess(ctx, util.ValidateToken(ctx, 2))
	return
	session := sessions.Default(ctx)
	sessData := session.Get("user")
	//dataType := reflect.TypeOf(sessData)
	//weCtx := public.ParserWeContext(ctx)
	//middleware.ResponseSuccess(ctx, weCtx)
	//wUser := weCtx.GetUser()
	//user := &dao.User{
	//	Id: wUser.Id,
	//	Role:wUser.Role,
	//}
	//dt := sessData.(uint8)
	resp := &LoginResp{}
	if err := json.Unmarshal([]byte(reflect.ValueOf(sessData).Interface().(string)), resp); err != nil {
		middleware.ResponseError(ctx, 2002, err)
	}
	session.Delete("user")
	if err := session.Save(); err != nil {
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	middleware.ResponseSuccess(ctx, resp)
	//data := json.Unmarshal([]byte(dt), &LoginResp{})
	//middleware.ResponseSuccess(ctx, data)
	return
}

func(c Controller) Logout(ctx *gin.Context){
	session := sessions.Default(ctx)
	session.Delete("user")
	middleware.ResponseSuccess(ctx, "success")
	return
}