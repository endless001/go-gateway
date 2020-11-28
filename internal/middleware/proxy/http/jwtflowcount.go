package http

import (
	"errors"
	"fmt"
	"github.com/e421083458/go_gateway/public"
	"github.com/gin-gonic/gin"
	"go-gateway/internal/middleware"
)

func HttpJwtFlowCountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Get("app")
		if !ok {
			c.Next()
			return
		}
		appCounter, err := public.FlowCounterHandler.GetCounter(public.FlowServicePrefix)
		if err != nil {
			middleware.ResponseError(c, 2002, err)
			c.Abort()
			return
		}
		appCounter.Increase()

		if appCounter.TotalCount > 1 {
			middleware.ResponseError(c, 2003, errors.New(fmt.Sprintf("租户日请求量限流 limit:%v current%v", 0, appCounter.TotalCount)))
			c.Abort()
			return
		}
		c.Next()
	}
}
