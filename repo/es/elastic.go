package es

import (
	"blog-backend/repo/config"
	"context"
	"github.com/olivere/elastic/v7"
)

// ElasticClient es客户端
type ElasticClient struct {
	cli *elastic.Client
}

func NewElasticClient() *ElasticClient {
	conf := config.GetElasticConf()
	url := conf.Addr
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(url),
		elastic.SetBasicAuth(conf.Username, conf.Password),
	)
	if err != nil {
		panic("elastic new client err " + err.Error())
	}
	_, _, err = client.Ping(url).Do(context.Background())
	if err != nil {
		panic("elastic ping err " + err.Error())
	}
	return &ElasticClient{
		cli: client,
	}
}
