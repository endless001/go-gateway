package main


import (
	"context"
	"github.com/e421083458/go_gateway/golang_common/lib"
	"github.com/gin-gonic/gin"
	"go-gateway/router"
	"log"
	"net/http"
	"time"
)

type  App struct {}

var (
	HttpServer *http.Server
)
func (a *App) Run()  {
	gin.SetMode(lib.GetStringConf("base.base.debug_mode"))
    r:= router.InitializeRoter()
	HttpServer = &http.Server{
    	Addr: lib.GetStringConf("base.http.addr"),
    	Handler: r,
    	ReadTimeout: time.Duration(lib.GetIntConf("base.http.read_timeout")) * time.Second,
    	WriteTimeout: time.Duration(lib.GetIntConf("base.http.write_timeout")) * time.Second,
    	MaxHeaderBytes: 1 <<uint(lib.GetIntConf("base.http.max_header_bytes")),
	}
	go func() {
		 log.Printf("[INFO] HttpServer run %v\n",lib.GetStringConf("base.http.addr"))
		 if err := HttpServer.ListenAndServe();err !=nil{
		 	log.Fatalf(" [ERROR] HttpServer run:%s\n",lib.GetStringConf("base.http.addr"),err)
		 }
	}()
}

func (a *App) Stop()  {
	ctx,cancel :=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	if err := HttpServer.Shutdown(ctx);err != nil{
		log.Fatalf("[ERROR] HttpServer Stop err:%v\n",err)
	}
	log.Printf("[INFO] HttpServer stop")
}
