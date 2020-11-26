package middleware

import (
	"errors"
	"fmt"
	"github.com/e421083458/go_gateway/golang_common/lib"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"runtime/debug"
)

func RecoveryMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(string(debug.Stack()))

				log.WithFields(log.Fields{
					"error": fmt.Sprint(err),
					"stack": string(debug.Stack()),
				}).Info("_com_panic")

				if lib.ConfBase.DebugMode != "debug" {
					ResponseError(c, 500, errors.New("内部错误"))
					return
				} else {
					ResponseError(c, 500, errors.New(fmt.Sprint(err)))
					return
				}
			}
		}()
		c.Next()
	}
}
