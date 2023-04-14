package ftp

import (
	"time"

	"blog-backend/repo/config"
	"github.com/jlaffaye/ftp"
)

func NewFtpProxy() (*ftp.ServerConn, error) {
	conf := config.GetFtpConf()
	conn, err := ftp.Dial(conf.Addr, ftp.DialWithTimeout(2000*time.Millisecond))
	if err != nil {
		return nil, err
	}
	if err = conn.Login(conf.Username, conf.Password); err != nil {
		return nil, err
	}
	//if err = conn.Quit(); err != nil {
	//	return nil, err
	//}
	return conn, nil
}
