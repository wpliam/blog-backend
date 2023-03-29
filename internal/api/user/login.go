package user

import (
	"blog-backend/global/proxy"
	"blog-backend/repo/auth/jwtauth"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginReply struct {
	Token string `json:"token"`
}

func (l *Login) Invoke(ctx *gin.Context, proxy proxy.Proxy) (interface{}, error) {
	if err := ctx.ShouldBindJSON(&l); err != nil {
		return nil, err
	}
	accountInfo, err := proxy.GetGormProxy().GetAccountInfo(l.Username)
	if err != nil {
		return nil, err
	}
	if l.Password != accountInfo.Password {
		return nil, fmt.Errorf("账号或密码错误")
	}
	token, err := jwtauth.DefaultJwtAuth.GenToken(ctx, accountInfo.ID)
	if err != nil {
		return nil, err
	}
	rsp := &LoginReply{
		Token: token,
	}
	return rsp, nil
}
