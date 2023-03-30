package server

import (
	"blog-backend/internal/api/article"
	"blog-backend/internal/api/banner"
	"blog-backend/internal/api/category"
	"blog-backend/internal/api/shared"
	"blog-backend/internal/api/tag"
	"blog-backend/internal/api/user"
	"blog-backend/util/resp"
	"github.com/gin-gonic/gin"
)

// initRouter 服务路由初始化
func (s *Server) initRouter() {
	s.router.Use(s.middle.Options())
	apiGroup := s.router.Group("api")
	s.initArticleRouter(apiGroup)
	s.initBannerRouter(apiGroup)
	s.initCategoryRouter(apiGroup)
	s.initTagRouter(apiGroup)
	s.initUserRouter(apiGroup)
	s.initSharedRouter(apiGroup)
}

func (s *Server) initArticleRouter(apiGroup *gin.RouterGroup) {
	a := article.NewArticleService(s.proxy)
	apiGroup.POST("search_article_list", s.wrapperHandler(a.SearchArticleList))
	apiGroup.GET("get_hot_article", s.wrapperHandler(a.GetHotArticle))
	apiGroup.GET("read_article/:articleID", s.wrapperHandler(a.ReadArticle))
	apiGroup.GET("get_article_archive", s.wrapperHandler(a.GetArticleArchive))
}

func (s *Server) initBannerRouter(apiGroup *gin.RouterGroup) {
	b := banner.NewBannerService(s.proxy)
	apiGroup.GET("get_banner_card", s.wrapperHandler(b.GetBannerCard))
}

func (s *Server) initCategoryRouter(apiGroup *gin.RouterGroup) {
	c := category.NewCategoryService(s.proxy)
	apiGroup.GET("get_category_card", s.wrapperHandler(c.GetCategoryCard))
}

func (s *Server) initTagRouter(apiGroup *gin.RouterGroup) {
	t := tag.NewTagService(s.proxy)
	apiGroup.GET("get_tag_card", s.wrapperHandler(t.GetTagList))
}

func (s *Server) initUserRouter(apiGroup *gin.RouterGroup) {
	u := user.NewUserService(s.proxy)
	apiGroup.POST("login", s.wrapperHandler(u.Login))
	apiGroup.POST("logout", s.wrapper(u.Logout))
	apiGroup.POST("refresh_token", s.middle.Refresh(), s.wrapperHandler(u.RefreshToken))
	apiGroup.GET("static_user_info/:uid", s.wrapperHandler(u.StaticUserInfo))
}

func (s *Server) initSharedRouter(apiGroup *gin.RouterGroup) {
	share := shared.NewSharedService(s.proxy)
	apiGroup.POST("add_view_count", s.wrapperHandler(share.AddViewCount))
	loginAuthGroup := apiGroup.Use(s.middle.LoginAuth())
	{
		loginAuthGroup.POST("give_collect", s.wrapperHandler(share.GiveCollect))
		loginAuthGroup.POST("give_thumb", s.wrapperHandler(share.GiveThumb))
		loginAuthGroup.POST("give_follow", s.wrapperHandler(share.GiveFollow))
	}
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
