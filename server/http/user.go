package http

import (
	"encoding/json"
	"fmt"
	"github.com/e421083458/go_gateway/public"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-gateway/ent"
	"go-gateway/middleware"
)

func UserLogin(c *gin.Context) {

}

func UserInfo(c *gin.Context) {
	sess := sessions.Default(c)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	userInfo := &ent.User{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), userInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	//middleware.ResponseSuccess(c, userInfo)
}

func UserSignout(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Delete(public.AdminSessionInfoKey)
	sess.Save()
	//middleware.ResponseSuccess(c, "ok")
}
