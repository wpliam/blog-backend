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

func (cli *RedisClient) ZExists(ctx context.Context, key string, member string) bool {
	score, err := cli.ZScore(ctx, key, member)
	if err == redis.Nil {
		return false
	}
	if err != nil {
		return false
	}
	return score > 0
}

func (cli *RedisClient) ZScan(ctx context.Context, key string, cursor uint64, match string) ([]string, uint64, error) {
	return cli.cli.ZScan(ctx, key, cursor, match, 10000).Result()
}

func (cli *RedisClient) ZRem(ctx context.Context, key string, member string) error {
	return cli.cli.ZRem(ctx, key, member).Err()
}
