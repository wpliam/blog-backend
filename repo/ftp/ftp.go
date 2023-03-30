package ftp

import (
	"github.com/jlaffaye/ftp"
)

type FtpClient struct {
	*ftp.ServerConn
}
