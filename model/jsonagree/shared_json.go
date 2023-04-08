package jsonagree

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

type GiveFollowReply struct {
	IsFollow bool `json:"isFollow"`
}

type CensusClockInfoReq struct {
	Uid int64 `json:"uid" binding:"min=1"`
}

type OneDay struct {
	ID         int    `json:"id"`
	Day        string `json:"day"`
	Points     int    `json:"points"`
	Experience int    `json:"experience"`
	IsClock    bool   `json:"isClock"`
}

// CensusClockInfoReply 统计签到详情
type CensusClockInfoReply struct {
	MonthClockNum      int `json:"monthClockNum"`      // 本月打卡次数
	ContinuousClockNum int `json:"continuousClockNum"` // 当月连续打卡天数
}
