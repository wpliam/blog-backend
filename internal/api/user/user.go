package user

import (
	"blog-backend/internal/service"
	"blog-backend/repo/auth/jwtauth"
	"blog-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/wpliap/common-wrap/log"
)

func NewUserService(proxyService service.ProxyService) service.UserService {
	return &userImpl{
		proxyService,
	}
}

type userImpl struct {
	service.ProxyService
}

// Login 登录
func (u *userImpl) Login(ctx *gin.Context) (interface{}, error) {
	var req *LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return u.LoginImpl(ctx, req)
}

// Logout 退出
func (u *userImpl) Logout(ctx *gin.Context) error {
	token := jwtauth.DefaultJwtAuth.GetToken(ctx)
	if err := u.GetRedisProxy().Del(ctx, token); err != nil {
		log.Errorf("Logout 退出成功")
	}
	return nil
}

// RefreshToken 刷新token
func (u *userImpl) RefreshToken(ctx *gin.Context) (interface{}, error) {
	var req *RefreshTokenReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	log.Infof("RefreshToken req:%+v", req)
	return u.RefreshTokenImpl(ctx, req)
}

// StaticUserInfo 统计用户信息
func (u *userImpl) StaticUserInfo(ctx *gin.Context) (interface{}, error) {
	uid := util.ParseInt64(ctx.Param("uid"))
	userInfo, err := u.StaticUserInfoImpl(ctx, uid)
	if err != nil {
		return nil, err
	}
	rsp := &StaticUserInfoReply{
		User: userInfo,
	}
	return rsp, nil
}
