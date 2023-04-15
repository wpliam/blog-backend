package service

import "github.com/gin-gonic/gin"

type AdminService interface {
	GetReadyReviewArticle(ctx *gin.Context) (interface{}, error)
	ArticleReview(ctx *gin.Context) error
}
