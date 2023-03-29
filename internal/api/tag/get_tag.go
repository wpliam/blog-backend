package tag

import (
	"blog-backend/global/proxy"
	"blog-backend/model"
	"github.com/gin-gonic/gin"
)

type GetTagCard struct {
}

type GetTagCardReply struct {
	Tags []*model.Tag `json:"tags"`
}

func (t *GetTagCard) Invoke(ctx *gin.Context, proxy proxy.Proxy) (interface{}, error) {
	tag, err := proxy.GetGormProxy().GetTagList()
	if err != nil {
		return nil, err
	}
	rsp := &GetTagCardReply{
		Tags: tag,
	}
	return rsp, nil
}
