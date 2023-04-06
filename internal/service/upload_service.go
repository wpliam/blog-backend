package service

import "github.com/gin-gonic/gin"

type UploadService interface {
	Upload(ctx *gin.Context) (interface{}, error)
}
