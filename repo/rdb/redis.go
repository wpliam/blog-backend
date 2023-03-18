package rdb

import (
	"github.com/go-redis/redis/v8"
	aredis "github.com/wpliap/common-wrap/redis"
	"sync"
)

var (
	redisProxy = make(map[string]*Client)
	rw         sync.RWMutex
)

const defaultRedisName = "redis"

type Client struct {
	redis.UniversalClient
}

func GetRedisClient() *Client {
	rw.RLock()
	defer rw.RUnlock()
	client, ok := redisProxy[defaultRedisName]
	if ok && client != nil {
		return client
	}
	redisCli := aredis.NewRedisProxy("blog.redis")
	client = &Client{
		redisCli,
	}
	redisProxy[defaultRedisName] = client
	return client
}
