package ftp

import (
	"github.com/jlaffaye/ftp"
)

type FtpClient struct {
	*ftp.ServerConn
}

func NewFtpClient(conn *ftp.ServerConn) *FtpClient {
	return &FtpClient{
		conn,
	}
}
