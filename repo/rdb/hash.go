package rdb

import "context"

// HIncrBy 自增
func (cli *RedisClient) HIncrBy(ctx context.Context, key string, field string, incr int64) (int64, error) {
	return cli.UniversalClient.HIncrBy(ctx, key, field, incr).Result()
}

func (cli *RedisClient) HSetNX(ctx context.Context, key string, field string, val interface{}) (bool, error) {
	return cli.UniversalClient.HSetNX(ctx, key, field, val).Result()
}

// HGet 获取key int64值
func (cli *RedisClient) HGet(ctx context.Context, key string, field string) (int64, error) {
	return cli.UniversalClient.HGet(ctx, key, field).Int64()
}

// HExists hash field 是否存在
func (cli *RedisClient) HExists(ctx context.Context, key string, field string) bool {
	result, err := cli.UniversalClient.HExists(ctx, key, field).Result()
	if err != nil {
		return false
	}
	return result
}
