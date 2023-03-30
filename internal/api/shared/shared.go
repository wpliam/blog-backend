package shared

import (
	"blog-backend/internal/service"
	"github.com/gin-gonic/gin"
)

func NewSharedService(proxyService service.ProxyService) service.SharedService {
	return &sharedImpl{
		proxyService,
	}
}

type sharedImpl struct {
	service.ProxyService
}

func (s *sharedImpl) AddViewCount(ctx *gin.Context) (interface{}, error) {
	return nil, nil
}

// GiveThumb 点赞
func (s *sharedImpl) GiveThumb(ctx *gin.Context) (interface{}, error) {
	var req *GiveThumbReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return s.GiveThumbImpl(ctx, req)
}

func (s *sharedImpl) GiveFollow(ctx *gin.Context) (interface{}, error) {
	return nil, nil
}

func (s *sharedImpl) GiveCollect(ctx *gin.Context) (interface{}, error) {
	return nil, nil
}
