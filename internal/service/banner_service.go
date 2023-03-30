package service

import "github.com/gin-gonic/gin"

type BannerService interface {
	GetBannerCard(ctx *gin.Context) (interface{}, error)
}
