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
)

const (
	MysqlDefaultTimeLayout string = "2006-01-02T15:04:05+08:00"
	TimeStampLayout        string = "20060102150405"      // TimeStampLayout 查询时间的精度，目前是秒
	NotifyLayout           string = "2006年01月02日15时04分"   // NotifyLayout 通知时间格式，目前是分钟
	TimeStampDayLayout     string = "20060102"            // TimeStampDayLayout 格式化时间，目前是天
	DayLayout              string = "2006年01月02日"         // DayLayout 时间格式话为天
	TimeLayout             string = "2006-01-02 15:04:05" // TimeLayout
	MonthSubTableSuffix    string = "2006-01"             // MonthSubTableSuffix 按月分后缀格式
	YearSubTableSuffix     string = "2006"                // YearSubTableSuffix 按年分后缀格式
	TimeStampForDownload   string = "2006_01_02_15_04_05" // TimeStampForDownload 下载使用的时间格式
)

const (
	SearchNewCreateArticle = 1 // SearchNewCreateArticle 按照创建时间排序
	SearchNewUpdateArticle = 2 // SearchNewUpdateArticle 按照更新时间排序
)
