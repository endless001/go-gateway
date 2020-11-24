package http

import (
	"context"
	"github.com/e421083458/go_gateway/golang_common/lib"
	"github.com/gin-gonic/gin"
	"go-gateway/conf"
	"go-gateway/dao"
	"log"
	"net/http"
	"time"
)

var (
	d          *dao.Dao
	HttpServer *http.Server
)

func Init(c *conf.Config) {
	gin.SetMode(lib.GetStringConf("base.base.debug_mode"))
	d = dao.New(c)
	r := InitializeRoter()
	HttpServer = &http.Server{
		Addr:           lib.GetStringConf("base.http.addr"),
		Handler:        r,
		ReadTimeout:    time.Duration(lib.GetIntConf("base.http.read_timeout")) * time.Second,
		WriteTimeout:   time.Duration(lib.GetIntConf("base.http.write_timeout")) * time.Second,
		MaxHeaderBytes: 1 << uint(lib.GetIntConf("base.http.max_header_bytes")),
	}
}

func Run() {
	go func() {
		log.Printf("[INFO] HttpServer run %v\n", lib.GetStringConf("base.http.addr"))
		if err := HttpServer.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServer run:%s\n", lib.GetStringConf("base.http.addr"), err)
		}
	}()
}

func Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpServer.Shutdown(ctx); err != nil {
		log.Fatalf("[ERROR] HttpServer Stop err:%v\n", err)
	}
	log.Printf("[INFO] HttpServer stop")
}

func InitializeRoter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	//router.Use(middlewares...)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
