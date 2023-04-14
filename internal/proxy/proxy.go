package proxy

import (
	"blog-backend/internal/service"
	"blog-backend/repo/es"
	"blog-backend/repo/mdb"
	"blog-backend/repo/rdb"
)

func NewProxyService(opts ...Option) service.ProxyService {
	opt := &proxy{
		dbCli:      mdb.NewMysqlClient(),
		redisCli:   rdb.NewRedisClient(),
		elasticCli: es.NewElasticClient(),
	}
	for _, o := range opts {
		o(opt)
	}
	return opt
}

func (p *proxy) GetGormProxy() *mdb.MysqlClient {
	return p.dbCli
}

func (p *proxy) GetElasticProxy() *es.ElasticClient {
	return p.elasticCli
}

func (p *proxy) GetRedisProxy() *rdb.RedisClient {
	return p.redisCli
}
