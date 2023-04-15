package server

import (
	"blog-backend/internal/api"
	"blog-backend/internal/file"
	"blog-backend/internal/proxy"
	"blog-backend/middleware"
	"blog-backend/repo/auth/jwtauth"
	"github.com/gin-gonic/gin"
	"io"
)

type Option func(*Server)

func defaultServerOption() *Server {
	proxyService := proxy.NewProxyService()
	gin.SetMode(gin.ReleaseMode)   // 生产模式
	gin.DefaultWriter = io.Discard // 禁用gin输出接口访问日志
	return &Server{
		router: gin.Default(),
		middle: &middleware.Middleware{
			Jwt:        jwtauth.DefaultJwtAuth,
			RedisProxy: proxyService.GetRedisProxy(),
		},
		MaxShutDownTimeout: 10,

		proxyService: proxyService,

		articleService:  api.NewArticleService(proxyService),
		bannerService:   api.NewBannerService(proxyService),
		categoryService: api.NewCategoryService(proxyService),
		commentService:  api.NewCommentService(proxyService),
		sharedService:   api.NewSharedService(proxyService),
		tagService:      api.NewTagService(proxyService),
		userService:     api.NewUserService(proxyService),

		adminService: api.NewAdminService(proxyService),

		uploadService:   file.NewUploadService(),
		downloadService: file.NewDownloadService(),
	}
}
