package es

import (
	"github.com/olivere/elastic/v7"
	"github.com/wpliap/common-wrap/elasticsearch"
)

// ElasticClient es客户端
type ElasticClient struct {
	*elastic.Client
}

func NewElasticClient(name string) *ElasticClient {
	return &ElasticClient{
		elasticsearch.NewElasticProxy(name),
	}
}
