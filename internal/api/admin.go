package api

import (
	"blog-backend/internal/service"
	"github.com/gin-gonic/gin"
)

func NewAdminService(proxyService service.ProxyService) service.AdminService {
	return &adminImpl{
		proxyService,
	}
}

type adminImpl struct {
	service.ProxyService
}

func (a *adminImpl) ArticleReview(ctx *gin.Context) error {
	return nil
}

func (a *adminImpl) GetReadyReviewArticle(ctx *gin.Context) (interface{}, error) {

	return nil, nil
}
