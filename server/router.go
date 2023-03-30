package server

import (
	"blog-backend/internal/api/article"
	"blog-backend/internal/api/banner"
	"blog-backend/internal/api/category"
	"blog-backend/internal/api/tag"
	"blog-backend/internal/api/user"
	"github.com/gin-gonic/gin"
	"github.com/wpliap/common-wrap/errs"
	"net/http"
)

// initRouter 服务路由初始化
func (s *Server) initRouter() {
	apiGroup := s.router.Group("api")
	s.initArticleRouter(apiGroup)
	s.initBannerRouter(apiGroup)
	s.initCategoryRouter(apiGroup)
	s.initTagRouter(apiGroup)
	s.initUserRouter(apiGroup)
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
}

type wrapperHandler func(ctx *gin.Context) (interface{}, error)

func (s *Server) wrapperHandler(h wrapperHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := h(ctx)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": errs.Code(err),
				"msg":  errs.Msg(err),
				"data": nil,
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "",
			"data": data,
		})
	}
}
