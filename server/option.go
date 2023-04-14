package server

import (
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
		router:              gin.Default(),
		port:                8888,
		DisableServerRouter: false,
		MaxShutDownTimeout:  10,
		proxy:               proxyService,
		middle: &middleware.Middleware{
			Jwt:        jwtauth.DefaultJwtAuth,
			RedisProxy: proxyService.GetRedisProxy(),
		},
	}
}
