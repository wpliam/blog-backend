package model

// Category 分类表
type Category struct {
	Model
	CategoryName string `json:"categoryName"`
	Cover        string `json:"cover"`
	AliasName    string `json:"aliasName"`
	Desc         string `json:"desc"`
	ParentID     int64  `json:"parentID"`
	Status       int    `json:"status"`
}

func (*Category) TableName() string {
	return CategoryTableName
}

// CategoryCard 分类卡片信息
type CategoryCard struct {
	Cid          int64                    `json:"cid" gorm:"column:category_id"`
	CategoryName string                   `json:"categoryName" gorm:"-:all"`
	Cover        string                   `json:"cover" gorm:"-:all"`
	ViewCount    int64                    `json:"viewCount"`
	Total        int64                    `json:"total" gorm:"-:all"`
	Articles     []*ArticleContentSummary `json:"articles" gorm:"-:all"`
}
