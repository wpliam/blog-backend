package tag

import (
	"blog-backend/internal/service"
	"github.com/gin-gonic/gin"
)

func NewTagService(proxyService service.ProxyService) service.TagService {
	return &tagImpl{
		proxyService,
	}
}

type tagImpl struct {
	service.ProxyService
}

func (t *tagImpl) GetTagList(ctx *gin.Context) (interface{}, error) {
	tag, err := t.GetGormProxy().GetTagList()
	if err != nil {
		return nil, err
	}
	rsp := &GetTagListReply{
		Tags: tag,
	}
	return rsp, nil
}
