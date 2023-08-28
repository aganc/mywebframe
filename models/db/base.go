package db

import "gorm.io/gorm"

type Model struct {
	ID       int64 `gorm:"primaryKey"`
	TenantID int64 `gorm:"column:tenant_id"`
	CreateAt int64 `gorm:"column:create_at;autoCreateTime"`
	UpdateAt int64 `gorm:"column:update_at;autoUpdateTime"`
}

func NotFoundToNil(err error) error {
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return err
}
