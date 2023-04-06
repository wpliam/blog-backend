package jsonagree

import "blog-backend/model"

type GetCategoryCardReq struct {
}

// GetCategoryCardReply 分类卡片
type GetCategoryCardReply struct {
	CategoryCard []*model.CategoryCard `json:"categoryCard"`
}

type GetCategoryListReq struct {
}

type GetCategoryListReply struct {
	CategoryList []*model.Category `json:"categoryList"`
}
