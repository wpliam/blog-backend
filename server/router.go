package server

import (
	"blog-backend/util/resp"
	"github.com/gin-gonic/gin"
)

// initRouter 服务路由初始化
func (s *Server) initRouter() {
	s.router.Use(s.middle.Options())
	apiGroup := s.router.Group("api")
	{
		// 文件下载
		apiGroup.GET("download/:filepath", s.downloadService.Download)
	}

	signGroup := apiGroup.Use(s.middle.SetUid(), s.middle.CheckSign())
	{
		// 文件上传
		signGroup.POST("upload", s.wrapperHandler(s.uploadService.Upload))
		// 搜索文章列表
		signGroup.POST("search_article_list", s.wrapperHandler(s.articleService.SearchArticleList))
		// 获取热门文章
		signGroup.GET("get_hot_article", s.wrapperHandler(s.articleService.GetHotArticle))
		// 读取文章
		signGroup.GET("read_article/:articleID", s.wrapperHandler(s.articleService.ReadArticle))
		// 获取文章归档
		signGroup.GET("get_article_archive", s.wrapperHandler(s.articleService.GetArticleArchive))
		// 搜索关键词流水
		signGroup.POST("search_keyword_flow", s.wrapperHandler(s.articleService.SearchKeywordFlow))

		// 添加访问量
		signGroup.GET("add_view_count/:articleID", s.wrapper(s.sharedService.AddViewCount))

		// 获取评论
		signGroup.POST("get_comment", s.wrapperHandler(s.commentService.GetComment))

		// 登录
		signGroup.POST("login", s.wrapperHandler(s.userService.Login))
		// 退出
		signGroup.POST("logout", s.wrapper(s.userService.Logout))
		// 刷新token
		signGroup.POST("refresh_token", s.wrapperHandler(s.userService.RefreshToken))
		// 统计用户信息
		signGroup.GET("census_user_info/:uid", s.wrapperHandler(s.userService.CensusUserInfo))
		// 获取用户信息
		signGroup.GET("get_user_info/:uid", s.wrapperHandler(s.userService.GetUserInfo))
		// 获取用户收藏列表
		signGroup.POST("get_user_collect_list", s.wrapperHandler(s.userService.GetUserCollectList))

		// 获取banner卡片
		signGroup.GET("get_banner_card", s.wrapperHandler(s.bannerService.GetBannerCard))

		// 获取标签列表
		signGroup.GET("get_tag_card", s.wrapperHandler(s.tagService.GetTagList))

		// 获取分类卡片
		signGroup.GET("get_category_card", s.wrapperHandler(s.categoryService.GetCategoryCard))
		// 获取分类列表
		signGroup.GET("get_category_list", s.wrapperHandler(s.categoryService.GetCategoryList))
	}

	loginGroup := signGroup.Use(s.middle.LoginAuth())
	{
		// 文章审核
		loginGroup.POST("article_review", s.wrapper(s.articleService.ArticleReview))
		// 写文章
		loginGroup.POST("write_article", s.middle.LoginAuth(), s.wrapper(s.articleService.WriteArticle))
		// 收藏
		loginGroup.POST("give_collect", s.wrapperHandler(s.sharedService.GiveCollect))
		// 点赞
		loginGroup.POST("give_thumb", s.wrapperHandler(s.sharedService.GiveThumb))
		// 关注
		loginGroup.POST("give_follow", s.wrapperHandler(s.sharedService.GiveFollow))
		// 签到
		loginGroup.GET("punch_clock", s.wrapper(s.sharedService.PunchClock))
		// 统计签到信息
		loginGroup.POST("census_clock_info", s.wrapperHandler(s.sharedService.CensusClockInfo))
		// 添加评论
		loginGroup.POST("add_comment", s.wrapper(s.commentService.AddComment))
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
