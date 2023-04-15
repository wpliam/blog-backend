package model

import "gorm.io/gorm"

// User 用户表
type User struct {
	Model
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Role     uint32 `yaml:"role"` // 0:普通用户 1:管理员
	Desc     string `json:"desc"`
}

func (*User) TableName() string {
	return UserTableName
}

func (u *User) FindAfter(db *gorm.DB) error {
	return nil
}

// Account 账号信息
type Account struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (Account) TableName() string {
	return UserTableName
}
