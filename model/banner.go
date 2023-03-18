package model

// Banner banner表
type Banner struct {
	Model
	BigTitle        string `json:"bigTitle"`
	SubTitle        string `json:"subTitle"`
	Weight          int    `json:"weight"`
	BackgroundImage string `json:"backgroundImage"`
}

func (*Banner) TableName() string {
	return BannerTableName
}
