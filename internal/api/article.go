package api

import (
	"blog-backend/constant"
	"blog-backend/internal/service"
	"blog-backend/model"
	"blog-backend/model/jsonagree"
	"blog-backend/util"
	"blog-backend/util/thread"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wpliap/common-wrap/log"
	"strconv"
	"strings"
)

func NewArticleService(proxyService service.ProxyService) service.ArticleService {
	return &articleImpl{
		proxyService,
	}
}

type articleImpl struct {
	service.ProxyService
}

// SearchArticleList 搜索文章列表
func (a *articleImpl) SearchArticleList(ctx *gin.Context) (interface{}, error) {
	var req *jsonagree.SearchArticleListReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	switch req.SearchType {
	case 1:
		articles, err := a.GetElasticProxy().SearchRandomArticle(ctx)
		if err != nil {
			return nil, err
		}
		rsp := &jsonagree.SearchArticleListReply{
			Articles: articles,
		}
		return rsp, nil
	default:
		if req.Keyword != "" {
			if err := a.GetGormProxy().AddSearchFlow(req.Keyword); err != nil {
				log.Errorf("SearchArticleListImpl AddSearchFlow err:%v", err)
			}
		}
		articles, total, err := a.GetElasticProxy().SearchArticleList(ctx, req)
		if err != nil {
			log.Errorf("SearchArticle search err:%v req:%+v", err, req)
			return nil, err
		}
		req.Page.SetTotal(total)
		rsp := &jsonagree.SearchArticleListReply{
			Page:     req.Page,
			Articles: articles,
		}
		return rsp, nil
	}
}

// GetArticleArchive 获取文章归档
func (a *articleImpl) GetArticleArchive(ctx *gin.Context) (interface{}, error) {
	rsp := &jsonagree.GetArticleArchiveReply{}
	dbCli := a.GetGormProxy()
	handler := make([]func() error, 0)
	handler = append(handler, func() error {
		articles, total, err := a.GetElasticProxy().SearchArticleList(ctx, &jsonagree.SearchArticleListReq{})
		if err != nil {
			return err
		}
		rsp.Article = articleGroupBy(articles)
		rsp.ArticleCount = total
		return nil
	})
	handler = append(handler, func() error {
		tags, err := dbCli.GetTagList()
		if err != nil {
			return err
		}
		rsp.Tags = tags
		rsp.TagCount = int64(len(tags))
		return nil
	})
	handler = append(handler, func() error {
		categoryList, err := dbCli.GetCategoryList()
		if err != nil {
			return err
		}
		rsp.Category = categoryList
		rsp.CategoryCount = int64(len(categoryList))
		return nil
	})
	if err := thread.GoAndWait(handler...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func articleGroupBy(articleList []*model.ArticleContentSummary) map[string][]*model.ArticleContentSummary {
	articleGroup := make(map[string][]*model.ArticleContentSummary)
	for _, article := range articleList {
		key := article.CreateTime.Format(constant.MonthSubTableSuffix)
		articleGroup[key] = append(articleGroup[key], article)
	}
	return articleGroup
}

// GetHotArticle 获取热门文章
func (a *articleImpl) GetHotArticle(ctx *gin.Context) (interface{}, error) {
	articles, err := a.GetGormProxy().GetHotArticle()
	if err != nil {
		return nil, err
	}
	rsp := &jsonagree.GetArticleReply{
		Articles: articles,
	}
	return rsp, nil
}

// ReadArticle 读取文章
func (a *articleImpl) ReadArticle(ctx *gin.Context) (interface{}, error) {
	articleID, err := strconv.ParseInt(ctx.Param("articleID"), 10, 64)
	if err != nil {
		return nil, err
	}
	dbCli := a.GetGormProxy()
	esCli := a.GetElasticProxy()
	redisCli := a.GetRedisProxy()
	summary, err := esCli.QueryArticleInfo(ctx, articleID)
	if err != nil {
		return nil, err
	}
	rsp := &jsonagree.ReadArticleReply{}
	articleContent := &model.ArticleContentInfo{}
	rsp.Article.ArticleContentSummary = summary
	handler := make([]func() error, 0)
	handler = append(handler, func() error {
		var err error
		articleContent, err = dbCli.GetArticleContentInfo(articleID)
		return err
	})
	handler = append(handler, func() error {
		var err error
		rsp.Next, err = dbCli.GetNextArticle(articleID)
		return err
	})
	handler = append(handler, func() error {
		var err error
		rsp.Prev, err = dbCli.GetPrevArticle(articleID)
		return err
	})
	handler = append(handler, func() error {
		var err error
		rsp.Article.CommentCount, err = dbCli.GetArticleCommentCount(articleID)
		return err
	})
	if len(summary.TagIDs) > 0 {
		handler = append(handler, func() error {
			var err error
			rsp.Tags, err = dbCli.GetTagList(summary.TagIDs...)
			return err
		})
	}
	if len(summary.RecommendIDs) > 0 {
		handler = append(handler, func() error {
			var err error
			rsp.Recommend, err = esCli.QueryArticleList(ctx, summary.RecommendIDs)
			return err
		})
	}
	if err = thread.GoAndWait(handler...); err != nil {
		log.Errorf("ReadArticle err:%v articleID:%s", err, articleID)
		return nil, err
	}
	articleStrID := fmt.Sprintf("%d", articleID)
	likeCount, err := redisCli.ZScore(ctx, constant.ArticleLikeCountKey, articleStrID)
	if err == nil {
		articleContent.LikeCount = int64(likeCount)
	}
	viewCount, err := redisCli.ZScore(ctx, constant.ArticleViewCountKey, articleStrID)
	if err == nil {
		articleContent.ViewCount = int64(viewCount)
	}
	collectCount, err := redisCli.HGet(ctx, constant.ArticleCollectCountKey, articleStrID)
	if err == nil {
		articleContent.CollectCount = collectCount
	}
	uid := util.GetUid(ctx)
	if uid > 0 {
		rsp.Article.IsLike = redisCli.SIsMember(ctx, util.GetUserArticleLikeKey(uid), articleID)
		rsp.Article.IsCollect = redisCli.SIsMember(ctx, util.GetUserCollectKey(uid), articleID)
	}
	rsp.Article.ArticleContentInfo = articleContent
	return rsp, nil
}

// SearchKeywordFlow 搜索关键词流水
func (a *articleImpl) SearchKeywordFlow(ctx *gin.Context) (interface{}, error) {
	var req *jsonagree.SearchKeywordFlowReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	flows, err := a.GetGormProxy().GetSearchFlow(req.Keyword)
	if err != nil {
		return nil, err
	}
	rsp := &jsonagree.SearchKeywordFlowReply{
		Flows: flows,
	}
	return rsp, nil
}

// WriteArticle 写文章
func (a *articleImpl) WriteArticle(ctx *gin.Context) error {
	var req *jsonagree.WriteArticleReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return err
	}
	log.Infof("WriteArticle req:%+v", req)
	article := &model.Article{
		ArticleBaseInfo: model.ArticleBaseInfo{
			CategoryID:  req.Cid,
			UserID:      util.GetUid(ctx),
			Title:       req.Title,
			Abstract:    req.Abstract,
			Status:      0,
			Cover:       req.Cover,
			ArticleType: req.ArticleType,
		},
		TagID:       strings.Join(a.getTagIDs(req.Tags), ","),
		RecommendID: strings.Join(a.getRecommends(req.Recommends), ","),
		Content:     req.Content,
	}
	if err := a.GetGormProxy().AddArticle(article); err != nil {
		log.Errorf("WriteArticle err:%v article:%+v", err, article)
		return err
	}
	return nil
}

func (a *articleImpl) getTagIDs(tags []string) []string {
	dbCli := a.GetGormProxy()
	var tagIDs []string
	for _, name := range tags {
		tag := &model.Tag{
			TagName: name,
		}
		if err := dbCli.FirstOrCreateTag(tag); err != nil {
			log.Errorf("getTagIDs FirstOrCreateTag err:%v", err)
			continue
		}
		tagIDs = append(tagIDs, fmt.Sprintf("%d", tag.ID))
	}
	return tagIDs
}

func (a *articleImpl) getRecommends(list []*model.Article) []string {
	var recommends []string
	for _, item := range list {
		recommends = append(recommends, fmt.Sprintf("%d", item.ID))
	}
	return recommends
}

// ArticleReview 文章审核
func (a *articleImpl) ArticleReview(ctx *gin.Context) error {
	// 超级管理员才有权限审核
	// todo 后面到db拉取权限
	//uid := util.GetUid(ctx)
	//if uid != 1 {
	//	return fmt.Errorf("没有审核权限")
	//}
	var req *jsonagree.ArticleReviewReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return err
	}
	// 默认通过
	dbCli := a.GetGormProxy()
	articleInfo, err := dbCli.GetArticleInfo(req.ArticleID)
	if err != nil {
		return err
	}
	//if articleInfo.Status == 1 {
	//	return fmt.Errorf("文章已是审核通过状态")
	//}
	articleInfo.Status = 1
	if err = dbCli.UpdateArticleStatus(articleInfo); err != nil {
		return err
	}
	if err = a.GetElasticProxy().AddArticleToEs(ctx, articleInfo.ArticleContentSummary()); err != nil {
		return err
	}
	return nil
}
