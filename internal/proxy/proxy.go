package proxy

import (
	"blog-backend/internal/service"
	"blog-backend/repo/es"
	"blog-backend/repo/ftp"
	"blog-backend/repo/mdb"
	"blog-backend/repo/rdb"
	"github.com/wpliap/common-wrap/elasticsearch"
	aftp "github.com/wpliap/common-wrap/ftp"
	agorm "github.com/wpliap/common-wrap/gorm"
	aredis "github.com/wpliap/common-wrap/redis"
)

func NewProxyService(opts ...Option) service.ProxyService {
	opt := &proxy{
		dbCli:      agorm.NewGormProxy("blog.mysql"),
		redisCli:   aredis.NewRedisProxy("blog.redis"),
		ftpCli:     aftp.NewFtpProxy("blog.ftp"),
		elasticCli: elasticsearch.NewElasticProxy("blog.elastic"),
	}
	for _, o := range opts {
		o(opt)
	}
	return opt
}

func (p *proxy) GetGormProxy() *mdb.MysqlClient {
	return &mdb.MysqlClient{DB: p.dbCli}
}

func (p *proxy) GetElasticProxy() *es.ElasticClient {
	return &es.ElasticClient{Client: p.elasticCli}
}

func (p *proxy) GetFtpProxy() *ftp.FtpClient {
	return &ftp.FtpClient{ServerConn: p.ftpCli}
}

func (p *proxy) GetRedisProxy() *rdb.RedisClient {
	return &rdb.RedisClient{UniversalClient: p.redisCli}
}
