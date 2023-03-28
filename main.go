package main

import (
	"blog-backend/internal/common/proxy"
	"blog-backend/router"
	"context"
	"github.com/wpliap/common-wrap/log"
	"net/http"
	"time"

	_ "github.com/wpliap/common-wrap"
)

func main() {
	route := router.InitRouter()
	svr := &http.Server{
		Addr:         ":8888",
		Handler:      route,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	if err := svr.ListenAndServe(); err != nil {
		panic("server start err " + err.Error())
	}
}

func rebuild() {
	p := proxy.New()
	ctx := context.Background()
	if err := p.GetEsProxy().DeleteIndex(ctx); err != nil {
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
		if err := p.GetEsProxy().AddArticleToEs(ctx, article.ArticleContentSummary(tagID, recommendID)); err != nil {
			log.Errorf("AddArticleToEs err:%v", err)
		}
	}
}
