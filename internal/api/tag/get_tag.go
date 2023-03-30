package tag

import (
	"blog-backend/model"
	"github.com/gin-gonic/gin"
)

type GetTagListReq struct {
}

type GetTagListReply struct {
	Tags []*model.Tag `json:"tags"`
}

func (t *tagImpl) GetTagListImpl(ctx *gin.Context) (*GetTagListReply, error) {
	tag, err := t.GetGormProxy().GetTagList()
	if err != nil {
		return nil, err
	}
	rsp := &GetTagListReply{
		Tags: tag,
	}
	return rsp, nil
}
