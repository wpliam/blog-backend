package ftp

import (
	"github.com/jlaffaye/ftp"
	aftp "github.com/wpliap/common-wrap/ftp"
)

type FtpClient struct {
	*ftp.ServerConn
}

func NewFtpClient(name string) *FtpClient {
	return &FtpClient{
		aftp.NewFtpProxy(name),
	}
}
