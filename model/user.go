package model

// User 用户表
type User struct {
	Model
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Desc     string `json:"desc"`
}

func (*User) TableName() string {
	return UserTableName
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
