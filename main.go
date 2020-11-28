package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"go-gateway/conf"
	"go-gateway/internal/server/http"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	conf.Load("conf/config.yaml")
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
	http.Run()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
