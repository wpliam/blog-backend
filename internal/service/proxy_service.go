package service

import (
	"blog-backend/repo/es"
	"blog-backend/repo/mdb"
	"blog-backend/repo/rdb"
)

type ProxyService interface {
	GetGormProxy() *mdb.MysqlClient
	GetElasticProxy() *es.ElasticClient
	GetRedisProxy() *rdb.RedisClient
}
