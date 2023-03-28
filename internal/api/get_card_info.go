package api

import (
	"blog-backend/internal/common/proxy"
	"blog-backend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wpliap/common-wrap/log"
	"strconv"
)

type GetCard struct {
	CardType int64 `json:"cardType"`
}

func (c *GetCard) Invoke(ctx *gin.Context, proxy proxy.Proxy) (interface{}, error) {
	cardType, err := strconv.ParseInt(ctx.Param("cardType"), 10, 64)
	if err != nil {
		return nil, err
	}
	c.CardType = cardType
	return c.GetCardInfo(ctx, proxy)
}

// GetCardInfo 获取卡片信息
func (c *GetCard) GetCardInfo(ctx *gin.Context, proxy proxy.Proxy) (interface{}, error) {
	switch c.CardType {
	case 1:
		return c.GetCategoryCard(ctx, proxy)
	case 2:
		return c.GetTagCard(ctx, proxy)
	case 3:
		return c.GetBanner(ctx, proxy)
	default:
		return nil, fmt.Errorf("cardType err")
	}
}

// GetCategoryCardReply 分类卡片
type GetCategoryCardReply struct {
	CategoryCard []*model.CategoryCard `json:"categoryCard"`
}

// GetCategoryCard 获取分类卡片
func (c *GetCard) GetCategoryCard(ctx *gin.Context, proxy proxy.Proxy) (*GetCategoryCardReply, error) {
	categoryCard, err := proxy.GetGormProxy().GetCategoryCard()
	if err != nil {
		return nil, err
	}
	page := model.NewPage(1, 5)
	for _, category := range categoryCard {
		articles, total, err := proxy.GetEsProxy().SearchArticleList(ctx, &model.SearchArticleParam{Cid: category.Cid, Page: page})
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

type GetTagCardReply struct {
	Tags []*model.Tag `json:"tags"`
}

// GetTagCard 获取标签卡片
func (c *GetCard) GetTagCard(ctx *gin.Context, proxy proxy.Proxy) (*GetTagCardReply, error) {
	tag, err := proxy.GetGormProxy().GetTagList()
	if err != nil {
		return nil, err
	}
	rsp := &GetTagCardReply{
		Tags: tag,
	}
	return rsp, nil
}

// GetBannerReply banner
type GetBannerReply struct {
	Banners []*model.Banner `json:"banners"`
}

// GetBanner 获取banner
func (c *GetCard) GetBanner(ctx *gin.Context, proxy proxy.Proxy) (*GetBannerReply, error) {
	banners, err := proxy.GetGormProxy().GetBannerList()
	if err != nil {
		return nil, err
	}
	rsp := &GetBannerReply{
		Banners: banners,
	}
	return rsp, nil
}
