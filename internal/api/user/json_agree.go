package user

import (
	"blog-backend/model"
)

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginReply struct {
	Token string      `json:"token"`
	User  *model.User `json:"user"`
}

// GetUserInfoReply 获取用户信息
type GetUserInfoReply struct {
	User     *model.User `json:"user"`     // 用户基本信息
	IsFollow bool        `json:"isFollow"` // 是否关注
	IsClock  bool        `json:"isClock"`  // 是否签到
}

type GetUserCollectListReq struct {
	Uid int64 `json:"uid" binding:"min=1"`
}

type GetUserCollectListRsp struct {
	Articles []*model.ArticleContentSummary `json:"articles"`
}

type CensusUserInfoReply struct {
	ArticleCount int64 `json:"articleCount"` // 用户发表了多少篇文章
	CommentCount int64 `json:"commentCount"` // 用户发表了多少评论
	HotCount     int64 `json:"hotCount"`     // 用户人气值

	FansCount    int64 `json:"fansCount"`    // 用户有多少粉丝
	CollectCount int64 `json:"collectCount"` // 用户收藏文章数
	LikeCount    int64 `json:"likeCount"`    // 获得多少点赞
}

type RefreshTokenReq struct {
	Uid   int64  `json:"uid" binding:"min=1"`
	Token string `json:"token" binding:"required"`
}

type RefreshTokenReply struct {
	Token string `json:"token"`
}
