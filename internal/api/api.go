package api

import (
	"blog-backend/global/proxy"
	"github.com/gin-gonic/gin"
)

type Client interface {
	Invoke(ctx *gin.Context, proxy proxy.Proxy) (interface{}, error)
}
