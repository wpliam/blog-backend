package rdb

import (
	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	redis.UniversalClient
}

// NewRedisClient ...
func NewRedisClient(cli redis.UniversalClient) *RedisClient {
	return &RedisClient{
		cli,
	}
}
