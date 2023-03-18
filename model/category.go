package model

// Category 分类表
type Category struct {
	Model
	CategoryName string `json:"categoryName"`
	AliasName    string `json:"aliasName"`
	Desc         string `json:"desc"`
	ParentID     int64  `json:"parentID"`
	Status       int    `json:"status"`
}

func (*Category) TableName() string {
	return CategoryTableName
}
