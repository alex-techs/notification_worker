package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"source.hitokoto.cn/hitokoto/notification-worker/src/aliyun/directmail"

	// 项目内文件
	"source.hitokoto.cn/hitokoto/notification-worker/src/config"
	"source.hitokoto.cn/hitokoto/notification-worker/src/event"

	// 外部依赖
	log "github.com/sirupsen/logrus"
	// "github.com/streadway/amqp"
)

// 程序信息
var (
	DEBUG = true
	v     bool
	c     string
)

func initLogger() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	if DEBUG { // 内编
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func init() {
	flag.BoolVar(&v, "v", false, "查看版本信息")
	flag.StringVar(&c, "c", "", "设定配置文件")
	flag.Parse()
	if v {
		fmt.Printf("NotificationWorker ©2020 MoeTeam All Rights Reserved. \n当前版本: %s \nGitCommit: %s \n", Version, GitCommit)
		os.Exit(0)
	}
	initLogger()
	config.Init(c)
	// 设置生产日记级别
	if config.Debug() {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	// 初始化阿里云 SDK
	directmail.InitAliyunDirectMail()
}

func main() {
	log.Infoln("服务已初始化，开始核心服务。程序版本：" + Version + "，构建于 " + runtime.Version() + "。 Git 标签：" + GitCommit)
	go event.InitRabbitMQEvent()
	select {} // 堵塞方法
}
