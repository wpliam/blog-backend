package model

import (
	"gorm.io/gorm"
	"time"
)

// Comment 评论表
type Comment struct {
	Model
	ParentID       int64  `json:"parentID"`       // 父id
	UserID         int64  `json:"userID"`         // 评论用户id
	Content        string `json:"content"`        // 评论内容
	ArticleID      int64  `json:"articleID"`      // 文章id
	ReplyCommentID int64  `json:"replyCommentID"` // 回复的评论id
	ReplyUserID    int64  `json:"replyUserID"`    // 回复的用户ID
	Likes          int64  `json:"likes"`          // 评论获赞量
	Status         int    `json:"status"`         // 评论状态
}

func (*Comment) TableName() string {
	return CommentTableName
}

func (c *Comment) AfterFind(db *gorm.DB) error {
	return nil
}

type CommentContent struct {
	ID         int64             `json:"id"`
	CreateTime time.Time         `json:"createTime"`
	Content    string            `json:"content"`
	User       *User             `json:"user"`
	ReplyUser  *User             `json:"replyUser"`
	SubComment []*CommentContent `json:"subComment"`
}
