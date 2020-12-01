package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-gateway/internal/conf"
	"log"
	"net/http"
	"time"
)

var (
	HttpServer *http.Server
)

func Run() {
	gin.SetMode("debug")
	r := RegisterRoutes()
	HttpServer = &http.Server{
		Addr:           conf.Conf.Server.Address,
		Handler:        r,
		ReadTimeout:    time.Duration(conf.Conf.Server.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(conf.Conf.Server.ReadTimeout) * time.Second,
		MaxHeaderBytes: 1 << uint(conf.Conf.Server.MaxHeaderBytes),
	}
	go func() {
		log.Printf("[INFO] HttpServer run %v\n", conf.Conf.Server.Address)
		if err := HttpServer.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServer run:%s\n", conf.Conf.Server.Address, err)
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
