package comment

import (
	"blog-backend/model"
	"blog-backend/repo/mdb"
	"github.com/gin-gonic/gin"
)

type GetCommentReq struct {
	ArticleID int64 `json:"articleID" binding:"required"`
}

type GetCommentRsp struct {
	Comments []*model.CommentInfo `json:"comments"`
}

// GetComment 获取评论信息
func GetComment(ctx *gin.Context) (interface{}, error) {
	var req *GetCommentReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	dbCli := mdb.GetGormClient()
	userIDs, err := dbCli.GetCommentUserIDs(req.ArticleID)
	if err != nil {
		return nil, err
	}

	userInfo, err := dbCli.BatchGetUserInfo(userIDs)
	if err != nil {
		return nil, err
	}
	commentList, err := dbCli.GetCommentInfo(req.ArticleID, 0)
	if err != nil {
		return nil, err
	}
	var list []*model.CommentInfo
	for _, root := range commentList {
		subComment, err := dbCli.GetCommentInfo(req.ArticleID, root.ID)
		if err != nil {
			return nil, err
		}
		var subCommentList []*model.CommentInfo
		for _, sub := range subComment {
			subCommentList = append(subCommentList, &model.CommentInfo{
				ID:            sub.ID,
				CreateTime:    sub.CreateTime,
				Content:       sub.Content,
				UserInfo:      userInfo[sub.UserID],
				ReplyUserInfo: userInfo[sub.ReplyUserID],
			})
		}
		info := &model.CommentInfo{
			ID:             root.ID,
			CreateTime:     root.CreateTime,
			Content:        root.Content,
			UserInfo:       userInfo[root.UserID],
			SubCommentInfo: subCommentList,
		}
		list = append(list, info)
	}
	rsp := &GetCommentRsp{
		Comments: list,
	}
	return rsp, nil
}
