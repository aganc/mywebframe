package admin

import (
	"airport/dto/user"
	"airport/message"
	srv "airport/service/user"
	"github.com/gin-gonic/gin"
)

// Login 用户登录
// @Summary   用户登录
// @Tags      用户
// @Accept    json
// @Product   json
// @Param     param  body      dto.HTTPLoginReq  true  "用户登录信息"
// @Success   200   {object}  dto.HTTPLoginRes
// @Router    /login [POST]
func Login(ctx *gin.Context) {
	args := &user.HTTPLoginReq{}

	err := ctx.ShouldBindJSON(args)
	if err != nil {
		message.HTTPFail(ctx, err)
		return
	}

	res, err := srv.Login(ctx, args)
	if err != nil {
		message.HTTPFail(ctx, err)
		return
	}

	message.HTTPSuccessWithData(ctx, res)
	return
}

// Logout 用户登出
// @Security  TokenAuth
// @Summary   用户登出
// @Tags      用户
// @Accept    json
// @Product   json
// @Success   200    {object}  dto.HTTPLoginRes
// @Router    /logout [POST]
func Logout(ctx *gin.Context) {
	return
}
