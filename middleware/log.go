package middleware

import "C"
import (
	"bytes"
	"github.com/e421083458/go_gateway/golang_common/lib"
	"github.com/e421083458/go_gateway/public"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)

func RequestInLog(c *gin.Context)  {
	traceContext :=lib.NewTrace()
	if traceId :=c.Request.Header.Get("com-header-rid");traceId != ""{
		traceContext.TraceId=traceId
	}
	if spanId :=c.Request.Header.Get("com-header-spanid");spanId != ""{
		traceContext.SpanId=spanId
	}
	c.Set("startExecTime",time.Now())
	c.Set("trace",traceContext)

	bodyBytes,_:=ioutil.ReadAll(c.Request.Body)
	c.Request.Body =ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	lib.Log.TagInfo(traceContext,"_com_request_in",map[string]interface{}{
		"uri" : c.Request.RequestURI,
		"method" : c.Request.Method,
		"args" : c.Request.PostForm,
		"body" : string(bodyBytes),
		"form" : c.ClientIP(),
	})
}
func RequestOutLog(c *gin.Context)  {
	endExecTime :=time.Now()
	response,_:=c.Get("response")
	st,_ := c.Get("startExecTime")
	startExecTime,_:=st.(time.Time)
	public.ComLogNotice(c,"_com_request_out", map[string]interface{}{
		"uri": c.Request.RequestURI,
		"method": c.Request.Method,
		"args" : c.Request.PostForm,
		"form" : c.ClientIP(),
		"response" :response,
		"proc_time" :endExecTime.Sub(startExecTime).Seconds(),
	})
}
func RequestLog() gin.HandlerFunc  {
	 return func(c *gin.Context) {
		 if lib.GetBoolConf("base.log.file_writer.on"){
		 	RequestInLog(c)
		 	defer RequestOutLog(c)
		 }
		 c.Next()
	 }
	 
}