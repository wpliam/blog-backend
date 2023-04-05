package rdb

import (
	"context"
)

func (cli *RedisClient) SetBit(ctx context.Context, key string, offset int64, value int) error {
	return cli.UniversalClient.SetBit(ctx, key, offset, value).Err()
}

func (cli *RedisClient) GetBit(ctx context.Context, key string, offset int64) bool {
	result, err := cli.UniversalClient.GetBit(ctx, key, offset).Result()
	if err != nil {
		return false
	}
	return result == 1
}
