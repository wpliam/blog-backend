package rdb

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func (cli *RedisClient) ZAdd(ctx context.Context, key string, member interface{}, score float64) error {
	return cli.cli.ZAdd(ctx, key, &redis.Z{
		Score:  score,
		Member: member,
	}).Err()
}

// ZIncrBy z set有序集合
func (cli *RedisClient) ZIncrBy(ctx context.Context, key string, member string, incr float64) (float64, error) {
	return cli.cli.ZIncrBy(ctx, key, incr, member).Result()
}

func (cli *RedisClient) ZScore(ctx context.Context, key string, member string) (float64, error) {
	return cli.cli.ZScore(ctx, key, member).Result()
}

func (cli *RedisClient) ZMScore(ctx context.Context, key string, member ...string) ([]float64, error) {
	return cli.cli.ZMScore(ctx, key, member...).Result()
}
