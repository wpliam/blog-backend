package article

import (
	"blog-backend/model"
	"blog-backend/repo/mdb"
	"github.com/gin-gonic/gin"
)

type ReadArticleReq struct {
	ArticleID int64 `json:"articleID"`
}

type ReadArticleRsp struct {
	Article *model.Article `json:"article"`
}

// ReadArticle 读取文章
func ReadArticle(ctx *gin.Context) (interface{}, error) {
	var req *ReadArticleReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	dbClient := mdb.GetGormClient()
	articleInfo, err := dbClient.GetArticleInfo(req.ArticleID)
	if err != nil {
		return nil, err
	}
	rsp := &ReadArticleRsp{
		Article: articleInfo,
	}
	return rsp, nil
}
