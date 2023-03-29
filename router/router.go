package router

import (
	"blog-backend/constant"
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
	apiGroup := r.Group("/api")
	{
		// 读取文章信息
		apiGroup.GET("read_article/:articleID", enter.Wrapper(constant.ReadArticleName))
		// 获取文章归档
		apiGroup.GET("get_article_archive", enter.Wrapper(constant.GetArticleArchiveName))
		// 搜索文章
		apiGroup.POST("search_article_list", enter.Wrapper(constant.SearchArticleListName))
		// 获取热门文章
		apiGroup.GET("get_hot_article", enter.Wrapper(constant.GetHotArticleName))
		// 获取banner卡片
		apiGroup.GET("get_banner_card", enter.Wrapper(constant.GetBannerCardName))
		// 获取分类卡片
		apiGroup.GET("get_category_card", enter.Wrapper(constant.GetCategoryCardName))
		// 获取标签卡片
		apiGroup.GET("get_tag_card", enter.Wrapper(constant.GetTagCardName))

		// 登录
		apiGroup.POST("login", enter.Wrapper(constant.LoginName))
	}
	return r
}
