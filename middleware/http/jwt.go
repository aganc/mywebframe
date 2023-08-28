package middleware

import (
	"strings"

	"airport/defines"
	"airport/utils"

	"airport/zlog"
	"github.com/gin-gonic/gin"
)

const (
	// TODO: ATTENTION: 仅测试用!!!
	// 当请求使用此token串时，直接通过校验，并认为用户ID为1
	// 此过程不会查询mysql或redis
	//_testOnlyToken       = "test.only.token"
	//_adminUserID   int64 = 1
	TokenType = "Bearer"
)

func JWT(ctx *gin.Context) {

	token := ctx.GetHeader("Authorization")

	// "Bearer xxxxxxxx" ==> ["Bearer", "xxxxxxxx"]
	splitted := strings.Split(token, " ")
	if len(splitted) != 2 || splitted[0] != TokenType || splitted[1] == "" {
		//message.HTTPFailWithStatus(ctx, message.ErrorJWTUnauthorized, http.StatusUnauthorized)
		ctx.Abort()
		return
	}

	token = splitted[1]
	j := utils.NewJWT()

	//{
	//	// TODO: ATTENTION: 仅测试用!!!
	//	if token == _testOnlyToken {
	//		ctx.Set(defines.CtxKeyUserID, _adminUserID)
	//		return
	//	}
	//}

	// 校验
	claims, err := j.ParseToken(ctx, token)
	if err != nil {
		//message.HTTPFailWithStatus(ctx, message.ErrorJWTUnauthorized, http.StatusUnauthorized)
		ctx.Abort()
		return
	}

	// 注入UserID
	ctx.Set(defines.CtxKeyUserID, claims.UserID)

	// 刷新
	if claims.NeedRefresh() {
		tk, err := j.GenToken(ctx, claims)
		if err != nil {
			// 令牌刷新失败，只记录日志
			zlog.Warnf(ctx, "[jwt] refresh jwt error: %s", err)
		} else {
			// new-token
			ctx.Header(defines.JWTTokenRefreshHeader, tk) // 将新token串写入header
		}
	}
}
