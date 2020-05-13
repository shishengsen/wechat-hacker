package router

import (
	"github.com/e421083458/gin_scaffold/api/manager"
	"github.com/e421083458/gin_scaffold/middleware"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/gin-gonic/gin"
)

func registerManager(router *gin.Engine) {
	// manager
	managerController := &manager.Controller{}
	managerNormalGroup := router.Group("/manager")
	managerNormalGroup.Use(middleware.RequestLog())
	managerNormalGroup.GET("/getUploadToken", managerController.GetUploadToken)
	managerNormalGroup.GET("/getToken", managerController.GetToken)
	managerAuthGroup := router.Group("/manager")
	managerAuthGroup.Use(middleware.TokenAuthMiddleWare(public.RoleManager),)
	{
		managerAuthGroup.POST("/createCustomer", managerController.CreateCustomer)
	}
}