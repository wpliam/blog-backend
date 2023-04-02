package shared

import (
	"blog-backend/constant"
	"blog-backend/internal/service"
	"blog-backend/util"
	"fmt"
	"github.com/gin-gonic/gin"
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
		redisCli.ZIncrBy(ctx, constant.ArticleViewCountKey, fmt.Sprintf("%d", articleID), 1)
	}
	return nil
}

// GiveThumb 点赞
func (s *sharedImpl) GiveThumb(ctx *gin.Context) (interface{}, error) {
	var req *GiveThumbReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	redisCli := s.GetRedisProxy()
	userKey := util.GetUserLikeKey(util.GetUid(ctx))
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
	articleKey := fmt.Sprintf("%d", req.ID)
	incr := -1
	if isLike {
		incr = 1
	}
	likeCount, err := redisCli.ZIncrBy(ctx, constant.ArticleLikeCountKey, articleKey, float64(incr))
	if err != nil {
		return nil, err
	}
	rsp := &GiveThumbReply{
		LikeCount: int64(likeCount),
		IsLike:    isLike,
	}
	return rsp, nil
}

// GiveFollow 关注
func (s *sharedImpl) GiveFollow(ctx *gin.Context) (interface{}, error) {
	var req *GiveFollowReq
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
	rsp := GiveFollowRsp{
		IsFollow: isFollow,
	}
	return rsp, nil
}

// GiveCollect 收藏
func (s *sharedImpl) GiveCollect(ctx *gin.Context) (interface{}, error) {
	var req *GiveCollectReq
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
	collectCount, err := redisCli.HIncrBy(ctx, constant.ArticleCollectCountKey,
		fmt.Sprintf("%d", req.ArticleID), int64(incr))
	if err != nil {
		return nil, err
	}
	rsp := &GiveCollectReply{
		CollectCount: collectCount,
		IsCollect:    isCollect,
	}
	return rsp, nil
}
