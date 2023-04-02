package service

import "github.com/gin-gonic/gin"

type SharedService interface {
	AddViewCount(ctx *gin.Context) error
	GiveThumb(ctx *gin.Context) (interface{}, error)
	GiveFollow(ctx *gin.Context) (interface{}, error)
	GiveCollect(ctx *gin.Context) (interface{}, error)
}
