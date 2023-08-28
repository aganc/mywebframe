/**************************************************************
 * Copyright (c) 2021 anxin.com, Inc. All Rights Reserved
 * User: zhangdongsheng<zhangdongsheng@anxin.com>
 * Date: 2021/9/5
 * Desc:
 **************************************************************/

package utils

import (
	"airport/env"
	"airport/zlog"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SetPanicLog(c *gin.Context, err interface{}) {
	// 请求url
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	if raw != "" {
		path = path + "?" + raw
	}
	// 请求报文
	body := c.GetString("body")

	fields := []zlog.Field{
		zlog.String("logId", zlog.GetLogID(c)),
		zlog.String("requestId", zlog.GetRequestID(c)),
		zlog.String("uri", path),
		zlog.String("refer", c.Request.Referer()),
		zlog.String("clientIp", c.ClientIP()),
		zlog.String("module", env.AppName),
		zlog.String("ua", c.Request.UserAgent()),
		zlog.String("host", c.Request.Host),
		zlog.String("body", body),
		zlog.String("errMsg", fmt.Sprintf("%+v", err)),
	}
	zlog.InfoLogger(c, "Panic[recover]", fields...)
}
