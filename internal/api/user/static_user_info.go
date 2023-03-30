package user

import (
	"blog-backend/util/thread"
	"github.com/gin-gonic/gin"
	"github.com/wpliap/common-wrap/log"
)

type NeedUserInfo struct {
	ID           int64  `json:"id"`           // 用户id
	Avatar       string `json:"avatar"`       // 用户图像
	Nickname     string `json:"nickname"`     // 用户昵称
	Desc         string `json:"desc"`         // 用户描述
	ArticleCount int64  `json:"articleCount"` // 用户发表了多少篇文章
	CommentCount int64  `json:"commentCount"` // 用户发表了多少评论
	HotCount     int64  `json:"hotCount"`     // 用户人气值
	FansCount    int64  `json:"fansCount"`    // 用户有多少粉丝
}

type StaticUserInfoReply struct {
	User *NeedUserInfo `json:"user"`
}

// StaticUserInfoImpl 统计用户信息
func (u *userImpl) StaticUserInfoImpl(ctx *gin.Context, uid int64) (*NeedUserInfo, error) {
	userInfo := &NeedUserInfo{
		FansCount: 10,
	}
	dbCli := u.GetGormProxy()
	handler := make([]func() error, 0)
	handler = append(handler, func() error {
		info, err := dbCli.GetUserInfo(uid)
		if err != nil {
			return err
		}
		userInfo.ID = info.ID
		userInfo.Nickname = info.Nickname
		userInfo.Avatar = info.Avatar
		userInfo.Desc = info.Desc
		return nil
	})
	handler = append(handler, func() error {
		var err error
		userInfo.CommentCount, err = dbCli.GetUserCommentCount(uid)
		return err
	})
	handler = append(handler, func() error {
		var err error
		userInfo.ArticleCount, err = dbCli.GetUserArticleCount(uid)
		return err
	})
	handler = append(handler, func() error {
		var err error
		userInfo.HotCount, err = dbCli.GetUserViewCount(uid)
		return err
	})
	if err := thread.GoAndWait(handler...); err != nil {
		log.Errorf("StaticUserInfo GoAndWait err:%v uid:%d", err, uid)
		return nil, err
	}
	return userInfo, nil
}
