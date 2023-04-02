package user

import (
	"blog-backend/internal/service"
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
	token := util.GetToken(ctx)
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
	return u.RefreshTokenImpl(ctx, req)
}

// StaticUserInfo 统计用户信息
func (u *userImpl) StaticUserInfo(ctx *gin.Context) (interface{}, error) {
	uid := util.ParseInt64(ctx.Param("uid"))
	rsp, err := u.StaticUserInfoImpl(ctx, uid)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// GetUserInfo 获取用户信息
func (u *userImpl) GetUserInfo(ctx *gin.Context) (interface{}, error) {
	uid := util.ParseInt64(ctx.Param("uid"))
	userInfo, err := u.GetGormProxy().GetUserInfo(uid)
	if err != nil {
		return nil, err
	}
	rsp := &GetUserInfoReply{
		User: userInfo,
	}
	// 如果登录了,且不是获取用户自己的信息,查询一下关注关系
	loginUid := util.GetUid(ctx)
	if loginUid > 0 && uid != loginUid {
		rsp.IsFollow = u.GetRedisProxy().SIsMember(ctx, util.GetUserFollowKey(loginUid), uid)
	}
	return rsp, nil
}
