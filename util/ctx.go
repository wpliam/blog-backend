package util

import (
	"github.com/gin-gonic/gin"
	"strings"
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

// GetStaffName 获取登录人
func GetStaffName(ctx *gin.Context) string {
	return ctx.GetHeader("staffname")
}

// GetArticleID 获取文章ID
func GetArticleID(ctx *gin.Context) int64 {
	return ParseInt64(ctx.Param("articleID"))
}

// GetToken 获取token
func GetToken(ctx *gin.Context) string {
	return strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
}
