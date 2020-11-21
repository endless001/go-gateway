package middleware

import (
	"errors"
	"fmt"
	"github.com/e421083458/go_gateway/golang_common/lib"
	"github.com/e421083458/go_gateway/public"
	"github.com/gin-gonic/gin"
	"runtime/debug"
)

func RecoveryMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err :=recover();err !=nil{
				fmt.Println(string(debug.Stack()))
				public.ComLogNotice(c,"_com_panic",map[string]interface{}{
					"error": fmt.Sprint(err),
					"stack": string(debug.Stack()),
				})
				if lib.ConfBase.DebugMode != "debug"{
					ResponseError(c,500,errors.New("内部错误"))
					return
				}else{
					ResponseError(c,500,errors.New(fmt.Sprint(err)))
					return
				}
			}
		}()
	    c.Next()
	}
}
