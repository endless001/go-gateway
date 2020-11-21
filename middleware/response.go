package middleware
import (
	"encoding/json"
	"fmt"
	"github.com/e421083458/go_gateway/golang_common/lib"
	"github.com/gin-gonic/gin"
	"strings"
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
	TraceId   interface{}  `json:trace_id`
	Stack     interface{}  `json:stack`
}

func ResponseError(c *gin.Context,code ResponseCode,err error)  {
	trace,_:=c.Get("trace")
	traceContext,_:=trace.(*lib.TraceContext)
	traceId := ""
	if traceContext != nil{
		traceId = traceContext.TraceId
	}
	stack := ""
	if c.Query("is_debug") =="1" ||lib.GetConfEnv()=="dev" {
		stack = strings.Replace(fmt.Sprintf("%+v",err),err.Error()+"\n","",-1)
	}
	resp := &Response{ErrorCode: code,ErrorMsg: err.Error(),Data: "",TraceId: traceId,Stack: stack}
	c.JSON(200,resp)
	response,_:=json.Marshal(resp)
	c.Set("response",string((response)))
	c.AbortWithError(200,err)
}

func  ResponseSuccess(c *gin.Context,data interface{})  {
	trace,_ := c.Get("trace")
	traceContext,_ := trace.(*lib.TraceContext)
	traceId := ""
	if traceContext != nil{
		traceId = traceContext.TraceId
	}
	resp := &Response{ErrorCode: SuccessCode,ErrorMsg: "",Data: data,TraceId: traceId}
	c.JSON(200,resp)
	response,_ := json.Marshal(resp)
	c.Set("response",string(response))
}
