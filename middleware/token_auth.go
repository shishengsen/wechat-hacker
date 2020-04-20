package middleware

import (
	"errors"
	"fmt"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/e421083458/gin_scaffold/util"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleWare(role public.RoleType) gin.HandlerFunc {
	return func(c *gin.Context) {
		isValid := util.ValidateToken(c, role)
		if !isValid {
			ResponseError(c, InternalErrorCode, errors.New(fmt.Sprintf("token is invalid: %v", c.GetHeader("Authroization"))))
			c.Abort()
			return
		}
		c.Next()
	}
}
