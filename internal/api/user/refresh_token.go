package user

import (
	"blog-backend/repo/auth/jwtauth"
	"blog-backend/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/wpliap/common-wrap/log"
)

type RefreshTokenReq struct {
	Uid   int64  `json:"uid"`
	Token string `json:"token" binding:"required"`
}

type RefreshTokenReply struct {
	Token string `json:"token"`
}

// RefreshTokenImpl ...
func (u *userImpl) RefreshTokenImpl(ctx *gin.Context, req *RefreshTokenReq) (*RefreshTokenReply, error) {
	id, err := u.GetRedisProxy().Get(ctx, req.Token)
	if err == redis.Nil {
		return nil, fmt.Errorf("token not exist")
	}
	uid := util.ParseInt64(id)
	if uid != req.Uid {
		return nil, fmt.Errorf("id error")
	}
	token, err := jwtauth.DefaultJwtAuth.GenToken(ctx, uid)
	if err != nil {
		return nil, err
	}
	// 将旧token删除
	if err = u.GetRedisProxy().Del(ctx, req.Token); err != nil {
		log.Errorf("RefreshTokenImpl del old token err:%v token:%v", err, token)
	}
	rsp := &RefreshTokenReply{
		Token: token,
	}
	return rsp, nil
}
