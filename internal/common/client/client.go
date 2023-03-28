package client

import (
	"blog-backend/internal/common/proxy"
	"github.com/gin-gonic/gin"
)

type Client interface {
	Invoke(ctx *gin.Context, proxy proxy.Proxy) (interface{}, error)
}
