package rdb

import "context"

// SAdd 集合中添加成员
func (cli *RedisClient) SAdd(ctx context.Context, key string, member interface{}) error {
	return cli.UniversalClient.SAdd(ctx, key, member).Err()
}

// SRem 集合中删除成员
func (cli *RedisClient) SRem(ctx context.Context, key string, member interface{}) error {
	return cli.UniversalClient.SRem(ctx, key, member).Err()
}

// SIsMember 成员是否在列表中
func (cli *RedisClient) SIsMember(ctx context.Context, key string, member interface{}) bool {
	result, err := cli.UniversalClient.SIsMember(ctx, key, member).Result()
	if err != nil {
		return false
	}
	return result
}

// SCard 获取set长度
func (cli *RedisClient) SCard(ctx context.Context, key string) (int64, error) {
	return cli.UniversalClient.SCard(ctx, key).Result()
}
