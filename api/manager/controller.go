package manager

import (
	"github.com/e421083458/gin_scaffold/api"
	"github.com/e421083458/gin_scaffold/middleware"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/e421083458/gin_scaffold/util"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

/**
	管理角色控制器
 */
type Controller struct {
}

var svr *Service

// 创建客户
func (c *Controller) CreateCustomer(ctx *gin.Context) {
	params := &api.RegisterInput{}
	if err := public.BindWithValidate(ctx, params); err != nil {
		middleware.ResponseError(ctx, 2001, err)
	}
	params.Role = public.RoleCustomer
	token, err := svr.CreateCustomer(params, ctx);if err != nil || token.Token == "" {
		middleware.ResponseError(ctx, 2001, err)
	}
	middleware.ResponseSuccess(ctx, token)
	return
}

// 获取用户token
func (c *Controller) GetToken(ctx *gin.Context) {
	token, err := util.MakeToken(public.RoleManager)
	if err != nil {
		middleware.ResponseError(ctx, 2001, err)
	}
	middleware.ResponseSuccess(ctx, token)
	return
}

// 获取七牛上传token
func (c *Controller) GetUploadToken(ctx *gin.Context) {
	bucket := "w-pub"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	accessKey := "ca2R6rWA7LN06mvEoW4ffyosac_MPdCIOQJmnluX"
	secretKey := "psDfwGy2g74PuDGI6caONnO8SbBTw89oesypy7oh"
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	middleware.ResponseSuccess(ctx, upToken)
	return
}
