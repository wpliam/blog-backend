package article

import (
	"blog-backend/internal/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func NewArticleService(proxyService service.ProxyService) service.ArticleService {
	return &articleImpl{
		proxyService,
	}
}

type articleImpl struct {
	service.ProxyService
}

// SearchArticleList 搜索文章列表
func (a *articleImpl) SearchArticleList(ctx *gin.Context) (interface{}, error) {
	var req *SearchArticleReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	switch req.SearchType {
	case 1:
		return a.SearchRandomArticle(ctx)
	default:
		return a.SearchArticleListImpl(ctx, req)
	}
}

// GetArticleArchive 获取文章归档
func (a *articleImpl) GetArticleArchive(ctx *gin.Context) (interface{}, error) {
	return a.GetArticleArchiveImpl(ctx)
}

// GetHotArticle 获取热门文章
func (a *articleImpl) GetHotArticle(ctx *gin.Context) (interface{}, error) {
	return a.GetHotArticleImpl(ctx)
}

// ReadArticle 读取文章
func (a *articleImpl) ReadArticle(ctx *gin.Context) (interface{}, error) {
	articleID, err := strconv.ParseInt(ctx.Param("articleID"), 10, 64)
	if err != nil {
		return nil, err
	}
	return a.ReadArticleImpl(ctx, articleID)
}
