package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"sync/atomic"
)

var conf atomic.Value

func init() {
	conf.Store(&AppConf{})
}

type AppConf struct {
	Port    string       `yaml:"port"`
	Mysql   *MysqlConf   `yaml:"mysql"`
	Redis   *RedisConf   `yaml:"redis"`
	Ftp     *FtpConf     `yaml:"ftp"`
	Elastic *ElasticConf `yaml:"elastic"`
	Sign    *SignConf    `yaml:"sign"`
}

type MysqlConf struct {
	Target string `yaml:"target"`
}

type RedisConf struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
}

type FtpConf struct {
	Addr     string `yaml:"addr"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type ElasticConf struct {
	Addr     string `yaml:"addr"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type SignConf struct {
	Appid string `yaml:"appid"`
	SKey  string `yaml:"s_key"`
}

// LoadConfig 加载配置
func LoadConfig() error {
	appConf := &AppConf{}
	content, err := os.ReadFile("application.yaml")
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(content, appConf); err != nil {
		return err
	}
	conf.Store(appConf)
	return nil
}

func getAppConf() *AppConf {
	return conf.Load().(*AppConf)
}

func GetPort() string {
	return getAppConf().Port
}

func GetMysqlConf() *MysqlConf {
	return getAppConf().Mysql
}

func GetFtpConf() *FtpConf {
	return getAppConf().Ftp
}

func GetRedisConf() *RedisConf {
	return getAppConf().Redis
}

func GetElasticConf() *ElasticConf {
	return getAppConf().Elastic
}

func GetSignConf() *SignConf {
	return getAppConf().Sign
}