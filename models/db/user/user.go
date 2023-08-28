package user

type User struct {
	ID     int64  `gorm:"primaryKey"`
	Name   string `gorm:"column:name"`
	Email  string `gorm:"column:email"`
	Passwd string `gorm:"column:password"`
	Tel    string `gorm:"column:tel"`
	Desc   string `gorm:"column:desc"`
}

func (*User) TableName() string {
	return "user"
}
