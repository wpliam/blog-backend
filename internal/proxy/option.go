package proxy

import (
	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

type proxy struct {
	dbCli      *gorm.DB
	redisCli   redis.UniversalClient
	elasticCli *elastic.Client
}

type Option func(*proxy)

func WithMysqlProxy(db *gorm.DB) Option {
	return func(p *proxy) {
		p.dbCli = db
	}
}

func WithRedisProxy(redisCli redis.UniversalClient) Option {
	return func(p *proxy) {
		p.redisCli = redisCli
	}
}

func WithElasticProxy(elasticCli *elastic.Client) Option {
	return func(p *proxy) {
		p.elasticCli = elasticCli
	}
}
