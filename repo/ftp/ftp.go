package ftp

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"github.com/wpliap/common-wrap/config"
	"time"
)

func NewFtpProxy() (*ftp.ServerConn, error) {
	cfg := config.GetConnConf("blog.ftp")
	addr := fmt.Sprintf("%s:%d", cfg.GetHost(), cfg.GetPort())
	conn, err := ftp.Dial(addr, ftp.DialWithTimeout(time.Duration(cfg.GetTimeout())*time.Millisecond))
	if err != nil {
		return nil, err
	}
	if err = conn.Login(cfg.GetUsername(), cfg.GetPassword()); err != nil {
		return nil, err
	}
	return conn, nil
}
