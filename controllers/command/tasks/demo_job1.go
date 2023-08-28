/**************************************************************
 * Copyright (c) 2021 anxin.com, Inc. All Rights Reserved
 * User: zhangdongsheng<zhangdongsheng@anxin.com>
 * Date: 2021/9/5
 * Desc:
 **************************************************************/

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
