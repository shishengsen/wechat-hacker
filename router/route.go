package router

import (
	"encoding/gob"
	"github.com/e421083458/gin_scaffold/api/client"
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
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//gin.DefaultErrorWriter = io.MultiWriter(f)
	gob.Register(User{})
	gob.Register(customer.LoginResp{})
	router := gin.Default()
	store, _ := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte("secret"))
	router.Use(sessions.Sessions("sess", store))
	router.Use(middlewares...)


	//demo
	//v1 := router.Group("/demo")
	//v1.Use(middleware.RecoveryMiddleware(), middleware.RequestLog(), middleware.IPAuthMiddleware(), middleware.TranslationMiddleware())
	//{
	//	api.DemoRegister(v1)
	//}

	//api
	//store := sessions.NewCookieStore([]byte("secret"))
	//apiNormalGroup := router.Group("/api")
	//apiController:=&api.Api{}
	//apiNormalGroup.Use(
	//	sessions.Sessions("mysession", store),
	//	middleware.RecoveryMiddleware(),
	//	middleware.RequestLog(),
	//	middleware.TranslationMiddleware())
	//apiNormalGroup.POST("/login",apiController.Login)
	//apiNormalGroup.GET("/loginout",apiController.LoginOut)


	//apiAuthGroup := router.Group("/api")
	//apiAuthGroup.Use(
	//	sessions.Sessions("mysession", store),
	//	middleware.RecoveryMiddleware(),
	//	middleware.RequestLog(),
	//	middleware.SessionAuthMiddleware(),
	//	middleware.TranslationMiddleware())
	//apiAuthGroup.GET("/user/listpage", apiController.ListPage)

	//customer
	customerController := &customer.Controller{}
	customerNormalGroup := router.Group("/customer")
	customerNormalGroup.POST("/login", customerController.Login)
	customerNormalGroup.GET("/test", customerController.Test)

	// client
	clientController := &client.Controller{}
	clientNormalGroup := router.Group("/client")
	clientNormalGroup.GET("/getToken", clientController.GetToken)
	clientNormalGroup.Use(middleware.TokenAuthMiddleWare(public.RoleClient))
	clientNormalGroup.GET("/push", clientController.Push)


	return router
}