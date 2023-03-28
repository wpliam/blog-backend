package proxy

import (
	"blog-backend/repo/es"
	"blog-backend/repo/ftp"
	"blog-backend/repo/mdb"
	"blog-backend/repo/rdb"
	"github.com/wpliap/common-wrap/elasticsearch"
	aftp "github.com/wpliap/common-wrap/ftp"
	agorm "github.com/wpliap/common-wrap/gorm"
	aredis "github.com/wpliap/common-wrap/redis"
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
		MysqlClient: mdb.NewMysqlClient(agorm.NewGormProxy("blog.mysql")),
		RedisClient: rdb.NewRedisClient(aredis.NewRedisProxy("blog.redis")),
		FtpClient:   ftp.NewFtpClient(aftp.NewFtpProxy("blog.ftp")),
		EsClient:    es.NewElasticClient(elasticsearch.NewElasticProxy("blog.elastic")),
	}
}
