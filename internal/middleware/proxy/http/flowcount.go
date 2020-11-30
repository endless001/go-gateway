package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-gateway/internal/flow/handler"
	"go-gateway/internal/middleware"
	"go-gateway/internal/util"
)

func HttpFlowCountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Get("service")
		if !ok {
			middleware.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}
		totalCounter, err := handler.FlowCounterHandler.GetCounter(util.FlowTotal)
		if err != nil {
			middleware.ResponseError(c, 4001, err)
			c.Abort()
			return
		}
		totalCounter.Increase()

		serviceCounter, err := handler.FlowCounterHandler.GetCounter(util.FlowServicePrefix)
		if err != nil {
			middleware.ResponseError(c, 4001, err)
			c.Abort()
			return
		}
		serviceCounter.Increase()
		c.Next()
	}
}
