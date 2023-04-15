package constant

import "time"

const (
	LoginJwtValidTime   = 24 * time.Hour // 登录有效时间
	LoginRedisValidTime = 2 * LoginJwtValidTime
)

const (
	EsArticleIndex = "blog_article_index"
)

const (
	ArticleLikeCountKey    = "article_like_count"
	ArticleViewCountKey    = "article_view_count"
	ArticleCollectCountKey = "article_collect_count"
	CommentLikeCountKey    = "comment_like_count"
)

const (
	MysqlDefaultTimeLayout string = "2006-01-02T15:04:05+08:00"
	TimeStampLayout        string = "20060102150405"
	TimeStampDayLayout     string = "20060102"
	TimeStampMonthLayout   string = "200601"
	TimeLayout             string = "2006-01-02 15:04:05"
	MonthSubTableSuffix    string = "2006-01"
)

const (
	SearchNewCreateArticle = 1 // SearchNewCreateArticle 按照创建时间排序
	SearchNewUpdateArticle = 2 // SearchNewUpdateArticle 按照更新时间排序
)

const (
	StateArticleEdit   = 0 // 文章编辑状态(草稿)
	StateArticlePush   = 1 // 文章待审核状态
	StateArticlePass   = 2 // 文章审核通过状态
	StateArticleReject = 3 // 文章审核拒绝状态
)
