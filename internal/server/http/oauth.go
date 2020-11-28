package http

import (
	"encoding/base64"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-gateway/internal/middleware"
	"go-gateway/internal/util"
	"strings"
	"time"
)

func Tokens(c *gin.Context) {

	splits := strings.Split(c.GetHeader("Authorization"), " ")
	if len(splits) != 2 {
		middleware.ResponseError(c, 2000, errors.New("用户名或密码格式错误"))
		return
	}
	appSecret, err := base64.StdEncoding.DecodeString(splits[1])
	if err != nil {
		middleware.ResponseError(c, 20002, err)
	}
	parts := strings.Split(string(appSecret), ":")
	if len(parts) != 2 {
		middleware.ResponseError(c, 2003, errors.New("用户名或密码格式错误"))
	}
	appInfo, err := d.GetTenant(c, parts[0], parts[1])
	if err != nil {
		middleware.ResponseError(c, 2004, err)
		return
	}
	if appInfo != nil {

		localtime, err := time.LoadLocation("")

		claims := jwt.StandardClaims{
			Issuer:    appInfo.AppID,
			ExpiresAt: time.Now().Add(util.JwtExpires * time.Second).In(localtime).Unix(),
		}
		if err != nil {
			middleware.ResponseError(c, 2004, err)
			return
		}

		middleware.ResponseSuccess(c, claims)
		return
	}
	middleware.ResponseError(c, 2005, errors.New("未匹配正确APP信息"))
}

func Signout(c *gin.Context) {

	sess := sessions.Default(c)
	sess.Delete(util.AdminSessionInfoKey)
	sess.Save()
	middleware.ResponseSuccess(c, "ok")

}
