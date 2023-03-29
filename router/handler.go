package router

import (
	"blog-backend/global/container"
	"blog-backend/global/proxy"
	"blog-backend/internal/api"
	"blog-backend/util/resp"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Carrier 载体
type Carrier struct {
	proxy.Proxy
}

// New ...
func New() *Carrier {
	enter := &Carrier{
		Proxy: proxy.New(),
	}
	return enter
}

// Wrapper 包装器
func (c *Carrier) Wrapper(key string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cli, ok := container.DefaultContainer.Get(key).(api.Client)
		if !ok {
			panic(fmt.Sprintf("key %s not exist", key))
		}
		data, err := cli.Invoke(ctx, c.Proxy)
		if err != nil {
			resp.ResponseFail(ctx, err)
			return
		}
		resp.ResponseOk(ctx, data)
	}
}
