package service

import "github.com/gin-gonic/gin"

type UserService interface {
	Login(ctx *gin.Context) (interface{}, error)
	Logout(ctx *gin.Context) error
	RefreshToken(ctx *gin.Context) (interface{}, error)
	StaticUserInfo(ctx *gin.Context) (interface{}, error)
	GetUserInfo(ctx *gin.Context) (interface{}, error)
}
