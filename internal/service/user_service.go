package service

import "github.com/gin-gonic/gin"

type UserService interface {
	Login(ctx *gin.Context) (interface{}, error)
}
