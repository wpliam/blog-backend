package router

import (
	"blog-backend/util/resp"
	"github.com/gin-gonic/gin"
)

type Handler func(ctx *gin.Context) (interface{}, error)

func WrapHandler(h Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := h(ctx)
		if err != nil {
			resp.ResponseFail(ctx, err)
			return
		}
		resp.ResponseOk(ctx, data)
	}
}
