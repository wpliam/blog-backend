package middleware

import (
	"blog-backend/repo/auth/jwtauth"
	"blog-backend/repo/rdb"
	"blog-backend/util/resp"
	"github.com/gin-gonic/gin"
	"time"
)

// Middleware 中间件
type Middleware struct {
	Jwt        *jwtauth.JwtAuth
	RedisProxy *rdb.RedisClient
}

// LoginAuth 登录中间件
func (middle *Middleware) LoginAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, err := middle.Jwt.ParseClaims(ctx)
		if err != nil {
			resp.ResponseUnauthorized(ctx, "未登录,请先登录")
			return
		}
		if !claims.VerifyExpiresAt(time.Now(), false) {
			resp.ResponseUnauthorized(ctx, "登录已过期,请从新登录")
			return
		}
		ctx.Set("uid", claims.Uid)
		ctx.Next()
	}
}
