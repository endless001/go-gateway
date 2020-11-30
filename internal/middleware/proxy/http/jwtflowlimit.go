package http

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gateway/internal/flow/handler"
	"go-gateway/internal/middleware"
	"go-gateway/internal/util"
)

func HttpJwtFlowLimtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Get("app")
		if !ok {
			c.Next()
			return
		}
		if true {
			clientLimiter, err := handler.FlowLimiterHandler.GetLimiter(
				util.FlowAppPrefix+"_"+c.ClientIP(),
				float64(0))
			if err != nil {
				middleware.ResponseError(c, 5001, err)
				c.Abort()
				return
			}
			if !clientLimiter.Allow() {
				middleware.ResponseError(c, 5002, errors.New(fmt.Sprintf("%v flow limt %v", c.ClientIP(), 0)))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
