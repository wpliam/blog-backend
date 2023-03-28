package article

import (
	"blog-backend/internal/common/proxy"
	"blog-backend/model"
	"blog-backend/util/thread"
	"github.com/gin-gonic/gin"
	"github.com/wpliap/common-wrap/log"
	"strconv"
)

type ReadArticle struct {
	ArticleID int64
}

type ReadArticleReply struct {
	Article   *model.Article                 `json:"article"`   // 文章信息
	Next      *model.Article                 `json:"next"`      // 下一篇文章
	Prev      *model.Article                 `json:"prev"`      // 上一篇文章
	Tags      []*model.Tag                   `json:"tags"`      // 文章标签
	Recommend []*model.ArticleContentSummary `json:"recommend"` // 文件推荐
	Comment   []*model.CommentContent        `json:"comment"`   // 文章评论
}

func (r *ReadArticle) Invoke(ctx *gin.Context, proxy proxy.Proxy) (interface{}, error) {
	articleID, err := strconv.ParseInt(ctx.Param("articleID"), 10, 64)
	if err != nil {
		return nil, err
	}
	r.ArticleID = articleID
	return r.ReadArticle(ctx, proxy)
}

// ReadArticle 读取文章
func (r *ReadArticle) ReadArticle(ctx *gin.Context, proxy proxy.Proxy) (*ReadArticleReply, error) {
	articleInfo, err := proxy.GetEsProxy().GetArticleInfo(ctx, r.ArticleID)
	if err != nil {
		return nil, err
	}
	rsp := &ReadArticleReply{}
	handler := make([]func() error, 0)
	handler = append(handler, func() error {
		var err error
		rsp.Article, err = proxy.GetGormProxy().GetArticleInfo(r.ArticleID)
		return err
	})
	handler = append(handler, func() error {
		var err error
		rsp.Next, err = proxy.GetGormProxy().GetNextArticle(r.ArticleID)
		return err
	})
	handler = append(handler, func() error {
		var err error
		rsp.Prev, err = proxy.GetGormProxy().GetPrevArticle(r.ArticleID)
		return err
	})
	handler = append(handler, func() error {
		var err error
		rsp.Tags, err = proxy.GetGormProxy().GetTagList(articleInfo.TagIDs...)
		return err
	})
	handler = append(handler, func() error {
		var err error
		rsp.Recommend, err = proxy.GetEsProxy().GetArticleList(ctx, articleInfo.RecommendIDs)
		return err
	})
	handler = append(handler, func() error {
		var err error
		rsp.Comment, err = GetArticleComment(r.ArticleID, proxy)
		return err
	})
	if err = thread.GoAndWait(handler...); err != nil {
		log.Errorf("ReadArticle err:%v articleID:%s", err, r.ArticleID)
		return nil, err
	}
	return rsp, nil
}

// GetArticleComment 获取文章的评论
func GetArticleComment(articleID int64, proxy proxy.Proxy) ([]*model.CommentContent, error) {
	userIDs, err := proxy.GetGormProxy().GetCommentUserIDs(articleID)
	if err != nil {
		return nil, err
	}
	userInfo, err := proxy.GetGormProxy().BatchGetUserInfo(userIDs)
	if err != nil {
		return nil, err
	}
	commentList, err := proxy.GetGormProxy().GetCommentInfo(articleID, 0)
	if err != nil {
		return nil, err
	}
	var list []*model.CommentContent
	for _, root := range commentList {
		subComment, err := proxy.GetGormProxy().GetCommentInfo(articleID, root.ID)
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
