package user

import (
	data "airport/data/user"
	"airport/dto/user"
	model "airport/models/db/user"
	"airport/zlog"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(ctx context.Context, req *user.AddUserInfo) error {
	if req.Passwd == "" || req.Name == "" || req.Email == "" {
		return errors.New("params is empty")
	}
	hashed := hashPassword(ctx, req.Passwd)

	userMdl := &model.User{
		Name:   req.Name,
		Email:  req.Email,
		Tel:    req.Tel,
		Desc:   req.Desc,
		Passwd: hashed,
	}

	err := data.AddUser(ctx, userMdl)

	return err

}

func hashPassword(ctx context.Context, pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 10)

	if err != nil {
		zlog.Warnf(ctx, "[HashPassword] error: %s", err)
		return ""
	}

	return string(hash)
}

func GetUserList(ctx context.Context, req *user.CurUserReq) (res *user.CurUserResp, err error) {
	list, count, err := data.GetUserList(ctx, req.Page.PageIndex, req.Page.PageSize)
	res.Items = list
	res.Total = count
	res.PageIndex = req.Page.PageIndex
	return res, err
}

func UpdateUserPwd(ctx context.Context, req *user.UpdateUserInfo) error {

	if req.ID == 0 || req.Passwd == "" {
		return errors.New("params is empty")
	}
	hashed := hashPassword(ctx, req.Passwd)

	userMdl := &model.User{
		ID:     req.ID,
		Passwd: hashed,
	}

	err := data.UpdateUserInfo(ctx, userMdl)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(ctx context.Context, req *user.DeleteUserReq) error {
	if req.ID == 0 {
		return errors.New("params is empty")
	}

	userMdl := &model.User{
		ID: req.ID,
	}

	err := data.DeleteUser(ctx, userMdl)
	if err != nil {
		return err
	}
	return nil
}
