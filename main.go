package main

import (
	"blog-backend/internal/proxy"
	"blog-backend/server"
	"context"
	"github.com/wpliap/common-wrap/log"

	_ "github.com/wpliap/common-wrap"
)

func main() {
	s := server.NewServer()
	s.Run()
}

func rebuild() {
	p := proxy.NewProxyService()
	ctx := context.Background()
	if err := p.GetElasticProxy().DeleteIndex(ctx); err != nil {
		log.Errorf("DeleteIndex err:%v", err)
		return
	}
	articles, err := p.GetGormProxy().GetAllArticle()
	if err != nil {
		log.Errorf("GetAllArticle err:%v", err)
		return
	}
	for _, article := range articles {
		tagID, err := p.GetGormProxy().GetArticleTagID(article.ID)
		if err != nil {
			log.Errorf("GetArticleTagID err:%v", err)
			continue
		}
		recommendID, err := p.GetGormProxy().GetRecommendID(article.ID)
		if err != nil {
			log.Errorf("GetRecommendID err:%v", err)
			continue
		}
		if err := p.GetElasticProxy().AddArticleToEs(ctx, article.ArticleContentSummary(tagID, recommendID)); err != nil {
			log.Errorf("AddArticleToEs err:%v", err)
		}
	}
}
