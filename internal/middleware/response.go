package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type ResponseCode int

const (
	SuccessCode ResponseCode = iota
	UndefErrorCode
	ValidErrorCode
	InternalErrorCode

	InvaildRequestErrorCode ResponseCode = 401
	CustomizeCode           ResponseCode = 1000
	GROUPALL_SAVE_FLOWERROR ResponseCode = 2001
)

type Response struct {
	Code    ResponseCode `josn:code`
	Message string       `json:message`
	Data    interface{}  `json:data`
}

func ResponseError(c *gin.Context, code ResponseCode, err error) {

	resp := &Response{Code: code, Message: err.Error(), Data: ""}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string((response)))
	c.AbortWithError(200, err)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	resp := &Response{Code: SuccessCode, Message: "", Data: data}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
