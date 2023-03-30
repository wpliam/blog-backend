package proxy

import (
	"github.com/go-redis/redis/v8"
	"github.com/jlaffaye/ftp"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

type proxy struct {
	dbCli      *gorm.DB
	redisCli   redis.UniversalClient
	ftpCli     *ftp.ServerConn
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

func WithFtpProxy(ftpCli *ftp.ServerConn) Option {
	return func(p *proxy) {
		p.ftpCli = ftpCli
	}
}

func WithElasticProxy(elasticCli *elastic.Client) Option {
	return func(p *proxy) {
		p.elasticCli = elasticCli
	}
}
