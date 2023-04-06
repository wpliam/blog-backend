package file

import (
	"blog-backend/constant"
	"blog-backend/internal/service"
	"blog-backend/repo/ftp"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wpliap/common-wrap/log"
	"path"
	"strings"
	"time"
)

const (
	downloadUrl = "http://localhost:8888/api/download/"
)

func NewUploadService() service.UploadService {
	return &uploadImpl{}
}

type uploadImpl struct {
}

func (u *uploadImpl) Upload(ctx *gin.Context) (interface{}, error) {
	file, err := ctx.FormFile("file")
	if err != nil {
		return nil, err
	}
	if !checkFilename(file.Filename) {
		return nil, fmt.Errorf("只支持jpg和png的图片上传")
	}
	reader, err := file.Open()
	if err != nil {
		return nil, err
	}
	ftpCli, err := ftp.NewFtpProxy()
	if err != nil {
		return nil, err
	}
	suffix := path.Ext(file.Filename)
	filename := fmt.Sprintf("%s_%s%s",
		strings.ReplaceAll(uuid.New().String(), "-", ""),
		time.Now().Format(constant.TimeStampLayout), suffix)
	if err = ftpCli.Stor(filename, reader); err != nil {
		return nil, err
	}
	dir, err := ftpCli.CurrentDir()
	if err != nil {
		return nil, err
	}
	result := make(map[string]interface{})
	result["url"] = downloadUrl + filename
	result["dir"] = dir
	log.Infof("upload 图片上传成功 filename:%s", filename)
	return result, nil
}

func checkFilename(filename string) bool {
	switch path.Ext(filename) {
	case ".jpg", ".png":
		return true
	default:
		return false
	}
}
