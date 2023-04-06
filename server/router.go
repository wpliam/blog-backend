package server

import (
	"blog-backend/internal/api"
	"blog-backend/internal/file"
	"blog-backend/util/resp"
	"github.com/gin-gonic/gin"
)

// initRouter 服务路由初始化
func (s *Server) initRouter() {
	s.router.Use(s.middle.Options())

	apiGroup := s.router.Group("api")

	upload := file.NewUploadService()
	apiGroup.POST("upload", s.wrapperHandler(upload.Upload))
	download := file.NewDownloadService()
	apiGroup.GET("download/:filepath", download.Download)

	apiGroup.Use(s.middle.SetUid(), s.middle.CheckSign())
	s.initArticleRouter(apiGroup)
	s.initBannerRouter(apiGroup)
	s.initCategoryRouter(apiGroup)
	s.initTagRouter(apiGroup)
	s.initUserRouter(apiGroup)
	s.initSharedRouter(apiGroup)
}

func (s *Server) initArticleRouter(apiGroup *gin.RouterGroup) {
	a := api.NewArticleService(s.proxy)
	apiGroup.POST("search_article_list", s.wrapperHandler(a.SearchArticleList))
	apiGroup.GET("get_hot_article", s.wrapperHandler(a.GetHotArticle))
	apiGroup.GET("read_article/:articleID", s.wrapperHandler(a.ReadArticle))
	apiGroup.GET("get_article_archive", s.wrapperHandler(a.GetArticleArchive))
	apiGroup.POST("search_keyword_flow", s.wrapperHandler(a.SearchKeywordFlow))

	apiGroup.POST("write_article", s.middle.LoginAuth(), s.wrapper(a.WriteArticle))
}

func (s *Server) initBannerRouter(apiGroup *gin.RouterGroup) {
	b := api.NewBannerService(s.proxy)
	apiGroup.GET("get_banner_card", s.wrapperHandler(b.GetBannerCard))
}

func (s *Server) initCategoryRouter(apiGroup *gin.RouterGroup) {
	c := api.NewCategoryService(s.proxy)
	apiGroup.GET("get_category_card", s.wrapperHandler(c.GetCategoryCard))
	apiGroup.GET("get_category_list", s.wrapperHandler(c.GetCategoryList))
}

func (s *Server) initTagRouter(apiGroup *gin.RouterGroup) {
	t := api.NewTagService(s.proxy)
	apiGroup.GET("get_tag_card", s.wrapperHandler(t.GetTagList))
}

func (s *Server) initUserRouter(apiGroup *gin.RouterGroup) {
	u := api.NewUserService(s.proxy)
	apiGroup.POST("login", s.wrapperHandler(u.Login))
	apiGroup.POST("logout", s.wrapper(u.Logout))
	apiGroup.POST("refresh_token", s.wrapperHandler(u.RefreshToken))
	apiGroup.GET("census_user_info/:uid", s.wrapperHandler(u.CensusUserInfo))
	apiGroup.GET("get_user_info/:uid", s.wrapperHandler(u.GetUserInfo))

	apiGroup.POST("get_user_collect_list", s.wrapperHandler(u.GetUserCollectList))
}

func (s *Server) initSharedRouter(apiGroup *gin.RouterGroup) {
	share := api.NewSharedService(s.proxy)
	apiGroup.GET("add_view_count/:articleID", s.wrapper(share.AddViewCount))

	apiGroup.POST("give_collect", s.middle.LoginAuth(), s.wrapperHandler(share.GiveCollect))
	apiGroup.POST("give_thumb", s.middle.LoginAuth(), s.wrapperHandler(share.GiveThumb))
	apiGroup.POST("give_follow", s.middle.LoginAuth(), s.wrapperHandler(share.GiveFollow))
	apiGroup.GET("punch_clock", s.middle.LoginAuth(), s.wrapper(share.PunchClock))
	apiGroup.POST("census_clock_info", s.middle.LoginAuth(), s.wrapperHandler(share.CensusClockInfo))
}

type wrapper func(ctx *gin.Context) error

func (s *Server) wrapper(h wrapper) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := h(ctx); err != nil {
			resp.ResponseFail(ctx, err)
			return
		}
		resp.ResponseOk(ctx, nil)
	}
}

type wrapperHandler func(ctx *gin.Context) (interface{}, error)

func (s *Server) wrapperHandler(h wrapperHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := h(ctx)
		if err != nil {
			resp.ResponseFail(ctx, err)
			return
		}
		resp.ResponseOk(ctx, data)
	}
}
