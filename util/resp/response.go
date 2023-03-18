package resp

import (
	"github.com/gin-gonic/gin"
	"github.com/wpliap/common-wrap/errs"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data interface{}
}

func ResponseOk(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, &Response{
		Code: 0,
		Msg:  "",
		Data: data,
	})
}

func ResponseFail(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, &Response{
		Code: errs.Code(err),
		Msg:  errs.Msg(err),
	})
}
