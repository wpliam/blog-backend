package util

import (
	"github.com/gin-gonic/gin"
)

// GetClientIP 获取请求ip
func GetClientIP(ctx *gin.Context) string {
	ip := ctx.ClientIP()
	if ip == "::1" {
		return "127.0.0.1"
	}
	return ip
}

// GetUid 获取uid
func GetUid(ctx *gin.Context) int64 {
	return ctx.GetInt64("uid")
}
