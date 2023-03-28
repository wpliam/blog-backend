package mdb

import (
	"gorm.io/gorm"
)

type MysqlClient struct {
	*gorm.DB
}

// NewMysqlClient 创建一个mysql client
func NewMysqlClient(db *gorm.DB) *MysqlClient {
	return &MysqlClient{
		db,
	}
}

func filterStatus() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", 1)
	}
}
