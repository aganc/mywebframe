package message

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HTTPRes struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// HTTPSuccess 业务成功
func HTTPSuccess(ctx *gin.Context, isComplete ...bool) {
	if len(isComplete) > 0 && isComplete[0] {
		ctx.JSON(200, &HTTPRes{Code: 0, Msg: "success", Data: struct{}{}})
	} else {
		ctx.JSON(200, struct{}{})
	}
}

// HTTPSuccessWithData 业务成功返回
// isComplete 表示是否以Code、Msg、Data完全结构返回数据
func HTTPSuccessWithData(ctx *gin.Context, data interface{}, isComplete ...bool) {
	var res interface{} = struct{}{}
	if data != nil {
		res = data
	}
	if len(isComplete) > 0 && isComplete[0] {
		ctx.JSON(200, &HTTPRes{Code: 0, Msg: "success", Data: res})
	} else {
		ctx.JSON(200, res)
	}
}

func HTTPFail(ctx *gin.Context, err error) {
	HTTPFailWithStatus(ctx, err, 200)
}

// HTTPFailWithData 业务失败带Data
func HTTPFailWithData(ctx *gin.Context, err error, data ...interface{}) {
	HTTPFailWithStatus(ctx, err, http.StatusOK, data...)
}

func HTTPFailWithStatus(ctx *gin.Context, err error, httpStatus int, data ...interface{}) {

	ctx.JSON(httpStatus, &HTTPRes{Code: 200, Msg: "system error"})
}

func HTTP400(ctx *gin.Context, err error) {
	HTTPFailWithStatus(ctx, err, http.StatusBadRequest)
}

func HTTP404(ctx *gin.Context, err error) {
	HTTPFailWithStatus(ctx, err, http.StatusNotFound)
}

