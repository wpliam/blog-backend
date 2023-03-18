package model

// ArticleRecommend 文章推荐表
type ArticleRecommend struct {
	Model
	ArticleID          int64 `json:"articleID"`
	RecommendArticleID int64 `json:"recommendArticleID"`
}

func (*ArticleRecommend) TableName() string {
	return ArticleRecommendTableName
}
