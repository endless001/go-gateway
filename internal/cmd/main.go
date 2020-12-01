package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"go-gateway/internal/conf"
	"go-gateway/internal/database"
	"go-gateway/internal/router"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	conf.Load("internal/conf/config.yaml")
	database.New(conf.Conf)
	// 设置日志格式为json格式
	log.SetFormatter(&log.JSONFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	log.SetLevel(log.WarnLevel)
}

func main() {

	flag.Parse()
	router.Run()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
