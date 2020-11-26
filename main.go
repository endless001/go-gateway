package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// 设置日志格式为json格式
	log.SetFormatter(&log.JSONFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	log.SetLevel(1)
}

func main() {

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Fatal("A group of walrus emerges from the ocean")

}
