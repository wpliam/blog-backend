package rdb

import (
	"github.com/go-redis/redis/v8"
	aredis "github.com/wpliap/common-wrap/redis"
)

type RedisClient struct {
	redis.UniversalClient
}

// NewRedisClient ...
func NewRedisClient(name string) *RedisClient {
	return &RedisClient{
		aredis.NewRedisProxy(name),
	}
}
