package router

import (
	"encoding/gob"
	"github.com/e421083458/gin_scaffold/api/customer"
	"github.com/e421083458/gin_scaffold/middleware"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

type User struct {
	// 用户ID
	Id uint32 `json:"id"`
	// 姓名
	UserName string `json:"userName"`
	// 手机号
	Phone string `json:"phone"`
	// 角色
	Role public.RoleType `json:"role"`
}

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	//写入gin日志
	gin.DisableConsoleColor()
	gob.Register(User{})
	gob.Register(customer.LoginResp{})
	router := gin.Default()
	router.Use(middleware.RegisterGlobalParams())
	router.Use(middleware.RequestLog())
	store, _ := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte("secret"))
	router.Use(sessions.Sessions("sess", store))
	router.Use(middlewares...)
	registerClient(router)
	registerCustomer(router)
	registerManager(router)
	return router
}
