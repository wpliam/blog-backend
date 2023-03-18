package tag

import (
	"blog-backend/model"
	"blog-backend/repo/mdb"
	"github.com/gin-gonic/gin"
)

type GetTagReq struct {
	ArticleID int64 `json:"articleID"`
}

type GetTagRsp struct {
	Tags []*model.Tag `json:"tags"`
}

// GetTag 获取tag
func GetTag(ctx *gin.Context) (interface{}, error) {
	var req *GetTagReq
	var tagIDList []int64
	var err error
	dbCli := mdb.GetGormClient()
	if req.ArticleID > 0 {
		tagIDList, err = dbCli.GetArticleTagID(req.ArticleID)
		if err != nil {
			return nil, err
		}
		if len(tagIDList) == 0 {
			return nil, nil
		}
	}
	tagList, err := dbCli.GetTagList(tagIDList...)
	if err != nil {
		return nil, err
	}
	rsp := &GetTagRsp{
		Tags: tagList,
	}
	return rsp, nil
}
