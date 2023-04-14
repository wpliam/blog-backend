package rdb

import (
	"context"
	"time"
)

// Exists key是否存在 存在返回1
func (cli *RedisClient) Exists(ctx context.Context, key string) bool {
	result, err := cli.cli.Exists(ctx, key).Result()
	if err != nil {
		return false
	}
	return result == 1
}

// Get 获取key的值
func (cli *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return cli.cli.Get(ctx, key).Result()
}

func (cli *RedisClient) GetInt64(ctx context.Context, key string) (int64, error) {
	return cli.cli.Get(ctx, key).Int64()
}

// Set ...
func (cli *RedisClient) Set(ctx context.Context, key string, value interface{}, expr time.Duration) error {
	return cli.cli.Set(ctx, key, value, expr).Err()
}

// SetNX set if not exists
func (cli *RedisClient) SetNX(ctx context.Context, key string, value interface{}, expr time.Duration) (bool, error) {
	return cli.cli.SetNX(ctx, key, value, expr).Result()
}

// SetEX ...
func (cli *RedisClient) SetEX(ctx context.Context, key string, value interface{}, expr time.Duration) error {
	return cli.cli.SetEX(ctx, key, value, expr).Err()
}

// Del 删除key
func (cli *RedisClient) Del(ctx context.Context, key string) error {
	return cli.cli.Del(ctx, key).Err()
}
