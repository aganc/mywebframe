package helpers

import (
	"airport/conf"
	"airport/env"
	"airport/zlog"
	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {
	PreInit()
	InitResource(engine)
}

func Clear() {
	// 服务结束时的清理工作，对应 Init() 初始化的资源
	//zlog.CloseLogger()
	//CloseKafkaProducer()
	//CloseRocketMq()
}

func PreInit() {
	// 用于日志中展示模块的名字，开发环境需要手动指定，容器中无需手动指定
	env.SetAppName("airport")

	// 配置加载
	conf.InitConf()

	// 日志初始化
	zlog.InitLog(conf.BasicConf.Log)
}

func InitResource(engine *gin.Engine) {
	// 初始化全局变量
	InitJob(engine)
	InitPgsql()
	InitRedis()
	InitHTTPClient()
	//InitMysql()
	//InitClickHouse()
	//InitRocketMq()
	//InitRpcxClient()
	//InitExternal()
	//InitEs()
	//InitHBase()
	//InitGCache()
	//InitKafkaProducer()
}
