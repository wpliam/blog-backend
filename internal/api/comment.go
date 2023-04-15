package api

import (
	"blog-backend/constant"
	"blog-backend/internal/service"
	"blog-backend/model"
	"blog-backend/model/jsonagree"
	"blog-backend/util"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewCommentService(proxyService service.ProxyService) service.CommentService {
	return &commentImpl{
		proxyService,
	}
}

type commentImpl struct {
	service.ProxyService
}

func (c *commentImpl) GetComment(ctx *gin.Context) (interface{}, error) {
	var req *jsonagree.GetCommentReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	articleID := req.ArticleID
	dbCli := c.GetGormProxy()
	userIDs, err := dbCli.GetCommentUserIDs(articleID)
	if err != nil {
		return nil, err
	}
	userInfo, err := dbCli.BatchGetUserInfo(userIDs)
	if err != nil {
		return nil, err
	}
	commentList, err := dbCli.GetCommentInfo(articleID, 0)
	if err != nil {
		return nil, err
	}
	var comments []*model.CommentContent
	for _, root := range commentList {
		subComment, err := dbCli.GetCommentInfo(articleID, root.ID)
		if err != nil {
			return nil, err
		}
		var subCommentList []*model.CommentContent
		for _, sub := range subComment {
			subCommentList = append(subCommentList, &model.CommentContent{
				ID:         sub.ID,
				LikeCount:  sub.LikeCount,
				CreateTime: sub.CreateTime,
				Content:    sub.Content,
				User:       userInfo[sub.UserID],
				ReplyUser:  userInfo[sub.ReplyUserID],
			})
		}
		info := &model.CommentContent{
			ID:         root.ID,
			LikeCount:  root.LikeCount,
			CreateTime: root.CreateTime,
			Content:    root.Content,
			User:       userInfo[root.UserID],
			SubComment: subCommentList,
		}
		comments = append(comments, info)
	}
	c.fillCommentLike(ctx, comments)
	rsp := &jsonagree.GetCommentReply{
		Comment: comments,
	}
	return rsp, nil
}

func (c *commentImpl) fillCommentLike(ctx *gin.Context, comments []*model.CommentContent) {
	uid := util.GetUid(ctx)
	redisCli := c.GetRedisProxy()
	fillFunc := func(ctx context.Context, comment *model.CommentContent) {
		idStr := fmt.Sprintf("%d", comment.ID)
		likes, err := redisCli.HGet(ctx, constant.CommentLikeCountKey, idStr)
		if err == nil {
			comment.LikeCount = likes
		}
		if uid > 0 {
			comment.IsLike = redisCli.SIsMember(ctx, util.GetUserCommentLikeKey(uid), idStr)
		}
	}
	for _, root := range comments {
		fillFunc(ctx, root)
		for _, sub := range root.SubComment {
			fillFunc(ctx, sub)
		}
	}
}

// AddComment 添加评论
func (c *commentImpl) AddComment(ctx *gin.Context) error {
	var req *jsonagree.AddCommentReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return err
	}
	comment := &model.Comment{
		ParentID:       req.ParentID,
		UserID:         util.GetUid(ctx),
		Content:        req.Content,
		ArticleID:      req.ArticleID,
		ReplyCommentID: req.ReplyCommentID,
		ReplyUserID:    req.ReplyUserID,
	}
	if err := c.GetGormProxy().SetCommentInfo(comment); err != nil {
		return err
	}
	return nil
}
