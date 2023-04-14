package proxy

import (
	"blog-backend/repo/es"
	"blog-backend/repo/mdb"
	"blog-backend/repo/rdb"
)

type proxy struct {
	dbCli      *mdb.MysqlClient
	redisCli   *rdb.RedisClient
	elasticCli *es.ElasticClient
}

type Option func(*proxy)
