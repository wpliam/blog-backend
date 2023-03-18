package comment

import (
	"blog-backend/model"
	"blog-backend/repo/auth/jwtauth"
	"blog-backend/repo/mdb"
	"github.com/gin-gonic/gin"
)

type AddCommentReq struct {
	Comment *model.Comment `json:"comment" binding:"required"`
}

func AddComment(ctx *gin.Context) (interface{}, error) {
	var req *AddCommentReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	userID, err := jwtauth.DefaultJwtAuth.ParseUserID(ctx)
	if err != nil {
		return nil, err
	}
	req.Comment.UserID = userID
	if err = mdb.GetGormClient().SetCommentInfo(req.Comment); err != nil {
		return nil, err
	}
	return nil, nil
}
