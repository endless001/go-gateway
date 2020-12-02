package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-gateway/internal/middleware"
	"go-gateway/internal/middleware/auth"
	"go-gateway/internal/service"
	"go-gateway/internal/util"
	"time"
)

type Login struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type ChangePassword struct {
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"`
}

type SessionInfo struct {
	ID        int       `json:"id"`
	UserName  string    `json:"user_name"`
	LoginTime time.Time `json:"login_time"`
}

type LoginResult struct {
	Token string `json:"token" form:"token" comment:"token" example:"token" validate:""`
}

type UserInfoResult struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	LoginTime    time.Time `json:"login_time"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	Roles        []string  `json:"roles"`
}

type UserController struct {
	us *service.UserService
}

func UserRegister(group *gin.RouterGroup) {
	c := &UserController{
		us: new(service.UserService),
	}
	group.POST("/login", c.Login)
	group.GET("/logout", auth.SessionAuthMiddleware(), c.Logout)
	group.GET("/userinfo", auth.SessionAuthMiddleware(), c.UserInfo)
	group.GET("/changepassword", auth.SessionAuthMiddleware(), c.ChangePassword)
}

func (u *UserController) Login(c *gin.Context) {
	var login Login

	if err := c.ShouldBind(&login); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	user, err := u.us.CheckUser(c, login.UserName, login.Password)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	sessionInfo := &SessionInfo{
		ID:        user.ID,
		UserName:  user.UserName,
		LoginTime: time.Now(),
	}

	json, err := json.Marshal(sessionInfo)

	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}

	session := sessions.Default(c)
	session.Set(util.AdminSessionInfoKey, string(json))
	session.Save()

	result := &LoginResult{Token: user.UserName}

	middleware.ResponseSuccess(c, result)

}

func (u *UserController) ChangePassword(c *gin.Context) {
	var password ChangePassword

	if err := c.ShouldBind(&password); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	session := sessions.Default(c)
	sessionInfo := session.Get(util.AdminSessionInfoKey)
	userSessionInfo := &SessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessionInfo)), userSessionInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	user, err := u.us.FindUser(c, userSessionInfo.UserName)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	saltPassword := util.GenSaltPassword(user.Salt, password.Password)
	user.Password = saltPassword

	if err := u.us.ChangePassword(c, user.ID, user.Password); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	middleware.ResponseSuccess(c, "")

}

func (u *UserController) UserInfo(c *gin.Context) {

	sess := sessions.Default(c)
	sessInfo := sess.Get(util.AdminSessionInfoKey)
	userSessionInfo := &SessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), userSessionInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	result := &UserInfoResult{
		ID:           userSessionInfo.ID,
		Name:         userSessionInfo.UserName,
		LoginTime:    userSessionInfo.LoginTime,
		Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Introduction: "I am a super administrator",
		Roles:        []string{"admin"},
	}

	middleware.ResponseSuccess(c, result)
}

func (u *UserController) Logout(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Delete(util.AdminSessionInfoKey)
	sess.Save()
	middleware.ResponseSuccess(c, "")
}
