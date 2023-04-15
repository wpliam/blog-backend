package middleware

import (
	"blog-backend/util"
	"github.com/gin-gonic/gin"
	"time"
)

// SetUid 这里只做uid的赋值
func (middle *Middleware) SetUid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid := middle.parseUid(ctx)
		ctx.Set("uid", uid)
		ctx.Next()
	}
}

func (middle *Middleware) parseUid(ctx *gin.Context) int64 {
	token := util.GetToken(ctx)
	if token == "" {
		return 0
	}
	uid, err := middle.RedisProxy.GetInt64(ctx, token)
	if err == nil {
		return uid
	}
	claims, err := middle.Jwt.ParseClaims(ctx)
	if err != nil {
		return 0
	}
	if claims.VerifyExpiresAt(time.Now(), false) {
		return claims.Uid
	}
	return 0
}
