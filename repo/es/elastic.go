package es

import (
	"github.com/olivere/elastic/v7"
)

// ElasticClient es客户端
type ElasticClient struct {
	*elastic.Client
}
