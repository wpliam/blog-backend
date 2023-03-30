package server

import (
	"blog-backend/internal/proxy"
	"blog-backend/middleware"
	"blog-backend/repo/auth/jwtauth"
	"github.com/gin-gonic/gin"
)

type Option func(*Server)

func defaultServerOption() *Server {
	proxyService := proxy.NewProxyService()
	return &Server{
		router:              gin.Default(),
		port:                8888,
		DisableServerRouter: false,
		MaxShutDownTimeout:  0,
		proxy:               proxyService,
		middle: &middleware.Middleware{
			Jwt:        jwtauth.DefaultJwtAuth,
			RedisProxy: proxyService.GetRedisProxy(),
		},
	}
}

// WithPort 端口设置
func WithPort(port uint16) Option {
	return func(s *Server) {
		s.port = port
	}
}

// WithServerRoute 传了路由将自己的路由禁用
func WithServerRoute(router *gin.Engine) Option {
	return func(s *Server) {
		s.router = router
		s.DisableServerRouter = true
	}
}
