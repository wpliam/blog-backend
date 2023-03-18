package es

import (
	"blog-backend/constant"
	"blog-backend/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/wpliap/common-wrap/log"
)

// SearchAllArticle 搜索所有文章
func (cli *client) SearchAllArticle(ctx context.Context) (*elastic.SearchResult, error) {
	return cli.SearchArticleList(ctx, nil, 0, nil)
}

// SearchArticleList es搜索文章
func (cli *client) SearchArticleList(ctx context.Context,
	article *model.ArticleContentSummary, order int, page *model.Page) (*elastic.SearchResult, error) {
	searchService := cli.Search(constant.EsArticleIndex)
	articleQueryCondition(searchService, article)
	articleOrderCondition(searchService, order)
	if page != nil && page.Limit > 0 && page.Offset > 0 {
		searchService.From((page.Offset - 1) * page.Limit).Size(page.Limit)
	}
	return searchService.Do(ctx)
}

func articleQueryCondition(searchService *elastic.SearchService, article *model.ArticleContentSummary) {
	if article == nil {
		return
	}
	query := elastic.NewBoolQuery()
	if article.Title != "" {
		query.Filter(elastic.NewWildcardQuery("title", article.Title))
	}
	if article.Category.Cid > 0 {
		query.Filter(elastic.NewTermQuery("category.cid", article.Category.Cid))
	}
	if article.Category.CategoryName != "" {
		query.Filter(elastic.NewTermQuery("category.categoryName", article.Category.CategoryName))
	}
	if article.User.Uid > 0 {
		query.Filter(elastic.NewTermQuery("user.uid", article.User.Uid))
	}
	searchService.Query(query)
}

func articleOrderCondition(searchService *elastic.SearchService, order int) {
	switch order {
	case constant.SearchNewArticle:
		searchService.Sort("createTime", false)
	case constant.SearchHotArticle:
		searchService.Sort("viewCount", false)
	case constant.SearchLikeArticle:
		searchService.Sort("likeCount", false)
	}
}

// AddArticleToEs 添加文章到es
func (cli *client) AddArticleToEs(ctx context.Context, article *model.ArticleContentSummary) error {
	exec := cli.Index().Index(constant.EsArticleIndex).Id(fmt.Sprintf("%d", article.ID)).BodyJson(article)
	resp, err := exec.Do(ctx)
	if err != nil {
		return err
	}
	log.Infof("InsertArticle resp:%+v", resp)
	return nil
}

// SearchRandomArticle 随机搜索文章
func (cli *client) SearchRandomArticle(ctx context.Context) (*elastic.SearchResult, error) {
	return cli.Search(constant.EsArticleIndex).
		SortBy(elastic.NewScriptSort(elastic.NewScript("Math.random()"), "number")).
		From(0).
		Size(1).
		Do(ctx)
}

// DeleteIndex 删除索引
func (cli *client) DeleteIndex(ctx context.Context) error {
	_, err := cli.Client.DeleteIndex(constant.EsArticleIndex).Do(ctx)
	return err
}

type ArticleCategoryAggregations struct {
	Aggregations struct {
		CategoryGroup struct {
			Buckets []struct {
				Key       int64 `json:"key"`
				DocCount  int64 `json:"doc_count"`
				ViewCount struct {
					Value float64 `json:"value"`
				} `json:"view_count"`
			} `json:"buckets"`
		} `json:"categoryGroup"`
	} `json:"aggregations"`
}

func (cli *client) AggregationsArticleCategory(ctx context.Context) (*ArticleCategoryAggregations, error) {
	// 根据category.cid分组
	cidGroup := elastic.NewTermsAggregation().Field("category.cid")

	searchSource := elastic.NewSearchSource()

	viewCount := elastic.NewSumAggregation().Field("viewCount")
	// viewCount sum值
	cidGroup.SubAggregation("view_count", viewCount)
	// view_count 降序排列
	cidGroup.OrderByAggregation("view_count", false)

	bucketSort := elastic.NewBucketSortAggregation().From(0).Size(2)
	cidGroup.SubAggregation("bucket_field", bucketSort)

	searchSource.Aggregation("categoryGroup", cidGroup)
	result, err := cli.Search().Index(constant.EsArticleIndex).SearchSource(searchSource).Do(ctx)
	if err != nil {
		return nil, err
	}

	content, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	articleCategoryGroup := &ArticleCategoryAggregations{}
	if err = json.Unmarshal(content, articleCategoryGroup); err != nil {
		return nil, err
	}
	return articleCategoryGroup, nil
}

// UpdateArticle 更新文章
func (cli *client) UpdateArticle(ctx context.Context, articleID int64, field map[string]interface{}) error {
	_, err := cli.Update().Index(constant.EsArticleIndex).
		Id(fmt.Sprintf("%d", articleID)).Doc(field).Do(ctx)
	return err
}
