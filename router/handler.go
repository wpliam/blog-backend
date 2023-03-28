package router

import (
	"blog-backend/internal/api"
	"blog-backend/internal/api/article"
	"blog-backend/internal/common/client"
	"blog-backend/internal/common/container"
	"blog-backend/internal/common/proxy"
	"blog-backend/util/resp"
	"fmt"
	"github.com/gin-gonic/gin"
)

const (
	SearchArticleListName = "search_article_list"
	GetCardInfoName       = "get_card_info"
	ReadArticleName       = "read_article"
	GetArticleArchiveName = "get_article_archive"
)

// Enter 入口
type Enter interface {
	WriteFunc()
	Wrapper(key string) gin.HandlerFunc
}

// Carrier 载体
type Carrier struct {
	container.Factory
	proxy.Proxy
}

// New 创建一个入口
func New() Enter {
	return &Carrier{
		Factory: container.New(),
		Proxy:   proxy.New(),
	}
}

// WriteFunc 写函数
func (c *Carrier) WriteFunc() {
	c.Set(SearchArticleListName, &article.SearchArticle{})
	c.Set(ReadArticleName, &article.ReadArticle{})
	c.Set(GetArticleArchiveName, &article.GetArticleArchive{})
	c.Set(GetCardInfoName, &api.GetCard{})
}

// Wrapper 包装器
func (c *Carrier) Wrapper(key string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cli, ok := c.Get(key).(client.Client)
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
