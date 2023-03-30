package category

import (
	"blog-backend/internal/service"
	"github.com/gin-gonic/gin"
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
	return c.GetCategoryCardImpl(ctx)
}
