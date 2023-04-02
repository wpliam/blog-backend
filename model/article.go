package model

// Article 文章表
type Article struct {
	ArticleBaseInfo
	Content      string    `json:"content" gorm:"column:content"`            // 文章内容
	LikeCount    int64     `json:"likeCount" gorm:"column:like_count"`       // 文章点赞数
	ViewCount    int64     `json:"viewCount" gorm:"column:view_count"`       // 文章阅读数量
	CollectCount int64     `json:"collectCount" gorm:"column:collect_count"` // 文章收藏数
	Category     *Category `json:"category"`                                 // 分类信息
	User         *User     `json:"user"`                                     // 用户信息
}

// TableName 文章表名
func (*Article) TableName() string {
	return ArticleTableName
}

// ArticleContentInfo 文章内容信息
type ArticleContentInfo struct {
	Content      string `json:"content" gorm:"column:content"`            // 文章内容
	LikeCount    int64  `json:"likeCount" gorm:"column:like_count"`       // 文章点赞数
	ViewCount    int64  `json:"viewCount" gorm:"column:view_count"`       // 文章阅读数量
	CollectCount int64  `json:"collectCount" gorm:"column:collect_count"` // 文章收藏数
}

// TableName 文章表名
func (*ArticleContentInfo) TableName() string {
	return ArticleTableName
}

// ArticleBaseInfo 文章基本信息
type ArticleBaseInfo struct {
	Model
	CategoryID  int64  `json:"cid" gorm:"column:category_id"`          // 文章分类id
	UserID      int64  `json:"uid" gorm:"column:user_id"`              // 文章作者id
	Title       string `json:"title" gorm:"column:title"`              // 文章标题
	Abstract    string `json:"abstract" gorm:"column:abstract"`        // 文章摘要
	Status      uint32 `json:"status" gorm:"column:status"`            // 文章状态 0:待审核 1:审核通过 2:审核拒绝
	Cover       string `json:"cover" gorm:"column:cover"`              // 文章封面
	ArticleType int    `json:"articleType" gorm:"column:article_type"` // 文章类型 0:原创 1:转载
}

// ArticleContentSummary 写入es的结构
type ArticleContentSummary struct {
	ArticleBaseInfo
	TagIDs       []int64 `json:"tagIDs"`       // 标签ID
	RecommendIDs []int64 `json:"recommendIDs"` // 文章推荐ID

	Nickname string `json:"nickname"` // 用户名
	Avatar   string `json:"avatar"`   // 用户图像

	CategoryName  string `json:"categoryName"`  // 分类名
	CategoryCover string `json:"categoryCover"` // 分类背景图
}

func (a *Article) ArticleContentSummary(tagIDs, recommendIDs []int64) *ArticleContentSummary {
	if a.Category == nil {
		a.Category = &Category{}
	}
	if a.User == nil {
		a.User = &User{}
	}
	summary := &ArticleContentSummary{
		ArticleBaseInfo: a.ArticleBaseInfo,
		TagIDs:          tagIDs,
		RecommendIDs:    recommendIDs,
		Nickname:        a.User.Nickname,
		Avatar:          a.User.Avatar,
		CategoryName:    a.Category.CategoryName,
		CategoryCover:   a.Category.Cover,
	}
	return summary
}

// SearchArticleParam 搜索文章的参数
type SearchArticleParam struct {
	Keyword string `json:"keyword"`
	Cid     int64  `json:"cid"`
	TagID   int64  `json:"tagID"`
	Order   int    `json:"order"`
	Uid     int64  `json:"uid"`
	Page    *Page  `json:"page"`
}
