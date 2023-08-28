package router

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func Tasks(engine *gin.Engine) {

	// 定时任务
	startCrontab(engine)
}

func startCrontab(engine *gin.Engine) {
	cronJob := InitCrontab(engine)
	//cronJob.AddFunc("@every 1d", tasks.DemoJob1)
	cronJob.Start()
}

func InitCrontab(engine *gin.Engine) (c *cron.Cron) {
	c = cron.New()
	return c
}
