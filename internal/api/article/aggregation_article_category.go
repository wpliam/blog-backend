package article

import (
	"blog-backend/constant"
	"blog-backend/model"
	"blog-backend/repo/es"
	"github.com/gin-gonic/gin"
	"github.com/wpliap/common-wrap/log"
)

type AggregationArticleCategoryRsp struct {
	Aggregations []*AggregationCategory `json:"aggregations"`
}

type AggregationCategory struct {
	Cid          int64                          `json:"cid"`
	CategoryName string                         `json:"categoryName"`
	Total        int64                          `json:"total"`
	ViewCount    int64                          `json:"viewCount"`
	Articles     []*model.ArticleContentSummary `json:"articles"`
}

// AggregationArticleCategory 聚合文章分类
func AggregationArticleCategory(ctx *gin.Context) (interface{}, error) {
	esCli := es.GetElasticClient()
	aggregations, err := esCli.AggregationsArticleCategory(ctx)
	if err != nil {
		return nil, err
	}

	var aggregationCategory []*AggregationCategory
	for _, item := range aggregations.Aggregations.CategoryGroup.Buckets {
		articleInfo := &model.ArticleContentSummary{}
		articleInfo.Category.Cid = item.Key
		searchResult, err := esCli.SearchArticleList(ctx, articleInfo, constant.SearchHotArticle, &model.Page{
			Offset: 1,
			Limit:  5,
		})
		if err != nil {
			log.Errorf("SearchArticleList err:%v cid:%d", err, item.Key)
			continue
		}

		aggregationCategory = append(aggregationCategory, &AggregationCategory{
			Cid:       item.Key,
			Total:     item.DocCount,
			ViewCount: int64(item.ViewCount.Value),
			Articles:  convertArticleResult(searchResult),
		})
	}
	for _, category := range aggregationCategory {
		for _, item := range category.Articles {
			category.CategoryName = item.Category.CategoryName
		}
	}
	rsp := &AggregationArticleCategoryRsp{
		Aggregations: aggregationCategory,
	}
	return rsp, nil
}
