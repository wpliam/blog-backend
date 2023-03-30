package banner

import (
	"blog-backend/model"
	"github.com/gin-gonic/gin"
)

type GetBannerCard struct {
}

type GetBannerReply struct {
	Banners []*model.Banner `json:"banners"`
}

func (b *bannerImpl) GetBannerCardImpl(ctx *gin.Context) (*GetBannerReply, error) {
	banners, err := b.GetGormProxy().GetBannerList()
	if err != nil {
		return nil, err
	}
	rsp := &GetBannerReply{
		Banners: banners,
	}
	return rsp, nil
}
