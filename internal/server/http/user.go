package http

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-gateway/internal/ent"
	"go-gateway/internal/middleware"
	"go-gateway/internal/util"
)

func UserLogin(c *gin.Context) {

}

func UserInfo(c *gin.Context) {
	sess := sessions.Default(c)
	sessInfo := sess.Get(util.AdminSessionInfoKey)
	userInfo := &ent.User{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), userInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	middleware.ResponseSuccess(c, userInfo)
}

func UserSignout(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Delete(util.AdminSessionInfoKey)
	sess.Save()
	middleware.ResponseSuccess(c, "ok")
}
