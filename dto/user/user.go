package user

import model "airport/models/db/user"

// AddUserInfo 添加用户参数
type AddUserInfo struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Tel    string `json:"telephone"`
	Passwd string `json:"password"`
	Desc   string `json:"description"`
}

// UpdateUserInfo 更新用户参数
type UpdateUserInfo struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Passwd string `json:"password"`
}

// CurUserReq 用户列表请求信息
type CurUserReq struct {
	Page Page `json:"page"`
}

// CurUserResp 用户列表响应信息
type CurUserResp struct {
	Items []*model.User `json:"items"`
	PageOutput
}

type DeleteUserReq struct {
	ID int64 `json:"id"`
}
