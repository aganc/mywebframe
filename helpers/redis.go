package helpers

import (
	"airport/base"
	"airport/conf"
	"airport/zlog"
)

// 推荐，直接使用
var RedisClient *base.Redis

// 初始化redis
func InitRedis() {
	//demoRedisConf := conf.RConf.Redis["demo"]
	//DemoRedisClient = base.InitRedisClient(demoRedisConf)

	c := conf.RConf.Redis["airport"]
	var err error
	RedisClient, err = base.InitRedisClient(c)
	if err != nil || RedisClient == nil {
		panic("init spdata redis failed!")
	} else {
		zlog.Debugf(nil, "init redis success %v", err)
	}
}

func CloseRedis() {
	_ = RedisClient.Close()
}
