package proxy

import (
	"blog-backend/repo/es"
	"blog-backend/repo/ftp"
	"blog-backend/repo/mdb"
	"blog-backend/repo/rdb"
)

// Proxy 代理
type Proxy interface {
	GetGormProxy() *mdb.MysqlClient
	GetRedisProxy() *rdb.RedisClient
	GetFtpProxy() *ftp.FtpClient
	GetEsProxy() *es.ElasticClient
}

func New(opts ...Option) Proxy {
	opt := defaultOption()
	for _, o := range opts {
		o(opt)
	}
	return opt
}

func defaultOption() *Options {
	return &Options{
		MysqlClient: mdb.NewMysqlClient("blog.mysql"),
		RedisClient: rdb.NewRedisClient("blog.redis"),
		FtpClient:   ftp.NewFtpClient("blog.ftp"),
		EsClient:    es.NewElasticClient("blog.elastic"),
	}
}
