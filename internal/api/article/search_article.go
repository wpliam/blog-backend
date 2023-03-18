package article

import (
	"blog-backend/model"
	"blog-backend/repo/es"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/wpliap/common-wrap/log"
)

// SearchArticleReq 搜索文章的请求体
type SearchArticleReq struct {
	Article *model.ArticleContentSummary `json:"article"`
	Page    *model.Page                  `json:"page"`
	Order   int                          `json:"order"`
}

// SearchArticleRsp 搜索文章的响应体
type SearchArticleRsp struct {
	Page     *model.Page                    `json:"page"`
	Articles []*model.ArticleContentSummary `json:"articles"`
}

// SearchArticle 搜索文章
func SearchArticle(ctx *gin.Context) (interface{}, error) {
	var req *SearchArticleReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Errorf("SearchArticle ShouldBindJSON err:%v req:%+v", err, req)
		return nil, err
	}
	esClient := es.GetElasticClient()
	searchResult, err := esClient.SearchArticleList(ctx, req.Article, req.Order, req.Page)
	if err != nil {
		log.Errorf("SearchArticle search err:%v req:%+v", err, req)
		return nil, err
	}
	req.Page.SetTotal(searchResult.TotalHits())
	rsp := &SearchArticleRsp{
		Page:     req.Page,
		Articles: convertArticleResult(searchResult),
	}
	return rsp, nil
}

// SearchRandomArticle 搜索随机文章
func SearchRandomArticle(ctx *gin.Context) (interface{}, error) {
	esClient := es.GetElasticClient()
	searchResult, err := esClient.SearchRandomArticle(ctx)
	if err != nil {
		return nil, err
	}
	rsp := &SearchArticleRsp{
		Articles: convertArticleResult(searchResult),
	}
	return rsp, nil
}

func convertArticleResult(searchResult *elastic.SearchResult) []*model.ArticleContentSummary {
	var articles []*model.ArticleContentSummary
	for _, hit := range searchResult.Hits.Hits {
		articleInfo := &model.ArticleContentSummary{}
		if err := json.Unmarshal(hit.Source, articleInfo); err != nil {
			log.Errorf("SearchArticle Source err:%v", err)
			continue
		}
		articles = append(articles, articleInfo)
	}
	return articles
}
