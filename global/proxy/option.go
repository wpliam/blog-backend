package proxy

import (
	"blog-backend/repo/es"
	"blog-backend/repo/ftp"
	"blog-backend/repo/mdb"
	"blog-backend/repo/rdb"
)

type Options struct {
	MysqlClient *mdb.MysqlClient
	RedisClient *rdb.RedisClient
	FtpClient   *ftp.FtpClient
	EsClient    *es.ElasticClient
}

type Option func(*Options)

func WithGormProxy(mysqlCli *mdb.MysqlClient) Option {
	return func(opt *Options) {
		opt.MysqlClient = mysqlCli
	}
}

func WithRedisProxy(redisCli *rdb.RedisClient) Option {
	return func(opt *Options) {
		opt.RedisClient = redisCli
	}
}

func WithFtpProxy(ftpCli *ftp.FtpClient) Option {
	return func(opt *Options) {
		opt.FtpClient = ftpCli
	}
}

func WithEsProxy(esProxy *es.ElasticClient) Option {
	return func(opt *Options) {
		opt.EsClient = esProxy
	}
}

func (opt *Options) GetGormProxy() *mdb.MysqlClient {
	return opt.MysqlClient
}

func (opt *Options) GetRedisProxy() *rdb.RedisClient {
	return opt.RedisClient
}

func (opt *Options) GetFtpProxy() *ftp.FtpClient {
	return opt.FtpClient
}

func (opt *Options) GetEsProxy() *es.ElasticClient {
	return opt.EsClient
}
