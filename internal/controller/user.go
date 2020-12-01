package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-gateway/internal/ent"
	"go-gateway/internal/middleware"
	"go-gateway/internal/middleware/auth"
	"go-gateway/internal/service"
	"go-gateway/internal/util"
)

type Login struct {
	UserName string `form:"username" json:"username" binding:"required"`

	Password string `form:"password" json:"password" binding:"required"`
}

type UserController struct {
	us *service.UserService
}

func UserRegister(group *gin.RouterGroup) {
	c := &UserController{
		us: new(service.UserService),
	}
	group.POST("/login", c.UserLogin)
	group.GET("/logout", auth.SessionAuthMiddleware(), c.UserSignout)
	group.GET("/userinfo", auth.SessionAuthMiddleware(), c.UserInfo)
}

func (u *UserController) UserLogin(c *gin.Context) {
	var login Login

	if err := c.ShouldBind(&login); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	user, _ := u.us.CheckUser(c, login.UserName, login.Password)
	c.AbortWithStatusJSON(200, user)

}

func (u *UserController) UserInfo(c *gin.Context) {
	sess := sessions.Default(c)
	sessInfo := sess.Get(util.AdminSessionInfoKey)
	userInfo := &ent.User{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), userInfo); err != nil {
		middleware.ResponseError(c, 200, err)
		return
	}
	middleware.ResponseSuccess(c, userInfo)
}

func (u *UserController) UserSignout(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Delete(util.AdminSessionInfoKey)
	sess.Save()
	middleware.ResponseSuccess(c, "ok")
}
