package middleware

import (
	"github.com/e421083458/gin_scaffold/public"
	"github.com/gin-gonic/gin"
)

func RegisterGlobalParams() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 注册自定义上下文
		registerContext(ctx)

		ctx.Next()
	}
}


func registerContext(ctx *gin.Context) {
	ctx.Set(public.WeContextKey, public.NewWeContext())
}
