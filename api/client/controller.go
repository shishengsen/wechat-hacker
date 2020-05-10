package client

import (
	"github.com/e421083458/gin_scaffold/middleware"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/e421083458/gin_scaffold/util"
	"github.com/e421083458/gin_scaffold/websocket"
	"github.com/gin-gonic/gin"
)

/**
	客户端设备控制器
 */
type Controller struct {

}

// 获取token
func(c *Controller) GetToken(ctx *gin.Context) {
	token, err := util.MakeToken(public.RoleClient)
	if err != nil {
		middleware.ResponseError(ctx, 2001, err)
	}
	middleware.ResponseSuccess(ctx, token)
	return
}

// 推送消息
func(c *Controller) Push(ctx *gin.Context) {
	msg := websocket.Message{}
	middleware.ResponseSuccess(ctx, msg)
	return
}

// 同步数据
func (c *Controller) SyncData(ctx *gin.Context){
	return
}

// 发送消息
func (c *Controller) MsgSend(ctx *gin.Context){
	return
}


