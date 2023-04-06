package api

import (
	"blog-backend/internal/service"
	"blog-backend/model"
	"blog-backend/model/jsonagree"
	"github.com/gin-gonic/gin"
	"github.com/wpliap/common-wrap/log"
)

func NewCategoryService(proxyService service.ProxyService) service.CategoryService {
	return &categoryImpl{
		proxyService,
	}
}

type categoryImpl struct {
	service.ProxyService
}

// GetCategoryCard 获取分类卡片
func (c *categoryImpl) GetCategoryCard(ctx *gin.Context) (interface{}, error) {
	categoryCard, err := c.GetGormProxy().GetCategoryCard()
	if err != nil {
		return nil, err
	}
	page := model.NewPage(1, 5)
	for _, category := range categoryCard {
		req := &jsonagree.SearchArticleListReq{Cid: category.Cid, Page: page}
		articles, total, err := c.GetElasticProxy().SearchArticleList(ctx, req)
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
	rsp := &jsonagree.GetCategoryCardReply{
		CategoryCard: categoryCard,
	}
	return rsp, nil
}

// GetCategoryList 获取分类列表
func (c *categoryImpl) GetCategoryList(ctx *gin.Context) (interface{}, error) {
	categoryList, err := c.GetGormProxy().GetCategoryList()
	if err != nil {
		return nil, err
	}
	rsp := &jsonagree.GetCategoryListReply{
		CategoryList: categoryList,
	}
	return rsp, nil
}
