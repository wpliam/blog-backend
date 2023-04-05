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

var defaultWeek = []*OneDay{
	{
		ID:         1,
		Day:        "一",
		Points:     10,
		Experience: 10,
		IsClock:    false,
	},
	{
		ID:         2,
		Day:        "二",
		Points:     20,
		Experience: 20,
		IsClock:    false,
	},
	{
		ID:         3,
		Day:        "三",
		Points:     30,
		Experience: 30,
		IsClock:    false,
	},
	{
		ID:         4,
		Day:        "四",
		Points:     40,
		Experience: 40,
		IsClock:    false,
	},
	{
		ID:         5,
		Day:        "五",
		Points:     50,
		Experience: 50,
		IsClock:    false,
	},
	{
		ID:         6,
		Day:        "六",
		Points:     60,
		Experience: 60,
		IsClock:    false,
	},
	{
		ID:         7,
		Day:        "七",
		Points:     70,
		Experience: 70,
		IsClock:    false,
	},
}

// CensusClockInfoReply 统计签到详情
type CensusClockInfoReply struct {
	MonthClockNum      int       `json:"monthClockNum"`      // 本月打卡次数
	ContinuousClockNum int       `json:"continuousClockNum"` // 当月连续打卡天数
	Days               []*OneDay `json:"days"`
}
