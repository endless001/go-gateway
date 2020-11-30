package http

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gateway/internal/middleware"
	"go-gateway/internal/util"
)

func HttpWhiteListMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Get("service")
		if !ok {
			middleware.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}
		iplist := []string{}
		if !util.InStringSlice(iplist, c.ClientIP()) {
			middleware.ResponseError(c, 3001, errors.New(fmt.Sprintf("%s not in white ip list", c.ClientIP())))
			c.Abort()
			return
		}

		c.Next()

	}
}
