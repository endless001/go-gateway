package http

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-gateway/internal/ent"
	"go-gateway/internal/util"
)

func UserLogin(c *gin.Context) {

}

func UserInfo(c *gin.Context) {
	sess := sessions.Default(c)
	sessInfo := sess.Get(util.AdminSessionInfoKey)
	userInfo := &ent.User{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), userInfo); err != nil {
		c.AbortWithError(200, err)
		return
	}
	c.AbortWithStatusJSON(200, userInfo)
	//middleware.ResponseSuccess(c, userInfo)
}

func UserSignout(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Delete(util.AdminSessionInfoKey)
	sess.Save()
	c.AbortWithStatusJSON(200, "ok")
	//middleware.ResponseSuccess(c, "ok")
}
