package banner

import (
	"blog-backend/global/proxy"
	"blog-backend/model"
	"github.com/gin-gonic/gin"
)

type GetBannerCard struct {
}

type GetBannerReply struct {
	Banners []*model.Banner `json:"banners"`
}

func (b *GetBannerCard) Invoke(ctx *gin.Context, proxy proxy.Proxy) (interface{}, error) {
	banners, err := proxy.GetGormProxy().GetBannerList()
	if err != nil {
		return nil, err
	}
	rsp := &GetBannerReply{
		Banners: banners,
	}
	return rsp, nil
}
