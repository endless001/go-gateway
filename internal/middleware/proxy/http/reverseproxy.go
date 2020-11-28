package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-gateway/internal/middleware"
)

func HttpReverseProxyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Get("service")
		if !ok {
			middleware.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}

	}
}
