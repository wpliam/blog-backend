package middleware

import (
	"blog-backend/repo/auth/jwtauth"
	"blog-backend/util/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Middleware 中间件
type Middleware struct {
	*jwtauth.JwtAuth
}

// LoginAuth 登录中间件
func (middle *Middleware) LoginAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := middle.ParseToken(ctx); err != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, &resp.Response{
				Code: http.StatusUnauthorized,
				Msg:  "未登录,请先登录",
			})
			return
		}
		ctx.Next()
	}
}
