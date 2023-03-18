package model

// Tag 标签表
type Tag struct {
	Model
	TagName   string `json:"tagName"`
	AliasName string `json:"aliasName"`
	Desc      string `json:"desc"`
	Status    int    `json:"status"`
}

func (*Tag) TableName() string {
	return TagTableName
}
