package service

import "github.com/gin-gonic/gin"

type DownloadService interface {
	Download(ctx *gin.Context)
}
