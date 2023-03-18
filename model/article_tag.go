package model

// ArticleTag 文章标签表
type ArticleTag struct {
	Model
	ArticleID int64 `json:"articleID"`
	TagID     int64 `json:"tagID"`
}

func (*ArticleTag) TableName() string {
	return ArticleTagTableName
}
