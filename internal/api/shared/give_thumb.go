package shared

import "github.com/gin-gonic/gin"

type GiveThumbReq struct {
	ID       int64  `json:"id"`
	LikeType uint32 `json:"likeType"` // 0:文章 1:评论
}
type GiveThumbReply struct {
}

func (s *sharedImpl) GiveThumbImpl(ctx *gin.Context, req *GiveThumbReq) (*GiveThumbReply, error) {

	return nil, nil
}
