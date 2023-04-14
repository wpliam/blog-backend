package rdb

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func (cli *RedisClient) SetBit(ctx context.Context, key string, offset int64, value int) error {
	return cli.cli.SetBit(ctx, key, offset, value).Err()
}

func (cli *RedisClient) GetBit(ctx context.Context, key string, offset int64) bool {
	result, err := cli.cli.GetBit(ctx, key, offset).Result()
	if err != nil {
		return false
	}
	return result == 1
}

func (cli *RedisClient) BitCount(ctx context.Context, key string) (int64, error) {
	return cli.cli.BitCount(ctx, key, &redis.BitCount{
		Start: 0,
		End:   -1,
	}).Result()
}

func (cli *RedisClient) BitGetDay(ctx context.Context, key string) ([]int64, error) {
	dayOfMonth := fmt.Sprintf("u%d", time.Now().Day())
	return cli.cli.BitField(ctx, key, "GET", dayOfMonth, 0).Result()
}
