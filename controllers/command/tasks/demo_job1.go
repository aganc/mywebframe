package tasks

import (
	"airport/zlog"
	"context"
	"time"
)

func DemoJob1() {
	for {
		ctx := context.Background()
		zlog.Infof(ctx, "开始")
		zlog.Infof(ctx, "哈哈哈哈")
		zlog.Infof(ctx, "结束")
		time.Sleep(3 * time.Second)
		break
	}
	return
}
