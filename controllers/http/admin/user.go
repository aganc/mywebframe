package admin

import (
	"airport/dto/user"
	"airport/message"
	srv "airport/service/user"
	"github.com/gin-gonic/gin"
)

func GetUserList(ctx *gin.Context) {
	args := &user.CurUserReq{}
	err := ctx.ShouldBindJSON(args)
	if err != nil {
		message.HTTPFail(ctx, err)
		return
	}
	list, err := srv.GetUserList(ctx, args)
	if err != nil {
		message.HTTPFail(ctx, err)
		return
	}
	message.HTTPSuccessWithData(ctx, list)
}

func CreateUser(ctx *gin.Context) {
	args := &user.AddUserInfo{}

	err := ctx.ShouldBindJSON(args)
	if err != nil {
		message.HTTPFail(ctx, err)
		return
	}

	err = srv.AddUser(ctx, args)
	if err != nil {
		message.HTTPFail(ctx, err)
		return
	}
	message.HTTPSuccessWithData(ctx, "success")
}

func UpdateUserInfo(ctx *gin.Context) {
	args := &user.UpdateUserInfo{}

	err := ctx.ShouldBindJSON(args)
	if err != nil {
		message.HTTPFail(ctx, err)
		return
	}

	err = srv.UpdateUserPwd(ctx, args)
	if err != nil {
		message.HTTPFail(ctx, err)
		return
	}
	message.HTTPSuccessWithData(ctx, "success")
}

func DeleteUserInfo(ctx *gin.Context) {
	args := &user.DeleteUserReq{}

	err := ctx.ShouldBindJSON(args)
	if err != nil {
		message.HTTPFail(ctx, err)
		return
	}

	err = srv.DeleteUser(ctx, args)
	if err != nil {
		message.HTTPFail(ctx, err)
		return
	}
	message.HTTPSuccessWithData(ctx, "success")
}
