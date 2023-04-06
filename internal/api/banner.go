package api

import (
	"blog-backend/internal/service"
	"blog-backend/model/jsonagree"
	"github.com/gin-gonic/gin"
)

func NewBannerService(proxyService service.ProxyService) service.BannerService {
	return &bannerImpl{
		proxyService,
	}
}

type bannerImpl struct {
	service.ProxyService
}

// GetBannerCard 获取banner卡片
func (b *bannerImpl) GetBannerCard(ctx *gin.Context) (interface{}, error) {
	banners, err := b.GetGormProxy().GetBannerList()
	if err != nil {
		return nil, err
	}
	rsp := &jsonagree.GetBannerReply{
		Banners: banners,
	}
	return rsp, nil
}
