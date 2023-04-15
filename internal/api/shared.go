package api

import (
	"blog-backend/constant"
	"blog-backend/internal/service"
	"blog-backend/model/jsonagree"
	"blog-backend/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wpliap/common-wrap/log"
	"time"
)

func NewSharedService(proxyService service.ProxyService) service.SharedService {
	return &sharedImpl{
		proxyService,
	}
}

type sharedImpl struct {
	service.ProxyService
}

// AddViewCount 添加文章阅读数量 ip+articleID 一分钟内只增加一次浏览量
func (s *sharedImpl) AddViewCount(ctx *gin.Context) error {
	ip := util.GetClientIP(ctx)
	articleID := util.GetArticleID(ctx)
	if articleID == 0 {
		return nil
	}
	redisCli := s.GetRedisProxy()
	success, err := redisCli.SetNX(ctx, util.GetArticleIPKey(ip, articleID), articleID, time.Minute)
	if err == nil && success {
		incr := float64(1)
		if exists := redisCli.ZExists(ctx, constant.ArticleViewCountKey, fmt.Sprintf("%d", articleID)); !exists {
			articleInfo, err := s.GetGormProxy().GetArticleContentInfo(articleID)
			if err == nil {
				incr += float64(articleInfo.ViewCount)
			}
		}
		if _, err = redisCli.ZIncrBy(ctx, constant.ArticleViewCountKey,
			fmt.Sprintf("%d", articleID), incr); err != nil {
			log.Errorf("AddViewCount ZIncrBy err:%v articleID:%d", err, articleID)
		}
	}
	return nil
}

// GiveThumb 点赞
func (s *sharedImpl) GiveThumb(ctx *gin.Context) (interface{}, error) {
	var req *jsonagree.GiveThumbReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	redisCli := s.GetRedisProxy()
	userKey := util.GetUserArticleLikeKey(util.GetUid(ctx))
	likeCountKey := constant.ArticleLikeCountKey
	// 给评论点赞
	if req.LikeType == 1 {
		userKey = util.GetUserCommentLikeKey(util.GetUid(ctx))
		likeCountKey = constant.CommentLikeCountKey
	}
	// 文章是否在用户的点赞列表
	isMember := redisCli.SIsMember(ctx, userKey, req.ID)
	var err error
	var isLike bool
	if isMember {
		err = redisCli.SRem(ctx, userKey, req.ID)
		isLike = false
	} else {
		err = redisCli.SAdd(ctx, userKey, req.ID)
		isLike = true
	}
	if err != nil {
		return nil, err
	}
	idKey := fmt.Sprintf("%d", req.ID)
	incr := -1
	if isLike {
		incr = 1
	}
	if exists := redisCli.ZExists(ctx, likeCountKey, idKey); !exists {
		dbCli := s.GetGormProxy()
		if req.LikeType == 0 {
			articleInfo, err := dbCli.GetArticleContentInfo(req.ID)
			if err == nil {
				incr += int(articleInfo.LikeCount)
			}
		} else if req.LikeType == 1 {
			commentInfo, err := dbCli.GetCommentByID(req.ID)
			if err == nil {
				incr += int(commentInfo.Likes)
			}
		}
	}

	likeCount, err := redisCli.ZIncrBy(ctx, likeCountKey, idKey, float64(incr))
	if err != nil {
		return nil, err
	}
	rsp := &jsonagree.GiveThumbReply{
		LikeCount: int64(likeCount),
		IsLike:    isLike,
	}
	return rsp, nil
}

// GiveFollow 关注
func (s *sharedImpl) GiveFollow(ctx *gin.Context) (interface{}, error) {
	var req *jsonagree.GiveFollowReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	var isFollow bool
	// 校验作者是否在该用户的关注列表
	uid := util.GetUid(ctx)
	redisCli := s.GetRedisProxy()
	isMember := redisCli.SIsMember(ctx, util.GetUserFollowKey(uid), req.AuthorID)
	if isMember {
		_ = redisCli.SRem(ctx, util.GetUserFollowKey(uid), req.AuthorID)
		_ = redisCli.SRem(ctx, util.GetUserFansKey(req.AuthorID), uid)
		isFollow = false
	} else {
		_ = redisCli.SAdd(ctx, util.GetUserFollowKey(uid), req.AuthorID)
		_ = redisCli.SAdd(ctx, util.GetUserFansKey(req.AuthorID), uid)
		isFollow = true
	}
	rsp := &jsonagree.GiveFollowReply{
		IsFollow: isFollow,
	}
	return rsp, nil
}

// GiveCollect 收藏
func (s *sharedImpl) GiveCollect(ctx *gin.Context) (interface{}, error) {
	var req *jsonagree.GiveCollectReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	// 文章是否在用户的收藏列表
	redisCli := s.GetRedisProxy()
	userKey := util.GetUserCollectKey(util.GetUid(ctx))
	var err error
	var isCollect bool
	isMember := redisCli.SIsMember(ctx, userKey, req.ArticleID)
	if isMember {
		err = redisCli.SRem(ctx, userKey, req.ArticleID)
		isCollect = false
	} else {
		err = redisCli.SAdd(ctx, userKey, req.ArticleID)
		isCollect = true
	}
	if err != nil {
		return nil, err
	}
	incr := -1
	if isCollect {
		incr = 1
	}
	// 如果redis hash中不存在数据,将db数据一并写入redis
	articleID := fmt.Sprintf("%d", req.ArticleID)
	if exists := redisCli.HExists(ctx, constant.ArticleCollectCountKey, articleID); !exists {
		dbCli := s.GetGormProxy()
		articleInfo, err := dbCli.GetArticleContentInfo(req.ArticleID)
		if err == nil {
			incr += int(articleInfo.CollectCount)
		}
	}

	collectCount, err := redisCli.HIncrBy(ctx, constant.ArticleCollectCountKey, articleID, int64(incr))
	if err != nil {
		return nil, err
	}
	rsp := &jsonagree.GiveCollectReply{
		CollectCount: collectCount,
		IsCollect:    isCollect,
	}
	return rsp, nil
}

// PunchClock 打卡
func (s *sharedImpl) PunchClock(ctx *gin.Context) error {
	uid := util.GetUid(ctx)
	if uid <= 0 {
		return fmt.Errorf("uid not exist")
	}
	if err := s.GetRedisProxy().SetBit(ctx, util.GetClockKey(uid), int64(time.Now().Local().Day()-1), 1); err != nil {
		return err
	}
	return nil
}

// CensusClockInfo 统计用户签到信息
func (s *sharedImpl) CensusClockInfo(ctx *gin.Context) (interface{}, error) {
	var req *jsonagree.CensusClockInfoReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	redisCli := s.GetRedisProxy()
	monthClockNum, err := redisCli.BitCount(ctx, util.GetClockKey(req.Uid))
	if err != nil {
		return nil, err
	}
	result, err := redisCli.BitGetDay(ctx, util.GetClockKey(req.Uid))
	if err != nil {
		return nil, err
	}
	var continuousClockNum int
	if len(result) > 0 {
		num := result[0]
		for {
			// 让这个数字与1做与运算,得到数字的最后一个bit位, 判断这个bit是否为0, 如果是0说明未签到,结束
			if (num & 1) == 0 {
				break
			}
			continuousClockNum++
			// 把数字右移一位,继续下一个bit位
			num = num >> 1
		}
	}
	rsp := &jsonagree.CensusClockInfoReply{
		MonthClockNum:      int(monthClockNum),
		ContinuousClockNum: continuousClockNum,
	}
	return rsp, nil
}
