/**************************************************************
 * Copyright (c) 2021 anxin.com, Inc. All Rights Reserved
 * User: zhangdongsheng<zhangdongsheng@anxin.com>
 * Date: 2021/9/5
 * Desc:
 **************************************************************/

package middleware

import (
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		//method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}
		//放行所有OPTIONS方法
		//if method == "OPTIONS" {
		//	c.AbortWithStatus(http.StatusNoContent)
		//}
		c.Next()
	}
}
