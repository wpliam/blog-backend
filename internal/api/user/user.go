package user

import (
	"blog-backend/constant"
	"blog-backend/internal/service"
	"blog-backend/repo/auth/jwtauth"
	"blog-backend/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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
	// 新token加入redis
	if err = u.GetRedisProxy().Set(ctx, token, uid, constant.LoginRedisValidTime); err != nil {
		return nil, err
	}
	// 将旧token删除
	if err = u.GetRedisProxy().Del(ctx, req.Token); err != nil {
		log.Errorf("RefreshToken del old token err:%v token:%v", err, token)
	}
	rsp := &RefreshTokenReply{
		Token: token,
	}
	log.Infof("RefreshToken success req.token:%s token:", req.Token, token)
	return rsp, nil
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

// GetUserCollectList 获取用户收藏列表
func (u *userImpl) GetUserCollectList(ctx *gin.Context) (interface{}, error) {
	var req *GetUserCollectListReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	redisCli := u.GetRedisProxy()
	members, err := redisCli.SMembers(ctx, util.GetUserCollectKey(req.Uid))
	if err != nil {
		return nil, err
	}
	if len(members) == 0 {
		return nil, nil
	}
	ids := make([]int64, len(members))
	for _, id := range members {
		userID := util.ParseInt64(id)
		if userID == 0 {
			continue
		}
		ids = append(ids, userID)
	}
	esCli := u.GetElasticProxy()
	articles, err := esCli.GetArticleList(ctx, ids)
	if err != nil {
		return nil, err
	}
	rsp := &GetUserCollectListRsp{
		Articles: articles,
	}
	return rsp, nil
}
