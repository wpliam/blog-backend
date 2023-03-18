package es

import (
	"github.com/olivere/elastic/v7"
	"github.com/wpliap/common-wrap/elasticsearch"
	"sync"
)

var (
	esProxy = make(map[string]*client)
	rw      sync.RWMutex
)

const (
	defaultElasticName = "elastic"
)

// Client es客户端
type client struct {
	*elastic.Client
}

// GetElasticClient 获取es
func GetElasticClient() *client {
	rw.RLock()
	defer rw.RUnlock()
	cli, ok := esProxy[defaultElasticName]
	if ok && cli != nil {
		return cli
	}
	esCli := elasticsearch.NewElasticProxy("blog.elastic")
	cli = &client{
		esCli,
	}
	esProxy[defaultElasticName] = cli
	return cli
}
