package mdb

import (
	agorm "github.com/wpliap/common-wrap/gorm"
	"gorm.io/gorm"
)

type MysqlClient struct {
	*gorm.DB
}

// NewMysqlClient 创建一个mysql client
func NewMysqlClient(name string) *MysqlClient {
	return &MysqlClient{
		agorm.NewGormProxy(name),
	}
}

func filterStatus() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", 1)
	}
}
