package zlog

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/smallnest/rpcx/share"
	"go.uber.org/zap/zapcore"
)

// util key
const (
	ContextKeyRequestID     = "requestId"
	ContextKeyLogID         = "logId"
	ContextKeyNoLog         = "_no_log"
	ContextKeyUri           = "_uri"
	ContextKeyUserID        = "userId"
	ContextKeyTenantID      = "tenantId"
	ContextKeyMachineID     = "machineId"
	ContextKeyExternalAddr  = "externalAddr"
	ContextKeyUserAuthGroup = "userAuthGroup"
)

const ContextValAllAuthGroup = "-1" // 所有分组权限都有
const ContextValNoAuthGroup = "-2"  // 所有分组权限都没有

// rpcx key
const (
	StartTime = "startTime"
)

// header key
const (
	TraceHeaderKey      = "Uber-Trace-Id"
	LogIDHeaderKey      = "X_BD_LOGID"
	LogIDHeaderKeyLower = "x_bd_logid"
)

// GetLogID 兼容虚拟机调用项目logid串联问题
func GetLogID(ctx context.Context) string {
	if ctx == nil {
		return genRequestId()
	}
	if ginCtx, ok := ctx.(*gin.Context); ok {
		return getLogIDFromGinContext(ginCtx)
	} else if shareCtx, ok := ctx.(*share.Context); ok {
		return getLogIDFromRpcxContext(shareCtx)
	} else {
		return genRequestId()
	}
}

func GetRequestID(ctx context.Context) string {
	if ctx == nil {
		return genRequestId()
	}

	if ginCtx, ok := ctx.(*gin.Context); ok {
		return getRequestIDFromGinContext(ginCtx)
	} else if shareCtx, ok := ctx.(*share.Context); ok {
		return getRequestIDFromRpcxContext(shareCtx)
	} else {
		return genRequestId()
	}
}

func GetKeyURI(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if ginCtx, ok := ctx.(*gin.Context); ok {
		uri := ginCtx.GetString(ContextKeyUri)
		return uri
	} else if shareCtx, ok := ctx.(*share.Context); ok {
		v := shareCtx.Value(ContextKeyUri)
		if v != nil {
			if uri, ok := v.(string); ok {
				return uri
			}
		}
	}

	return ""
}

func GetUserID(ctx context.Context) int64 {
	userID := int64(0)
	if ctx == nil {
		return userID
	}
	if ginCtx, ok := ctx.(*gin.Context); ok {
		userID = getUserIDFromGinContext(ginCtx)
	} else if shareCtx, ok := ctx.(*share.Context); ok {
		userID = getUserIDFromRpcxContext(shareCtx)
	}
	return userID
}

func GetTenantID(ctx context.Context) int64 {
	tenantID := int64(0)
	if ctx == nil {
		return tenantID
	}
	if ginCtx, ok := ctx.(*gin.Context); ok {
		tenantID = getTenantIDFromGinContext(ginCtx)
	} else if shareCtx, ok := ctx.(*share.Context); ok {
		tenantID = getTenantIDFromRpcxContext(shareCtx)
	}
	return tenantID
}

func GetMachineID(ctx context.Context) string {
	machineID := ""
	if ctx == nil {
		return machineID
	}
	if ginCtx, ok := ctx.(*gin.Context); ok {
		machineID = getMachineIDFromGinContext(ginCtx)
	} else if shareCtx, ok := ctx.(*share.Context); ok {
		machineID = getMachineIDFromRpcxContext(shareCtx)
	}
	return machineID
}

func GetExternalAddress(ctx context.Context) string {
	addr := ""
	if ctx == nil {
		return addr
	}
	if ginCtx, ok := ctx.(*gin.Context); ok {
		addr = getExternalAddressFromGinContext(ginCtx)
	} else if shareCtx, ok := ctx.(*share.Context); ok {
		addr = getExternalAddressFromRpcxContext(shareCtx)
	}
	return addr
}

func GetUserAuthGroup(ctx context.Context) string {
	userAuthGroup := ""
	if ctx == nil {
		return userAuthGroup
	}
	if ginCtx, ok := ctx.(*gin.Context); ok {
		userAuthGroup = getUserAuthGroupFromGinContext(ginCtx)
	} else if shareCtx, ok := ctx.(*share.Context); ok {
		userAuthGroup = getUserAuthGroupFromRpcxContext(shareCtx)
	}
	return userAuthGroup
}

func GetCustomKV(ctx context.Context, key string) string {
	value := ""
	if ctx == nil {
		return value
	}
	if ginCtx, ok := ctx.(*gin.Context); ok {
		value = getCustomValueFromGinContext(ginCtx, key)
	} else if shareCtx, ok := ctx.(*share.Context); ok {
		value = getCustomValueFromRpcxContext(shareCtx, key)
	}
	return value
}

func getCustomValueFromGinContext(ctx *gin.Context, key string) string {
	if ctx == nil {
		return ""
	}
	// 从ctx中获取
	if val, ok := ctx.Get(key); ok {
		return fmt.Sprintf("%v", val)
	}
	return ""
}

func getCustomValueFromRpcxContext(ctx *share.Context, key string) string {
	if ctx == nil {
		return ""
	}
	// 从ctx中获取
	if r := ctx.Value(key); r != nil {
		return fmt.Sprintf("%v", r)
	}
	return ""
}

func getUserIDFromGinContext(ctx *gin.Context) int64 {
	if ctx == nil {
		return 0
	}
	// 从ctx中获取
	var userID int64
	userID = ctx.GetInt64(ContextKeyUserID)
	return userID
}

func getUserIDFromRpcxContext(ctx *share.Context) int64 {
	if ctx == nil {
		return 0
	}
	// 从ctx中获取
	var userID int64
	var ok bool
	if r := ctx.Value(ContextKeyUserID); r != nil {
		if userID, ok = r.(int64); !ok {
			userID = 0
		}
	}
	return userID
}

func getTenantIDFromGinContext(ctx *gin.Context) int64 {
	if ctx == nil {
		return 0
	}
	// 从ctx中获取
	var tenantID int64
	tenantID = ctx.GetInt64(ContextKeyTenantID)
	return tenantID
}

func getTenantIDFromRpcxContext(ctx *share.Context) int64 {
	if ctx == nil {
		return 0
	}
	// 从ctx中获取
	var tenantID int64
	var ok bool
	if r := ctx.Value(ContextKeyTenantID); r != nil {
		if tenantID, ok = r.(int64); !ok {
			tenantID = 0
		}
	}
	return tenantID
}

func getMachineIDFromGinContext(ctx *gin.Context) string {
	if ctx == nil {
		return ""
	}
	// 从ctx中获取
	var machineID string
	machineID = ctx.GetString(ContextKeyMachineID)
	return machineID
}

func getExternalAddressFromGinContext(ctx *gin.Context) string {
	if ctx == nil {
		return ""
	}
	// 从ctx中获取
	var addr string
	addr = ctx.GetString(ContextKeyExternalAddr)
	return addr
}

func getUserAuthGroupFromGinContext(ctx *gin.Context) string {
	if ctx == nil {
		return ""
	}
	// 从ctx中获取
	var userAuthGroup string
	userAuthGroup = ctx.GetString(ContextKeyUserAuthGroup)
	return userAuthGroup
}

func getMachineIDFromRpcxContext(ctx *share.Context) string {
	if ctx == nil {
		return ""
	}
	// 从ctx中获取
	var machineID string
	var ok bool
	if r := ctx.Value(ContextKeyMachineID); r != nil {
		if machineID, ok = r.(string); !ok {
			machineID = ""
		}
	}
	return machineID
}

func getExternalAddressFromRpcxContext(ctx *share.Context) string {
	if ctx == nil {
		return ""
	}
	// 从ctx中获取
	var addr string
	var ok bool
	if r := ctx.Value(ContextKeyExternalAddr); r != nil {
		if addr, ok = r.(string); !ok {
			addr = ""
		}
	}
	return addr
}

func getUserAuthGroupFromRpcxContext(ctx *share.Context) string {
	if ctx == nil {
		return ""
	}
	// 从ctx中获取
	var userAuthGroup string
	var ok bool
	if r := ctx.Value(ContextKeyUserAuthGroup); r != nil {
		if userAuthGroup, ok = r.(string); !ok {
			userAuthGroup = ""
		}
	}
	return userAuthGroup
}

func getLogIDFromGinContext(ctx *gin.Context) string {
	if ctx == nil {
		return genRequestId()
	}

	// 上次获取到的
	if logID := ctx.GetString(ContextKeyLogID); logID != "" {
		return logID
	}

	// 尝试从header中获取
	var logID string
	if ctx.Request != nil && ctx.Request.Header != nil {
		logID = ctx.GetHeader(LogIDHeaderKey)
		if logID == "" {
			logID = ctx.GetHeader(LogIDHeaderKeyLower)
		}
	}

	// 无法从上游获得，不展示logid，弱化logid
	if logID == "" {
		logID = genRequestId()
	}

	ctx.Set(ContextKeyLogID, logID)
	return logID
}

func getLogIDFromRpcxContext(ctx *share.Context) string {
	if ctx == nil {
		return genRequestId()
	}
	// 从ctx中获取
	var logID string
	var ok bool
	if r := ctx.Value(ContextKeyLogID); r != nil {
		if logID, ok = r.(string); !ok {
			logID = ""
		}
	}

	// 新生成
	if logID == "" {
		logID = genRequestId()
	}

	ctx.SetValue(ContextKeyLogID, logID)
	return logID
}

func getRequestIDFromGinContext(ctx *gin.Context) string {
	if ctx == nil {
		return genRequestId()
	}

	// 从ctx中获取
	if r := ctx.GetString(ContextKeyRequestID); r != "" {
		return r
	}

	// 优先从header中获取
	var requestId string
	if ctx.Request != nil && ctx.Request.Header != nil {
		requestId = ctx.Request.Header.Get(TraceHeaderKey)
	}

	// 新生成
	if requestId == "" {
		requestId = genRequestId()
	}

	ctx.Set(ContextKeyRequestID, requestId)
	return requestId
}

func getRequestIDFromRpcxContext(ctx *share.Context) string {
	if ctx == nil {
		return genRequestId()
	}
	// 从ctx中获取
	var requestId string
	var ok bool
	if r := ctx.Value(ContextKeyRequestID); r != nil {
		if requestId, ok = r.(string); !ok {
			requestId = ""
		}
	}

	// 新生成
	if requestId == "" {
		requestId = genRequestId()
	}

	ctx.SetValue(ContextKeyRequestID, requestId)
	return requestId
}

func genRequestId() (requestId string) {
	// 随机生成 todo: 随机生成的格式是否要统一成trace的格式
	usec := uint64(time.Now().UnixNano())
	requestId = strconv.FormatUint(usec&0x7FFFFFFF|0x80000000, 10)
	return requestId
}

func SetNoLogFlag(ctx context.Context) {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		ginCtx.Set(ContextKeyNoLog, true)
	} else if rpcxCtx, ok := ctx.(*share.Context); ok {
		rpcxCtx.SetValue(ContextKeyNoLog, true)
	}
}

func NoLog(ctx context.Context, lvl zapcore.Level) bool {
	// 防止低级别日志不输出但仍然频繁的创建zapLogger(ctx) OR sugaredLogger(ctx)对象，减轻GC压力
	//if !GetZapLogger().Core().Enabled(lvl) {
	//	return true
	//}

	if ctx == nil {
		return false
	}

	if ginCtx, ok := ctx.(*gin.Context); ok {
		flag, ok := ginCtx.Get(ContextKeyNoLog)
		if ok && flag == true {
			return true
		}
	} else if rpcxCtx, ok := ctx.(*share.Context); ok {
		flag := rpcxCtx.Value(ContextKeyNoLog)
		if flag != nil {
			if f, ok := flag.(bool); ok && f {
				return true
			}
		}
	}

	return false
}
