package mdb

import (
	"gorm.io/gorm"
)

type MysqlClient struct {
	*gorm.DB
}

func filterStatus() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", 1)
	}
}
