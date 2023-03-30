package middleware

import (
	"blog-backend/util"
	"blog-backend/util/resp"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// Refresh 刷新token
func (middle *Middleware) Refresh() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := middle.refresh(ctx); err != nil {
			resp.ResponseUnauthorized(ctx, "刷新失败")
			return
		}
		ctx.Next()
	}
}

func (middle *Middleware) refresh(ctx *gin.Context) error {
	// 从头部获取token,刷新的前提条件是必须用旧token来换新token
	token := middle.Jwt.GetToken(ctx)
	if token == "" {
		return fmt.Errorf("token is empty")
	}
	uid, err := middle.RedisProxy.Get(ctx, token)
	// redis中不存在记录,可能是主动退出token已经删除获取token在redis中已过期,不允许刷新
	if err == redis.Nil {
		return fmt.Errorf("token not exist")
	}
	ctx.Set("uid", util.ParseInt64(uid))
	return nil
}
