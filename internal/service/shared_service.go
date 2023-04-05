package service

import "github.com/gin-gonic/gin"

type SharedService interface {
	// AddViewCount 添加访问量
	AddViewCount(ctx *gin.Context) error
	// GiveThumb 点赞
	GiveThumb(ctx *gin.Context) (interface{}, error)
	// GiveFollow 关注
	GiveFollow(ctx *gin.Context) (interface{}, error)
	// GiveCollect 收藏
	GiveCollect(ctx *gin.Context) (interface{}, error)
	// PunchClock 打卡
	PunchClock(ctx *gin.Context) error
	// CensusClockInfo 统计签到信息
	CensusClockInfo(ctx *gin.Context) (interface{}, error)
}
