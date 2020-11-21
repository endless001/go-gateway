package controller

import (
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"go-gateway/middleware"
	"strings"
)

type OAuthController struct{}

func (o *OAuthController) Tokens(c *gin.Context) {

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

}
