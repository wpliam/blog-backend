package service

import "github.com/gin-gonic/gin"

type UserService interface {
	Login(ctx *gin.Context) (interface{}, error)
	Logout(ctx *gin.Context) error
	RefreshToken(ctx *gin.Context) (interface{}, error)
	// StaticUserInfo 统计用户文章,评论,粉丝...信息
	StaticUserInfo(ctx *gin.Context) (interface{}, error)
	// GetUserInfo 获取用户信息
	GetUserInfo(ctx *gin.Context) (interface{}, error)
	// GetUserCollectList 获取用户收藏的文章列表
	GetUserCollectList(ctx *gin.Context) (interface{}, error)
}
