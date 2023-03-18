package ftp

import (
	"github.com/jlaffaye/ftp"
	aftp "github.com/wpliap/common-wrap/ftp"
	"sync"
)

var (
	ftpProxy = make(map[string]*Client)
	rw       sync.RWMutex
)

const defaultFtpName = "ftp"

type Client struct {
	*ftp.ServerConn
}

func GetFtpClient() *Client {
	rw.RLock()
	defer rw.RUnlock()
	client, ok := ftpProxy[defaultFtpName]
	if ok && client != nil {
		return client
	}
	ftpCli := aftp.NewFtpProxy("blog.ftp")
	client = &Client{
		ftpCli,
	}
	ftpProxy[defaultFtpName] = client
	return client
}
