package mdb

import (
	agorm "github.com/wpliap/common-wrap/gorm"
	"gorm.io/gorm"
	"sync"
)

var (
	gormProxy = make(map[string]*client)
	rw        sync.RWMutex
)

const defaultGormName = "gorm"

type client struct {
	*gorm.DB
}

func GetGormClient() *client {
	rw.RLock()
	defer rw.RUnlock()
	cli, ok := gormProxy[defaultGormName]
	if ok && cli != nil {
		return cli
	}
	gormCli := agorm.NewGormProxy("blog.mysql")
	cli = &client{
		gormCli,
	}
	gormProxy[defaultGormName] = cli
	return cli
}
