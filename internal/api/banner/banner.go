package banner

import (
	"blog-backend/internal/service"
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
	return b.GetBannerCardImpl(ctx)
}
