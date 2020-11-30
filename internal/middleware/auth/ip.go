package auth

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gateway/internal/conf"
	"go-gateway/internal/middleware"
)

func IpAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isMatched := false
		for _, host := range conf.Conf.Server.AllowIp {
			if c.ClientIP() == host {
				isMatched = true
			}
		}
		if !isMatched {
			middleware.ResponseError(c, middleware.InternalErrorCode, errors.New(fmt.Sprintf("%v,not in iplist", c.ClientIP())))
			c.Abort()
			return
		}
		c.Next()
	}
}
