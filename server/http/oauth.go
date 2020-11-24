package http

import (
	"encoding/base64"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/e421083458/go_gateway/golang_common/lib"
	"github.com/e421083458/go_gateway/public"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-gateway/middleware"
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
	appInfo, err := d.GetApp(c, parts[0], parts[1])
	if err != nil {
		middleware.ResponseError(c, 2004, err)
		return
	}
	if appInfo != nil {

		claims := jwt.StandardClaims{
			Issuer:    appInfo.AppID,
			ExpiresAt: time.Now().Add(public.JwtExpires * time.Second).In(lib.TimeLocation).Unix(),
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
	sess.Delete(public.AdminSessionInfoKey)
	sess.Save()
	middleware.ResponseSuccess(c, "ok")

}
