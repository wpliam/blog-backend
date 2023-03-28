package signauth

import (
	"github.com/gin-gonic/gin"
)

const (
	appid = "liwenping"
	sKey  = "pingguo"
)

// SignInfo 签名鉴权
type SignInfo struct {
	Appid     string `json:"appid" binding:"required"`
	Timestamp int64  `json:"timestamp" binding:"required"`

}

// CheckSign 检查签名
func CheckSign(ctx *gin.Context) error {
	return nil
}
