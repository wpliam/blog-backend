package jsonagree

import "blog-backend/model"

type GetCommentReq struct {
	ArticleID int64 `json:"articleID" binding:"min=1"`
}

type GetCommentReply struct {
	Comment []*model.CommentContent `json:"comment"`
}

type AddCommentReq struct {
	ParentID       int64  `json:"parentID"`
	Content        string `json:"content" binding:"required"`
	ArticleID      int64  `json:"articleID" binding:"min=1"`
	ReplyCommentID int64  `json:"replyCommentID"`
	ReplyUserID    int64  `json:"replyUserID"`
}
