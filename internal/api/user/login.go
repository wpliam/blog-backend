package user

import (
	"blog-backend/repo/auth/jwtauth"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginReply struct {
	Token string `json:"token"`
}

func (u *userImpl) LoginImpl(ctx *gin.Context, req *LoginReq) (*LoginReply, error) {
	accountInfo, err := u.GetGormProxy().GetAccountInfo(req.Username)
	if err != nil {
		return nil, err
	}
	if req.Password != accountInfo.Password {
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
