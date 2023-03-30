package article

import (
	"blog-backend/constant"
	"blog-backend/model"
	"blog-backend/util/thread"
	"github.com/gin-gonic/gin"
)

type GetArticleArchiveReq struct {
}

type GetArticleArchiveReply struct {
	Article       map[string][]*model.ArticleContentSummary `json:"article"`
	ArticleCount  int64                                     `json:"articleCount"`
	Tags          []*model.Tag                              `json:"tags"`
	TagCount      int64                                     `json:"tagCount"`
	Category      []*model.Category                         `json:"category"`
	CategoryCount int64                                     `json:"categoryCount"`
}

// GetArticleArchiveImpl 文章归档
func (a *articleImpl) GetArticleArchiveImpl(ctx *gin.Context) (*GetArticleArchiveReply, error) {
	rsp := &GetArticleArchiveReply{}
	handler := make([]func() error, 0)
	handler = append(handler, func() error {
		articles, total, err := a.GetElasticProxy().SearchArticleList(ctx, &model.SearchArticleParam{})
		if err != nil {
			return err
		}
		rsp.Article = articleGroupBy(articles)
		rsp.ArticleCount = total
		return nil
	})
	handler = append(handler, func() error {
		tags, err := a.GetGormProxy().GetTagList()
		if err != nil {
			return err
		}
		rsp.Tags = tags
		rsp.TagCount = int64(len(tags))
		return nil
	})
	handler = append(handler, func() error {
		categoryList, err := a.GetGormProxy().GetCategoryList()
		if err != nil {
			return err
		}
		rsp.Category = categoryList
		rsp.CategoryCount = int64(len(categoryList))
		return nil
	})
	if err := thread.GoAndWait(handler...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func articleGroupBy(articleList []*model.ArticleContentSummary) map[string][]*model.ArticleContentSummary {
	articleGroup := make(map[string][]*model.ArticleContentSummary)
	for _, article := range articleList {
		key := article.CreateTime.Format(constant.MonthSubTableSuffix)
		articleGroup[key] = append(articleGroup[key], article)
	}
	return articleGroup
}
