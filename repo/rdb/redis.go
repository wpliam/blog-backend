package rdb

import (
	"blog-backend/repo/config"
	"context"
	"github.com/go-redis/redis/v8"
)

// RedisClient redis客户端
type RedisClient struct {
	cli redis.UniversalClient
}

func NewRedisClient() *RedisClient {
	conf := config.GetRedisConf()
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{conf.Addr},
		Password: conf.Password,
		DB:       0,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic("redis client ping err " + err.Error())
	}
	return &RedisClient{
		cli: client,
	}
}
