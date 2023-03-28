package router

import (
	"blog-backend/middleware"
	"blog-backend/repo/auth/jwtauth"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	middle := &middleware.Middleware{
		Jwt: jwtauth.DefaultJwtAuth,
	}
	r.Use(middle.Options())
	enter := New()
	enter.WriteFunc()
	apiGroup := r.Group("/api")
	{
		// 获取卡片信息 1:获取分类卡片 2:获取标签卡片 3:获取banner
		apiGroup.GET("get_card_info/:cardType", enter.Wrapper(GetCardInfoName))
		// 读取文章信息
		apiGroup.GET("read_article/:articleID", enter.Wrapper(ReadArticleName))
		// 获取文章归档
		apiGroup.GET("get_article_archive", enter.Wrapper(GetArticleArchiveName))
		// 搜索文章
		apiGroup.POST("search_article_list", enter.Wrapper(SearchArticleListName))
	}
	return r
}
