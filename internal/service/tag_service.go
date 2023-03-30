package service

import "github.com/gin-gonic/gin"

type TagService interface {
	GetTagList(ctx *gin.Context) (interface{}, error)
}
