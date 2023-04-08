package jsonagree

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

type ReadArticle struct {
	ArticleID int64
}

type ReadArticleReply struct {
	Article struct {
		*model.ArticleContentSummary
		*model.ArticleContentInfo
		IsLike       bool  `json:"isLike"`
		IsCollect    bool  `json:"isCollect"`
		CommentCount int64 `json:"commentCount"`
	} `json:"article"`
	Next      *model.Article                 `json:"next"`      // 下一篇文章
	Prev      *model.Article                 `json:"prev"`      // 上一篇文章
	Tags      []*model.Tag                   `json:"tags"`      // 文章标签
	Recommend []*model.ArticleContentSummary `json:"recommend"` // 文件推荐
}

type SearchArticleListReq struct {
	Keyword string      `json:"keyword"`
	Title   string      `json:"title"`
	Cid     int64       `json:"cid"`
	TagID   int64       `json:"tagID"`
	Order   int         `json:"order"`
	Uid     int64       `json:"uid"`
	Page    *model.Page `json:"page"`

	SearchType int `json:"searchType"` // 0:搜索文章 1:搜索随机文章
}

// SearchArticleListReply 搜索文章响应体
type SearchArticleListReply struct {
	Page     *model.Page                    `json:"page"`
	Articles []*model.ArticleContentSummary `json:"articles"`
}

type GetArticleArchiveReq struct {
}

type GetArticleArchiveReply struct {
	Article       map[string][]*model.ArticleContentSummary `json:"article"`
	ArticleCount  int64                                     `json:"articleCount"`
	Tags          []*model.Tag                              `json:"tags"`
	TagCount      int64                                     `json:"tagCount"`
	Category      []*model.Category                         `json:"category"`
	CategoryCount int64                                     `json:"categoryCount"`
}

type WriteArticleReq struct {
	Title       string           `json:"title" binding:"min=1"`
	Abstract    string           `json:"abstract"`
	Content     string           `json:"content" binding:"min=1"`
	Cover       string           `json:"cover" binding:"min=1"`
	Tags        []string         `json:"tags"`
	Recommends  []*model.Article `json:"recommends"`
	Cid         int64            `json:"cid" binding:"min=1"`
	ArticleType uint32           `json:"articleType"`
}

type ArticleReviewReq struct {
	ArticleID int64 `json:"articleID"`
	Pass      int64 `json:"pass"` // 0:不通过 1:通过
}

type ArticleReviewRsp struct {
}
