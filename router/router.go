package router

import (
	"blog-backend/internal/api/article"
	"blog-backend/internal/api/banner"
	"blog-backend/internal/api/tag"
	"blog-backend/middleware"
	"blog-backend/repo/auth/jwtauth"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	middle := &middleware.Middleware{
		JwtAuth: jwtauth.DefaultJwtAuth,
	}
	r.Use(middle.Options())
	apiGroup := r.Group("/api")
	{
		apiGroup.POST("search_article", WrapHandler(article.SearchArticle))
		apiGroup.GET("search_random_article", WrapHandler(article.SearchRandomArticle))
		apiGroup.POST("read_article", WrapHandler(article.ReadArticle))
		apiGroup.POST("aggregation_article_category", WrapHandler(article.AggregationArticleCategory))

		apiGroup.POST("get_banner", WrapHandler(banner.GetBanner))
		apiGroup.POST("get_tag", WrapHandler(tag.GetTag))
	}
	return r
}
