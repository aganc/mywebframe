package metadata

import (
	"airport/defines"
	"context"
	"github.com/gin-gonic/gin"
)

func CurrentLoginID(ctx context.Context) (adminID int64) {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		adminID = ginCtx.GetInt64(defines.CtxKeyUserID)
	} else {
		adminID, _ = ctx.Value(defines.CtxKeyUserID).(int64)
	}
	return adminID
}

// SetOutsidePlatformCode 设置外部接口访问的平台编码
func SetOutsidePlatformCode(ctx *gin.Context, platformCode string) {
	ctx.Set(defines.CtxKeyOutsidePlatformCode, platformCode)
}

// GetOutsidePlatformCode 获取外部接口访问的平台编码
func GetOutsidePlatformCode(ctx context.Context) (platformCode string) {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		platformCode = ginCtx.GetString(defines.CtxKeyOutsidePlatformCode)
	} else {
		platformCode, _ = ctx.Value(defines.CtxKeyOutsidePlatformCode).(string)
	}
	return platformCode
}
