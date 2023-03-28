package model

import "time"

const (
	ArticleTableName          = "t_article"           // 文章表
	BannerTableName           = "t_banner"            // banner表
	TagTableName              = "t_tag"               // 标签表
	CategoryTableName         = "t_category"          // 分类表
	CommentTableName          = "t_comment"           // 评论表
	UserTableName             = "t_user"              // 用户表
	ArticleTagTableName       = "t_article_tag"       // 文章标签表
	ArticleRecommendTableName = "t_article_recommend" // 文章推荐表
)

// Model 基本字段
type Model struct {
	ID         int64     `json:"id" gorm:"primaryKey"`
	CreateTime time.Time `json:"createTime" gorm:"<-:false"`
	UpdateTime time.Time `json:"updateTime" gorm:"<-:false"`
}

// Page 分页
type Page struct {
	Offset int   `json:"offset"`
	Limit  int   `json:"limit"`
	Total  int64 `json:"total"`
}

// NewPage ...
func NewPage(offset, limit int) *Page {
	return &Page{
		Offset: offset,
		Limit:  limit,
	}
}

func (p *Page) SetTotal(total int64) {
	p.Total = total
}

//Name string `gorm:"<-:create"` // 允许读和创建
//Name string `gorm:"<-:update"` // 允许读和更新
//Name string `gorm:"<-"`        // 允许读和写（创建和更新）
//Name string `gorm:"<-:false"`  // 允许读，禁止写
//Name string `gorm:"->"`        // 只读（除非有自定义配置，否则禁止写）
//Name string `gorm:"->;<-:create"` // 允许读和写
//Name string `gorm:"->:false;<-:create"` // 仅创建（禁止从 db 读）
//Name string `gorm:"-"`  // 通过 struct 读写会忽略该字段
//Name string `gorm:"-:all"`        // 通过 struct 读写、迁移会忽略该字段
//Name string `gorm:"-:migration"`  // 通过 struct 迁移会忽略该字段
