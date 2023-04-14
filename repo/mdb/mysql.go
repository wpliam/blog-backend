package mdb

import (
	"blog-backend/repo/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MysqlClient mysql客户端
type MysqlClient struct {
	cli *gorm.DB
}

func NewMysqlClient() *MysqlClient {
	conf := config.GetMysqlConf()
	db, err := gorm.Open(mysql.Open(conf.Target), &gorm.Config{})
	if err != nil {
		panic("gorm open err " + err.Error())
	}
	return &MysqlClient{
		cli: db,
	}
}

func filterStatus() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", 1)
	}
}
