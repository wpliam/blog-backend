package main

import (
	"blog-backend/repo/es"
	"blog-backend/repo/mdb"
	"blog-backend/router"
	"context"
	_ "github.com/wpliap/common-wrap"
	"github.com/wpliap/common-wrap/log"
)

func main() {
	route := router.InitRouter()
	if err := route.Run(":8888"); err != nil {
		panic("server start err " + err.Error())
	}
}

func renovate() {
	ctx := context.Background()
	esCli := es.GetElasticClient()
	if err := esCli.DeleteIndex(ctx); err != nil {
		log.Errorf("DeleteIndex err:%v", err)
		return
	}
	dbCli := mdb.GetGormClient()
	articles, err := dbCli.GetAllArticle()
	if err != nil {
		log.Errorf("GetAllArticle err:%v", err)
		return
	}
	for _, item := range articles {
		if err = esCli.AddArticleToEs(ctx, item.ArticleContentSummary()); err != nil {
			log.Errorf("AddArticleToEs err:%v", err)
		}
	}
}
