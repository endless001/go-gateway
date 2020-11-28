package auth

import (
	"errors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-gateway/internal/middleware"
	"go-gateway/internal/util"
)

func SessionAuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		session := sessions.Default(c)
		if adminInfo, ok := session.Get(util.AdminSessionInfoKey).(string); !ok || adminInfo == "" {
			middleware.ResponseError(c, middleware.InternalErrorCode, errors.New("user not login"))
			c.Abort()
			return
		}
		c.Next()
	}
}
