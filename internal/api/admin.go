package api

import (
	"blog-backend/internal/service"
	"blog-backend/model/jsonagree"
	"blog-backend/util"
	"github.com/gin-gonic/gin"
)

func NewAdminService(proxyService service.ProxyService) service.AdminService {
	return &adminImpl{
		proxyService,
	}
}

type adminImpl struct {
	service.ProxyService
}

func (a *adminImpl) ArticleReview(ctx *gin.Context) error {

	return nil
}

func (a *adminImpl) GetReadyReviewArticle(ctx *gin.Context) (interface{}, error) {
	uid := util.GetUid(ctx)
	if uid == 0 {
		return nil, nil
	}
	dbCli := a.GetGormProxy()
	userInfo, err := dbCli.GetUserInfo(uid)
	if err != nil {
		return nil, err
	}
	// 普通用户没有文章审核权限,无需返回
	if userInfo.Role == 0 {
		return nil, nil
	}
	var req *jsonagree.GetReadyReviewArticleReq
	if err = ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	article, count, err := dbCli.GetReadyReviewArticle(req.Page)
	if err != nil {
		return nil, err
	}
	req.Page.SetTotal(count)
	rsp := &jsonagree.GetReadyReviewArticleReply{
		Articles: article,
		Page:     req.Page,
	}
	return rsp, nil
}
