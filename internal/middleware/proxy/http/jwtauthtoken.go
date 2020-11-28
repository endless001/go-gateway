package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-gateway/internal/middleware"
	"go-gateway/internal/util"
	"strings"
)

func HttpJwtAuthTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Get("service")
		if !ok {
			middleware.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}
		token := strings.ReplaceAll(c.GetHeader("Authorization"), "Bearer ", "")
		//appMatched := false
		if token != "" {
			_, err := util.JwtDecode(token)
			if err != nil {
				middleware.ResponseError(c, 2002, err)
				c.Abort()
				return
			}
		}

	}
}
