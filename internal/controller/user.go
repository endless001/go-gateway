package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-gateway/internal/ent"
	"go-gateway/internal/service"
	"go-gateway/internal/util"
)

type UserController struct {
	us *service.UserService
}

func UserRegister(group *gin.RouterGroup) {
	c := &UserController{
		us: new(service.UserService),
	}
	group.POST("/login", c.UserLogin)
	group.GET("/logout", c.UserSignout)
}

func (u *UserController) UserLogin(c *gin.Context) {
	user, _ := u.us.CheckUser(c, "123456", "1234")
	c.AbortWithStatusJSON(200, user)
}

func (u *UserController) UserInfo(c *gin.Context) {
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

func (u *UserController) UserSignout(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Delete(util.AdminSessionInfoKey)
	sess.Save()
	c.AbortWithStatusJSON(200, "ok")
	//middleware.ResponseSuccess(c, "ok")
}
