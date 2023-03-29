package article

import (
	"blog-backend/global/proxy"
	"blog-backend/model"
	"github.com/gin-gonic/gin"
	"github.com/wpliap/common-wrap/log"
)

type SearchArticle struct {
	Keyword string      `json:"keyword"`
	Cid     int64       `json:"cid"`
	TagID   int64       `json:"tagID"`
	Order   int         `json:"order"`
	Page    *model.Page `json:"page"`

	SearchType int `json:"searchType"` // 0:搜索文章 1:搜索随机文章
}

// SearchArticleReply 搜索文章响应体
type SearchArticleReply struct {
	Page     *model.Page                    `json:"page"`
	Articles []*model.ArticleContentSummary `json:"articles"`
}

func (s *SearchArticle) Invoke(ctx *gin.Context, proxy proxy.Proxy) (interface{}, error) {
	if err := ctx.ShouldBindJSON(&s); err != nil {
		return nil, err
	}
	switch s.SearchType {
	case 1:
		return s.SearchRandomArticle(ctx, proxy)
	default:
		return s.SearchArticleList(ctx, proxy)
	}
}

// SearchArticleList 搜索文章列表
func (s *SearchArticle) SearchArticleList(ctx *gin.Context, proxy proxy.Proxy) (*SearchArticleReply, error) {
	param := s.SearchArticleParam()
	log.Infof("param page:%+v", param.Page)
	articles, total, err := proxy.GetEsProxy().SearchArticleList(ctx, param)
	if err != nil {
		log.Errorf("SearchArticle search err:%v param:%+v", err, param)
		return nil, err
	}
	param.Page.SetTotal(total)
	rsp := &SearchArticleReply{
		Page:     param.Page,
		Articles: articles,
	}
	return rsp, nil
}

// SearchRandomArticle 搜索随机文章
func (s *SearchArticle) SearchRandomArticle(ctx *gin.Context, proxy proxy.Proxy) (*SearchArticleReply, error) {
	articles, err := proxy.GetEsProxy().SearchRandomArticle(ctx)
	if err != nil {
		return nil, err
	}
	rsp := &SearchArticleReply{
		Articles: articles,
	}
	return rsp, nil
}

// SearchArticleParam es检索参数转换
func (s *SearchArticle) SearchArticleParam() *model.SearchArticleParam {
	return &model.SearchArticleParam{
		Keyword: s.Keyword,
		Cid:     s.Cid,
		TagID:   s.TagID,
		Order:   s.Order,
		Page:    s.Page,
	}
}
