package article

import (
	"blog-backend/model"
	"github.com/gin-gonic/gin"
	"github.com/wpliap/common-wrap/log"
)

type SearchArticleReq struct {
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

// SearchArticleListImpl 搜索文章列表
func (a *articleImpl) SearchArticleListImpl(ctx *gin.Context, req *SearchArticleReq) (*SearchArticleReply, error) {
	param := a.SearchArticleParam(req)
	log.Infof("param page:%+v", param.Page)
	articles, total, err := a.GetElasticProxy().SearchArticleList(ctx, param)
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
func (a *articleImpl) SearchRandomArticle(ctx *gin.Context) (*SearchArticleReply, error) {
	articles, err := a.GetElasticProxy().SearchRandomArticle(ctx)
	if err != nil {
		return nil, err
	}
	rsp := &SearchArticleReply{
		Articles: articles,
	}
	return rsp, nil
}

// SearchArticleParam es检索参数转换
func (a *articleImpl) SearchArticleParam(req *SearchArticleReq) *model.SearchArticleParam {
	return &model.SearchArticleParam{
		Keyword: req.Keyword,
		Cid:     req.Cid,
		TagID:   req.TagID,
		Order:   req.Order,
		Page:    req.Page,
	}
}
