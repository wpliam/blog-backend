package article

import (
	"blog-backend/constant"
	"blog-backend/model"
	"blog-backend/util"
	"blog-backend/util/thread"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wpliap/common-wrap/log"
)

type ReadArticle struct {
	ArticleID int64
}

type ReadArticleReply struct {
	Article struct {
		*model.ArticleContentSummary
		*model.ArticleContentInfo
		IsLike    bool `json:"isLike"`
		IsCollect bool `json:"isCollect"`
	} `json:"article"`
	Next      *model.Article                 `json:"next"`      // 下一篇文章
	Prev      *model.Article                 `json:"prev"`      // 上一篇文章
	Tags      []*model.Tag                   `json:"tags"`      // 文章标签
	Recommend []*model.ArticleContentSummary `json:"recommend"` // 文件推荐
	Comment   []*model.CommentContent        `json:"comment"`   // 文章评论
}

// ReadArticleImpl 读取文章
func (a *articleImpl) ReadArticleImpl(ctx *gin.Context, articleID int64) (*ReadArticleReply, error) {
	dbCli := a.GetGormProxy()
	esCli := a.GetElasticProxy()
	redisCli := a.GetRedisProxy()
	summary, err := esCli.GetArticleInfo(ctx, articleID)
	if err != nil {
		return nil, err
	}
	rsp := &ReadArticleReply{}
	articleContent := &model.ArticleContentInfo{}
	rsp.Article.ArticleContentSummary = summary
	handler := make([]func() error, 0)
	handler = append(handler, func() error {
		var err error
		articleContent, err = dbCli.GetArticleContentInfo(articleID)
		return err
	})
	handler = append(handler, func() error {
		var err error
		rsp.Next, err = dbCli.GetNextArticle(articleID)
		return err
	})
	handler = append(handler, func() error {
		var err error
		rsp.Prev, err = dbCli.GetPrevArticle(articleID)
		return err
	})
	if len(summary.TagIDs) > 0 {
		handler = append(handler, func() error {
			var err error
			rsp.Tags, err = dbCli.GetTagList(summary.TagIDs...)
			return err
		})
	}
	if len(summary.RecommendIDs) > 0 {
		handler = append(handler, func() error {
			var err error
			rsp.Recommend, err = esCli.GetArticleList(ctx, summary.RecommendIDs)
			return err
		})
	}
	handler = append(handler, func() error {
		var err error
		rsp.Comment, err = a.GetArticleComment(articleID)
		return err
	})
	if err = thread.GoAndWait(handler...); err != nil {
		log.Errorf("ReadArticle err:%v articleID:%s", err, articleID)
		return nil, err
	}
	articleStrID := fmt.Sprintf("%d", articleID)
	likeCount, err := redisCli.ZScore(ctx, constant.ArticleLikeCountKey, articleStrID)
	if err == nil {
		articleContent.LikeCount = int64(likeCount)
	}
	viewCount, err := redisCli.ZScore(ctx, constant.ArticleViewCountKey, articleStrID)
	if err == nil {
		articleContent.ViewCount = int64(viewCount)
	}
	collectCount, err := redisCli.HGet(ctx, constant.ArticleCollectCountKey, articleStrID)
	if err == nil {
		articleContent.CollectCount = collectCount
	}
	uid := util.GetUid(ctx)
	if uid > 0 {
		rsp.Article.IsLike = redisCli.SIsMember(ctx, util.GetUserLikeKey(uid), articleID)
		rsp.Article.IsCollect = redisCli.SIsMember(ctx, util.GetUserCollectKey(uid), articleID)
	}
	rsp.Article.ArticleContentInfo = articleContent
	return rsp, nil
}

// GetArticleComment 获取文章的评论
func (a *articleImpl) GetArticleComment(articleID int64) ([]*model.CommentContent, error) {
	dbCli := a.GetGormProxy()
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
	var list []*model.CommentContent
	for _, root := range commentList {
		subComment, err := dbCli.GetCommentInfo(articleID, root.ID)
		if err != nil {
			return nil, err
		}
		var subCommentList []*model.CommentContent
		for _, sub := range subComment {
			subCommentList = append(subCommentList, &model.CommentContent{
				ID:         sub.ID,
				CreateTime: sub.CreateTime,
				Content:    sub.Content,
				User:       userInfo[sub.UserID],
				ReplyUser:  userInfo[sub.ReplyUserID],
			})
		}
		info := &model.CommentContent{
			ID:         root.ID,
			CreateTime: root.CreateTime,
			Content:    root.Content,
			User:       userInfo[root.UserID],
			SubComment: subCommentList,
		}
		list = append(list, info)
	}
	return list, nil
}
