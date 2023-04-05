package tag

import "blog-backend/model"

type GetTagListReq struct {
}

type GetTagListReply struct {
	Tags []*model.Tag `json:"tags"`
}
