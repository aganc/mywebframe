package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"

	"airport/conf"
	"airport/helpers"
	"airport/router"
)

func main() {
	// gin
	engine := gin.New()
	/*
		tips: 此处初始化的资源是http和server共用的资源。
	*/
	// 初始化资源
	helpers.Init(engine)
	defer helpers.Clear()

	httpServer(engine)
}

// web server: http web 服务，容器里监听特定端口(:8080)实现web请求
func httpServer(engine *gin.Engine) {
	// 初始化http服务路由
	router.HTTP(engine)
	// MQ 消费回调路由
	//router.MQ(engine)
	// app内定时任务
	router.Tasks(engine)

	// 启动web server
	if err := Start(engine, conf.BasicConf.HTTP); err != nil {
		panic(err.Error())
	}
}

func Start(engine *gin.Engine, conf conf.ServerConfig) error {
	// todo: 后续根据环境区分，正式环境不允许用户指定端口
	appServer := endless.NewServer(conf.Address, engine)

	// 超时时间 (如果设置的太小，可能导致接口响应时间超过该值，进而导致504错误)
	if conf.ReadTimeout > 0 {
		appServer.ReadTimeout = conf.ReadTimeout
	}

	if conf.WriteTimeout > 0 {
		appServer.WriteTimeout = conf.WriteTimeout
	}

	// 监听http端口
	if err := appServer.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
