package mdb

import (
	"blog-backend/model"
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

func filterStatus(status int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", status)
	}
}

func addPage(page *model.Page) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == nil || page.Limit == 0 {
			return db
		}
		return db.Offset((page.Offset - 1) * page.Limit).Limit(page.Limit)
	}
}
