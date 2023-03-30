package server

import (
	"blog-backend/internal/proxy"
	"github.com/gin-gonic/gin"
)

type Option func(*Server)

func defaultServerOption() *Server {
	return &Server{
		router:              gin.Default(),
		port:                8888,
		DisableServerRouter: false,
		MaxShutDownTimeout:  0,
		proxy:               proxy.NewProxyService(),
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
