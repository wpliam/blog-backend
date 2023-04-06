package file

import (
	"blog-backend/internal/service"
	"blog-backend/repo/ftp"
	"blog-backend/util/resp"
	"github.com/gin-gonic/gin"
	"github.com/wpliap/common-wrap/log"
	"io"
	"net/http"
)

func NewDownloadService() service.DownloadService {
	return &downloadImpl{}
}

type downloadImpl struct {
}

func (d *downloadImpl) Download(ctx *gin.Context) {
	filepath := ctx.Param("filepath")
	ftpCli, err := ftp.NewFtpProxy()
	if err != nil {
		log.Errorf("Download ftp conn err:%v", err)
		resp.ResponseFail(ctx, err)
		return
	}
	retr, err := ftpCli.Retr(filepath)
	if err != nil {
		log.Errorf("Download ftp retr err:%v", err)
		resp.ResponseFail(ctx, err)
		return
	}
	content, err := io.ReadAll(retr)
	if err != nil {
		log.Errorf("Download ReadAll err:%v", err)
		resp.ResponseFail(ctx, err)
		return
	}
	ctx.Data(http.StatusOK, "application/octet-stream", content)
}
