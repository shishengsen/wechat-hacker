package router

import (
	"github.com/e421083458/gin_scaffold/api/client"
	"github.com/e421083458/gin_scaffold/middleware"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/gin-gonic/gin"
)

func registerClient(router *gin.Engine) {
	// client
	clientController := &client.Controller{}
	clientNormalGroup := router.Group("/client")
	clientNormalGroup.GET("/getToken", clientController.GetToken)
	clientNormalGroup.Use(middleware.TokenAuthMiddleWare(public.RoleClient))
	clientNormalGroup.GET("/push", clientController.Push)
	clientAuthGroup := router.Group("/client")
	clientAuthGroup.Use(middleware.TokenAuthMiddleWare(public.RoleClient),)
	{

	}
}
