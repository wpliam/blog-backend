package banner

import (
	"blog-backend/model"
	"blog-backend/repo/mdb"
	"github.com/gin-gonic/gin"
)

// GetBannerRsp 获取banner响应体
type GetBannerRsp struct {
	Banners []*model.Banner `json:"banners"`
}

// GetBanner 获取banner
func GetBanner(ctx *gin.Context) (interface{}, error) {
	banners, err := mdb.GetGormClient().GetBannerList()
	if err != nil {
		return nil, err
	}
	rsp := &GetBannerRsp{
		Banners: banners,
	}
	return rsp, nil
}
