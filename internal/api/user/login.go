package user

import (
	"blog-backend/constant"
	"blog-backend/model"
	"blog-backend/repo/auth/jwtauth"
	"blog-backend/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wpliap/common-wrap/log"
	"time"
)

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginReply struct {
	Token string      `json:"token"`
	User  *model.User `json:"user"`
}

// LoginImpl 登录实现
func (u *userImpl) LoginImpl(ctx *gin.Context, req *LoginReq) (*LoginReply, error) {
	dbCli := u.GetGormProxy()
	accountInfo, err := dbCli.GetAccountInfo(req.Username)
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
	userInfo, err := dbCli.GetUserInfo(accountInfo.ID)
	if err != nil {
		return nil, err
	}
	if err = u.GetRedisProxy().Set(ctx, token, accountInfo.ID, constant.LoginRedisValidTime); err != nil {
		log.Errorf("Login redis set err:%v", err)
	}
	go func() {
		field := make(map[string]interface{})
		field["ip"] = util.GetClientIP(ctx)
		field["last_login_time"] = time.Now().Format(constant.TimeLayout)
		if err = u.GetGormProxy().UpdateUserInfo(accountInfo.ID, field); err != nil {
			log.Errorf("Login UpdateUserInfo err:%v", err)
		}
	}()
	rsp := &LoginReply{
		Token: token,
		User:  userInfo,
	}
	return rsp, nil
}
