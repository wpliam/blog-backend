package user

import (
	"blog-backend/util"
	"blog-backend/util/thread"
	"github.com/gin-gonic/gin"
	"github.com/wpliap/common-wrap/log"
)

type StaticUserInfoReply struct {
	ArticleCount int64 `json:"articleCount"` // 用户发表了多少篇文章
	CommentCount int64 `json:"commentCount"` // 用户发表了多少评论
	HotCount     int64 `json:"hotCount"`     // 用户人气值

	FansCount    int64 `json:"fansCount"`    // 用户有多少粉丝
	CollectCount int64 `json:"collectCount"` // 用户收藏文章数
	LikeCount    int64 `json:"likeCount"`    // 获得多少点赞
}

// StaticUserInfoImpl 统计用户信息
func (u *userImpl) StaticUserInfoImpl(ctx *gin.Context, uid int64) (*StaticUserInfoReply, error) {
	rsp := &StaticUserInfoReply{
		FansCount: 10,
	}
	dbCli := u.GetGormProxy()
	redisCli := u.GetRedisProxy()
	handler := make([]func() error, 0)
	handler = append(handler, func() error {
		var err error
		rsp.CommentCount, err = dbCli.GetUserCommentCount(uid)
		return err
	})
	handler = append(handler, func() error {
		var err error
		rsp.ArticleCount, err = dbCli.GetUserArticleCount(uid)
		return err
	})
	handler = append(handler, func() error {
		var err error
		rsp.HotCount, err = dbCli.GetUserViewCount(uid)
		return err
	})
	handler = append(handler, func() error {
		var err error
		rsp.LikeCount, err = redisCli.SCard(ctx, util.GetUserLikeKey(uid))
		return err
	})
	handler = append(handler, func() error {
		var err error
		rsp.CollectCount, err = redisCli.SCard(ctx, util.GetUserCollectKey(uid))
		return err
	})
	handler = append(handler, func() error {
		var err error
		rsp.FansCount, err = redisCli.SCard(ctx, util.GetUserFansKey(uid))
		return err
	})
	if err := thread.GoAndWait(handler...); err != nil {
		log.Errorf("StaticUserInfo GoAndWait err:%v uid:%d", err, uid)
		return nil, err
	}
	return rsp, nil
}
