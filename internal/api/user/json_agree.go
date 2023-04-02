package user

import (
	"blog-backend/model"
)

type GetUserInfoReply struct {
	User     *model.User `json:"user"`
	IsFollow bool        `json:"isFollow"`
}

type GetUserCollectListReq struct {
	Uid int64 `json:"uid" binding:"min=1"`
}

type GetUserCollectListRsp struct {
	Articles []*model.ArticleContentSummary `json:"articles"`
}
