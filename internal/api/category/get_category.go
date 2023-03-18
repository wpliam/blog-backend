package category

import (
	"blog-backend/model"
	"blog-backend/repo/mdb"
	"github.com/gin-gonic/gin"
)

type GetCategoryReq struct {
	Cid int64 `json:"cid"`
}

type GetCategoryRsp struct {
	Category []*model.Category `json:"categorys"`
}

// GetCategory 获取分类
func GetCategory(ctx *gin.Context) (interface{}, error) {
	dbCli := mdb.GetGormClient()
	categoryList, err := dbCli.GetCategoryList()
	if err != nil {
		return nil, err
	}
	rsp := &GetCategoryRsp{
		Category: categoryList,
	}
	return rsp, nil
}
