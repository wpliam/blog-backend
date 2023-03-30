package article

import (
	"blog-backend/model"
	"github.com/gin-gonic/gin"
)

type GetHotArticle struct {
}

type GetArticleReply struct {
	Articles []*model.Article `json:"articles"`
}

// GetHotArticleImpl 获取热门文章
func (a *articleImpl) GetHotArticleImpl(ctx *gin.Context) (interface{}, error) {
	articles, err := a.GetGormProxy().GetHotArticle()
	if err != nil {
		return nil, err
	}
	rsp := GetArticleReply{
		Articles: articles,
	}
	return rsp, nil
}
