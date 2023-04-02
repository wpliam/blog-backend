package shared

type GiveThumbReq struct {
	ID       int64  `json:"id" binding:"min=1"`
	LikeType uint32 `json:"likeType"` // 0:文章 1:评论
}
type GiveThumbReply struct {
	LikeCount int64 `json:"likeCount"`
	IsLike    bool  `json:"isLike"`
}

type GiveCollectReq struct {
	ArticleID int64 `json:"articleID" binding:"min=1"`
}

type GiveCollectReply struct {
	CollectCount int64 `json:"collectCount"`
	IsCollect    bool  `json:"isCollect"`
}

type GiveFollowReq struct {
	AuthorID int64 `json:"authorID" binding:"min=1"` // 要关注的作者id
}

type GiveFollowRsp struct {
	IsFollow bool `json:"isFollow"`
}
