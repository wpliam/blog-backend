package category

import (
	"blog-backend/model"
	"github.com/gin-gonic/gin"
	"github.com/wpliap/common-wrap/log"
)

type GetCategoryCard struct {
}

// GetCategoryCardReply 分类卡片
type GetCategoryCardReply struct {
	CategoryCard []*model.CategoryCard `json:"categoryCard"`
}

func (c *categoryImpl) GetCategoryCardImpl(ctx *gin.Context) (*GetCategoryCardReply, error) {
	categoryCard, err := c.GetGormProxy().GetCategoryCard()
	if err != nil {
		return nil, err
	}
	page := model.NewPage(1, 5)
	for _, category := range categoryCard {
		articles, total, err := c.GetElasticProxy().SearchArticleList(ctx, &model.SearchArticleParam{Cid: category.Cid, Page: page})
		if err != nil {
			log.Errorf("GetCategoryCard SearchArticleList err:%v", err)
			continue
		}
		if len(articles) == 0 {
			continue
		}
		category.Cover = articles[0].CategoryCover
		category.CategoryName = articles[0].CategoryName
		category.Total = total
		category.Articles = articles
	}
	rsp := &GetCategoryCardReply{
		CategoryCard: categoryCard,
	}
	return rsp, nil
}
