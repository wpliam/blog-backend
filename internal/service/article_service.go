package service

import "github.com/gin-gonic/gin"

type ArticleService interface {
	SearchArticleList(ctx *gin.Context) (interface{}, error)
	GetArticleArchive(ctx *gin.Context) (interface{}, error)
	GetHotArticle(ctx *gin.Context) (interface{}, error)
	ReadArticle(ctx *gin.Context) (interface{}, error)

	SearchKeywordFlow(ctx *gin.Context) (interface{}, error)
}
