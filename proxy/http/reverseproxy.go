package http

import (
	"github.com/gin-gonic/gin"
	"go-gateway/middleware"
	"net/http"
	"net/http/httputil"
	"strings"
)

func NewLoadBalanceReverseProxy(c *gin.Context, trans *http.Transport) *httputil.ReverseProxy {
	director := func(req *http.Request) {

	}
	modifyFunc := func(resp *http.Response) error {
		if strings.Contains(resp.Header.Get("Connection"), "Upgrade") {
			return nil
		}
		return nil
	}
	errFunc := func(w http.ResponseWriter, r *http.Request, err error) {
		middleware.ResponseError(c, 999, err)
	}
	return &httputil.ReverseProxy{Director: director, ModifyResponse: modifyFunc, ErrorHandler: errFunc}
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}
