package main

import (
	"blog-backend/constant"
	"blog-backend/global/container"
	"blog-backend/global/proxy"
	"blog-backend/internal/api/article"
	"blog-backend/internal/api/banner"
	"blog-backend/internal/api/category"
	"blog-backend/internal/api/tag"
	"blog-backend/internal/api/user"
	"blog-backend/router"
	"context"
	"github.com/wpliap/common-wrap/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/wpliap/common-wrap"
)

func init() {
	defaultContainer := container.DefaultContainer
	defaultContainer.Set(constant.SearchArticleListName, &article.SearchArticle{})
	defaultContainer.Set(constant.ReadArticleName, &article.ReadArticle{})
	defaultContainer.Set(constant.GetArticleArchiveName, &article.GetArticleArchive{})
	defaultContainer.Set(constant.GetHotArticleName, &article.GetHotArticle{})
	defaultContainer.Set(constant.GetBannerCardName, &banner.GetBannerCard{})
	defaultContainer.Set(constant.GetCategoryCardName, &category.GetCategoryCard{})
	defaultContainer.Set(constant.GetTagCardName, &tag.GetTagCard{})

	defaultContainer.Set(constant.LoginName, &user.Login{})
}

func main() {
	route := router.InitRouter()
	svr := &http.Server{
		Addr:         ":8888",
		Handler:      route,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	go func() {
		if err := svr.ListenAndServe(); err != nil {
			log.Errorf("ListenAndServe err %v", err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	switch <-c {
	case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT:
		if err := svr.Shutdown(context.Background()); err != nil {
			log.Errorf("server Shutdown err %v", err)
		}
	default:
		log.Infof("未知信号")
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
