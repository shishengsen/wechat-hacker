package router

import (
	"github.com/e421083458/gin_scaffold/api/customer"
	"github.com/e421083458/gin_scaffold/middleware"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/gin-gonic/gin"
)

func registerCustomer(router *gin.Engine) {
	//customer
	customerController := &customer.Controller{}
	customerNormalGroup := router.Group("/customer")
	customerNormalGroup.POST("/login", customerController.Login)
	customerNormalGroup.POST("/modifyPassword", customerController.ModifyPassword)
	customerNormalGroup.GET("/getToken", customerController.GetToken)
	customerAuthGroup := router.Group("/customer")

	customerAuthGroup.Use(middleware.TokenAuthMiddleWare(public.RoleCustomer),)
	{

	}
}
