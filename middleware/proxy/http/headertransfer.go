package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-gateway/middleware"
	"strings"
)

func HttpHeaderTransferMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Get("service")
		if !ok {
			middleware.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}

		for _, item := range strings.Split("", ",") {
			items := strings.Split(item, " ")
			if len(items) != 3 {
				continue
			}
			if items[0] == "add" || items[0] == "edit" {
				c.Request.Header.Set(items[1], items[2])
			}
			if items[0] == "del" {
				c.Request.Header.Del(items[1])
			}
		}
		c.Next()
	}
}
