package user

import (
	"blog-backend/repo/auth/jwtauth"
	"blog-backend/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wpliap/common-wrap/log"
)

type RefreshTokenReq struct {
	Uid   int64  `json:"uid" binding:"gt=1"`
	Token string `json:"token" binding:"required"`
}

type RefreshTokenReply struct {
	Token string `json:"token"`
}

// RefreshTokenImpl 中间件层有做redis的校验,这里只需要校验id与请求的id是否一致
func (u *userImpl) RefreshTokenImpl(ctx *gin.Context, req *RefreshTokenReq) (*RefreshTokenReply, error) {
	uid := util.GetUid(ctx)
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
