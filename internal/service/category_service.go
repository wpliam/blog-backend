package service

import "github.com/gin-gonic/gin"

type CategoryService interface {
	GetCategoryCard(ctx *gin.Context) (interface{}, error)
	GetCategoryList(ctx *gin.Context) (interface{}, error)
}
