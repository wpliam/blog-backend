package jsonagree

import "blog-backend/model"

type GetReadyReviewArticleReq struct {
	Page *model.Page `yaml:"page"`
}

type GetReadyReviewArticleReply struct {
	Articles []*model.Article `yaml:"articles"`
	Page     *model.Page      `yaml:"page"`
}
