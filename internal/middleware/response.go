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
	ErrorCode ResponseCode `josn:errcode`
	ErrorMsg  string       `json:errmsg`
	Data      interface{}  `json:data`
}

func ResponseError(c *gin.Context, code ResponseCode, err error) {

	resp := &Response{ErrorCode: code, ErrorMsg: err.Error(), Data: ""}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string((response)))
	c.AbortWithError(200, err)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	resp := &Response{ErrorCode: SuccessCode, ErrorMsg: "", Data: data}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
