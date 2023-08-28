package middleware

import (
	"airport/utils"
	"github.com/gin-gonic/gin"
	"net"
	"os"
	"strings"
)

func Recover(ctx *gin.Context) {
	defer CatchRecoverRpc(ctx)
	ctx.Next()
}

// 针对rpc接口的处理
func CatchRecoverRpc(c *gin.Context) {
	// panic捕获
	if err := recover(); err != nil {
		utils.SetPanicLog(c, err)
		if checkBrokenPipe(c, err) {
			handleBrokenPipePanic(c)
		} else {
			handleCommanPanic(c)
		}
	}
}

func checkBrokenPipe(ctx *gin.Context, err interface{}) (brokenPipe bool) {
	if ne, ok := err.(*net.OpError); ok {
		if se, ok := ne.Err.(*os.SyscallError); ok {
			if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
				brokenPipe = true
			}
		}
	}
	return
}

func handleBrokenPipePanic(c *gin.Context) {
	c.Abort()
}

func handleCommanPanic(c *gin.Context) {
	panicReport(c)
}

func panicReport(ctx *gin.Context) {
	//track := "Module:" + env.AppName + "\n\nLogId:" + zlog.GetLogID(ctx) + "\n\nRequestId:" + zlog.GetLogID(ctx) + "\n\n"
	//httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
	//httpBody := ctx.GetString("body")
	//debugStackStart := "\n\n-----------------------以下是调用栈信息-----------------------\n\n"
	//panicContent := track + string(httpRequest) + httpBody + debugStackStart + string(debug.Stack())
	//zlog.Fatal(ctx, panicContent)
	//err := dingding.Send(dingding.DingRobotBusiness, panicContent)
	//if err != nil {
	//	zlog.Warnf(ctx, "ding ding send fail. error: %v", err)
	//}
	//
	//emailContent := strings.ReplaceAll(strings.ReplaceAll(panicContent, "\r\n", "\n"), "\n", "\r\n")
	//_, err = mail.SendMail(ctx, []string{"zxxxx@xxx.com"}, "ntpanic", emailContent)
	//if err != nil {
	//	zlog.Warnf(ctx, "send mail fail. error: %v", err)
	//}
}
