package article

import (
	"blog-backend/constant"
	"blog-backend/model"
	"blog-backend/repo/es"
	"blog-backend/util"
	"github.com/gin-gonic/gin"
)

type GetArticleArchiveRsp struct {
	Article map[string][]*model.ArticleContentSummary `json:"article"`
	Count   int64                                     `json:"count"`
}

// GetArticleArchive 文章归档
func GetArticleArchive(ctx *gin.Context) (interface{}, error) {
	searchResult, err := es.GetElasticClient().SearchAllArticle(ctx)
	if err != nil {
		return nil, err
	}
	articleList := convertArticleResult(searchResult)
	articleGroupBy(articleList)
	rsp := GetArticleArchiveRsp{
		Article: articleGroupBy(articleList),
		Count:   searchResult.TotalHits(),
	}
	return rsp, nil
}

func articleGroupBy(articleList []*model.ArticleContentSummary) map[string][]*model.ArticleContentSummary {
	articleGroup := make(map[string][]*model.ArticleContentSummary)
	for _, item := range articleList {
		key := util.ParseDateTime(constant.MonthSubTableSuffix, item.CreateTime)
		articleGroup[key] = append(articleGroup[key], item)
	}
	return articleGroup
}
