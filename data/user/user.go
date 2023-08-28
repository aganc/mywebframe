package user

import (
	"airport/helpers"
	model "airport/models/db/user"
	"context"
	"errors"
)

// QueryOne 根据用户信息查询首条记录
func QueryOne(ctx context.Context, user *model.User) (*model.User, error) {
	db := helpers.PgsqlClientNtExam.WithContext(ctx).Model(&model.User{})

	reply := &model.User{}
	db = db.Where(user).First(reply)

	if db.Error != nil {
		return nil, errors.New("not found query user")
	}

	return reply, nil
}

// AddUser 添加用户
func AddUser(ctx context.Context, user *model.User) (err error) {
	db := helpers.PgsqlClientNtExam.WithContext(ctx).Model(&model.User{})

	db = db.Create(user)

	if db.Error != nil {
		return errors.New("add user error")
	}

	return nil
}

func DeleteUser(ctx context.Context, user *model.User) (err error) {
	db := helpers.PgsqlClientNtExam.WithContext(ctx).Model(&model.User{})

	db = db.Where(user).Delete(user)
	if db.Error != nil {
		return errors.New("delete user error")
	}

	return nil
}

func GetUserList(ctx context.Context, limit, offset int) (list []*model.User, count int64, err error) {
	db := helpers.PgsqlClientNtExam.WithContext(ctx).Model(&model.User{})

	err = db.Count(&count).Error
	if err != nil {
		return nil, 0, errors.New("get user list count error")
	}

	db = db.Limit(limit).Offset(offset).Find(&list)
	if db.Error != nil {
		return nil, 0, errors.New("get user list error")
	}

	return list, count, nil
}

func UpdateUserInfo(ctx context.Context, user *model.User) error {
	db := helpers.PgsqlClientNtExam.WithContext(ctx).Model(&model.User{})

	db = db.Where("id = ?", user.ID).Updates(user)
	if db.Error != nil {
		return errors.New("update user info error")
	}

	return nil
}
