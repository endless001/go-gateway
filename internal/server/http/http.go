package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-gateway/internal/conf"
	"go-gateway/internal/dao"
	"log"
	"net/http"
	"time"
)

var (
	d          *dao.Dao
	HttpServer *http.Server
)

func Run() {
	gin.SetMode("release")
	d = dao.New(conf.Conf)
	r := InitializeRoter()
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

func InitializeRoter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
