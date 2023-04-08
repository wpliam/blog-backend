package service

import "github.com/gin-gonic/gin"

type CommentService interface {
	AddComment(ctx *gin.Context) error
	GetComment(ctx *gin.Context) (interface{}, error)
}
