package model

// SearchFlow 搜索流水
type SearchFlow struct {
	Model
	Keyword string `json:"keyword"`
	Flag    uint32 `json:"flag"`
	Version int    `json:"version"`
}

func (*SearchFlow) TableName() string {
	return "t_search_flow"
}
