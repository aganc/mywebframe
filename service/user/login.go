package user

import (
	data "airport/data/user"
	"airport/defines"
	"airport/dto/user"
	model "airport/models/db/user"
	"airport/utils"
	"airport/zlog"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// Login 请求登录，通过邮箱和密码登录
func Login(ctx context.Context, req *user.HTTPLoginReq) (reply *user.HTTPLoginRes, err error) {

	// 用户校验，注意免密登录用户不可页面登录
	var userMdl *model.User
	var jwtCheckType int64

	userMdl, err = generalAuth(ctx, req.Username, req.Password)
	if userMdl != nil || err != nil {
		return nil, err
	}

	// 生成JWT
	j := utils.NewJWT()
	claims := j.NewClaims(userMdl.ID, jwtCheckType)
	token, err := j.GenToken(ctx, claims)
	if err != nil {
		zlog.Warnf(ctx, "generating jwt token failed: %s", err)
		return
	}

	// 更新用户最后登录token
	//err = updateUserToken(ctx, claims.UserID, claims.TokenID)
	//if err != nil {
	//	zlog.Warnf(ctx, "updating latest user token failed: %s", err)
	//	return nil, message.ErrorRedisSet
	//}

	reply = &user.HTTPLoginRes{
		Token: user.JWTToken{
			TokenType: defines.TokenType,
			Token:     token,
		},
		User: user.User{
			ID:    userMdl.ID,
			Name:  userMdl.Name,
			Email: userMdl.Email,
		},
	}
	return
}

// generalAuth 普通用户登录校验
func generalAuth(ctx context.Context, username, password string) (user *model.User, err error) {
	user, err = data.QueryOne(ctx, &model.User{Name: username})
	if user == nil {
		if err != nil {
			zlog.Warnf(ctx, "[user auth] query user failed: %s", err)
		}
		err = errors.New("user auth error")
		return
	}

	if !checkPassword(user.Passwd, password) {
		err = errors.New("user pwd auth error")
		return
	}

	return
}

func checkPassword(hash, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass)) == nil
}
