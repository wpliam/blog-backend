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

// SearchArticleList es搜索文章
func (cli *ElasticClient) SearchArticleList(ctx context.Context, param *model.SearchArticleParam) ([]*model.ArticleContentSummary, int64, error) {
	searchService := cli.Search(constant.EsArticleIndex)
	query := elastic.NewBoolQuery()
	if param.Keyword != "" {
		query.Filter(elastic.NewWildcardQuery("title", param.Keyword))
	}
	if param.Cid > 0 {
		query.Filter(elastic.NewTermQuery("cid", param.Cid))
	}
	if param.TagID > 0 {
		query.Filter(elastic.NewBoolQuery().Must(elastic.NewTermQuery("tagIDs", param.TagID)))
	}
	searchService.Query(query)
	switch param.Order {
	case constant.SearchNewCreateArticle:
		searchService.Sort("createTime", false)
	case constant.SearchNewUpdateArticle:
		searchService.Sort("updateTime", false)
	}
	if param != nil && param.Page != nil {
		searchService.From((param.Page.Offset - 1) * param.Page.Limit).Size(param.Page.Limit)
	}
	searchResult, err := searchService.Do(ctx)
	if err != nil {
		return nil, 0, err
	}
	return convertArticleResult(searchResult), searchResult.TotalHits(), nil
}

// GetArticleInfo 根据id搜索文章信息
func (cli *ElasticClient) GetArticleInfo(ctx context.Context, articleID int64) (*model.ArticleContentSummary, error) {
	resp, err := cli.Get().Index(constant.EsArticleIndex).Id(fmt.Sprintf("%d", articleID)).Do(ctx)
	if err != nil {
		return nil, err
	}
	data, err := resp.Source.MarshalJSON()
	if err != nil {
		return nil, err
	}
	summary := &model.ArticleContentSummary{}
	if err = json.Unmarshal(data, summary); err != nil {
		return nil, err
	}
	return summary, nil
}

// GetArticleList 获取文章列表
func (cli *ElasticClient) GetArticleList(ctx context.Context, ids []int64) ([]*model.ArticleContentSummary, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	var multi []*elastic.MultiGetItem
	for _, id := range ids {
		query := elastic.NewMultiGetItem().Index(constant.EsArticleIndex).Id(fmt.Sprintf("%d", id))
		multi = append(multi, query)
	}
	resp, err := cli.MultiGet().Add(multi...).Do(ctx)
	if err != nil {
		return nil, err
	}
	var contentSummary []*model.ArticleContentSummary
	for _, doc := range resp.Docs {
		content, err := doc.Source.MarshalJSON()
		if err != nil {
			continue
		}
		summary := &model.ArticleContentSummary{}
		if err = json.Unmarshal(content, summary); err != nil {
			continue
		}
		contentSummary = append(contentSummary, summary)
	}
	return contentSummary, nil
}

// AddArticleToEs 添加文章到es
func (cli *ElasticClient) AddArticleToEs(ctx context.Context, article *model.ArticleContentSummary) error {
	exec := cli.Index().Index(constant.EsArticleIndex).Id(fmt.Sprintf("%d", article.ID)).BodyJson(article)
	resp, err := exec.Do(ctx)
	if err != nil {
		return err
	}
	log.Infof("InsertArticle resp:%+v", resp)
	return nil
}

// SearchRandomArticle 随机搜索文章
func (cli *ElasticClient) SearchRandomArticle(ctx context.Context) ([]*model.ArticleContentSummary, error) {
	searchResult, err := cli.Search(constant.EsArticleIndex).
		SortBy(elastic.NewScriptSort(elastic.NewScript("Math.random()"), "number")).
		From(0).
		Size(10).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	return convertArticleResult(searchResult), nil
}

// DeleteIndex 删除索引
func (cli *ElasticClient) DeleteIndex(ctx context.Context) error {
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

func (cli *ElasticClient) AggregationsArticleCategory(ctx context.Context) (*ArticleCategoryAggregations, error) {
	// 根据category.cid分组
	cidGroup := elastic.NewTermsAggregation().Field("cid")

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
func (cli *ElasticClient) UpdateArticle(ctx context.Context, articleID int64, field map[string]interface{}) error {
	_, err := cli.Update().Index(constant.EsArticleIndex).
		Id(fmt.Sprintf("%d", articleID)).Doc(field).Do(ctx)
	return err
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
