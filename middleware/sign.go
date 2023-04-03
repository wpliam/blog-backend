package middleware

import (
	"blog-backend/util"
	"blog-backend/util/resp"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wpliap/common-wrap/log"
	"net/http"
)

const (
	Appid = "appid"
	SKey  = "key"
)

type SignInfo struct {
	Appid     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
	Sign      string `json:"sign"`
}

// CheckSign 鉴权检查
func (middle *Middleware) CheckSign() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		signInfo := &SignInfo{}
		request := ctx.Request
		if request.Method == http.MethodGet || request.Method == http.MethodPost {
			signInfo.Appid = ctx.Query("appid")
			signInfo.Timestamp = util.ParseInt64(ctx.Query("timestamp"))
			signInfo.Sign = ctx.Query("sign")
		}
		if err := checkMd5(signInfo); err != nil {
			log.Errorf("CheckSign checkMd5 err:%v", err)
			resp.ResponseForbidden(ctx, err.Error())
			return
		}
		ctx.Next()
	}
}

func checkMd5(sign *SignInfo) error {
	if sign.Appid != Appid {
		return fmt.Errorf("appid not exist")
	}
	h := md5.New()
	h.Write([]byte(fmt.Sprintf("%s%d%s", sign.Appid, sign.Timestamp, SKey)))
	str := hex.EncodeToString(h.Sum(nil))
	if sign.Sign != str {
		return fmt.Errorf("sign auth faile")
	}
	return nil
}
