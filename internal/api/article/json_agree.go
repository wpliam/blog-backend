package article

import "blog-backend/model"

type GetHotArticle struct {
}

type GetArticleReply struct {
	Articles []*model.Article `json:"articles"`
}

type SearchKeywordFlowReq struct {
	Keyword string `json:"keyword"`
}

type SearchKeywordFlowReply struct {
	Flows []*model.SearchFlow `json:"flows"`
}
