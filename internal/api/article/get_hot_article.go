package article

import (
	"blog-backend/global/proxy"
	"blog-backend/model"
	"github.com/gin-gonic/gin"
)

type GetHotArticle struct {
}

type GetArticleReply struct {
	Articles []*model.Article `json:"articles"`
}

// Invoke 获取热门文章
func (a *GetHotArticle) Invoke(ctx *gin.Context, proxy proxy.Proxy) (interface{}, error) {
	articles, err := proxy.GetGormProxy().GetHotArticle()
	if err != nil {
		return nil, err
	}
	rsp := GetArticleReply{
		Articles: articles,
	}
	return rsp, nil
}
