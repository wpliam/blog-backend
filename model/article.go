package model

import (
	"blog-backend/constant"
	"blog-backend/util"
	"gorm.io/gorm"
)

// Article 文章表
type Article struct {
	ArticleBaseInfo
	CategoryID int64     `json:"cid" gorm:"column:category_id"` // 文章分类id
	UserID     int64     `json:"uid" gorm:"column:user_id"`     // 文章作者id
	Content    string    `json:"content" gorm:"column:content"` // 文章内容
	Category   *Category `json:"category"`                      // 分类信息
	User       *User     `json:"user"`                          // 用户信息
}

func (*Article) TableName() string {
	return ArticleTableName
}

func (a *Article) AfterFind(db *gorm.DB) error {
	a.CreateTime = util.ParseDateTime(constant.TimeLayout, a.CreateTime)
	a.UpdateTime = util.ParseDateTime(constant.TimeLayout, a.UpdateTime)
	return nil
}

// ArticleBaseInfo 文章基本信息
type ArticleBaseInfo struct {
	Model
	Title        string `json:"title" gorm:"column:title"`                // 文章标题
	Abstract     string `json:"abstract" gorm:"column:abstract"`          // 文章摘要
	Status       uint32 `json:"status" gorm:"column:status"`              // 文章状态 0:待审核 1:审核通过 2:审核拒绝
	Cover        string `json:"cover" gorm:"column:cover"`                // 文章封面
	ArticleType  int    `json:"articleType" gorm:"column:article_type"`   // 文章类型 0:原创 1:转载
	LikeCount    int64  `json:"likeCount" gorm:"column:like_count"`       // 文章点赞数
	ViewCount    int64  `json:"viewCount" gorm:"column:view_count"`       // 文章阅读数量
	CollectCount int64  `json:"collectCount" gorm:"column:collect_count"` // 文章收藏数
}

// ArticleContentSummary 写入es的结构
type ArticleContentSummary struct {
	ArticleBaseInfo
	Category struct {
		Cid          int64  `json:"cid"`
		CategoryName string `json:"categoryName"`
	} `json:"category"`
	User struct {
		Uid      int64  `json:"uid"`
		Nickname string `json:"nickname"`
	} `json:"user"`
}

func (a *Article) ArticleContentSummary() *ArticleContentSummary {
	if a.Category == nil {
		a.Category = &Category{}
	}
	if a.User == nil {
		a.User = &User{}
	}
	summary := &ArticleContentSummary{}
	summary.ArticleBaseInfo = a.ArticleBaseInfo
	summary.Category.Cid = a.CategoryID
	summary.Category.CategoryName = a.Category.CategoryName
	summary.User.Uid = a.UserID
	summary.User.Nickname = a.User.Nickname
	return summary
}
