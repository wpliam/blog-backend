package conf

import (
	"github.com/spf13/viper"
)

// AppConf 配置信息
type AppConf struct {
	Client *ClientConfig `yaml:"client"`
}

// ClientConfig 客户端配置
type ClientConfig struct {
	Service []struct {
		Name     string `yaml:"name"`
		Host     string `yaml:"host"`
		Port     uint16 `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"service"`
}

// InitConf 初始化配置
func InitConf() error {
	viper.SetConfigFile("blog.yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	blogConf := &AppConf{}
	if err := viper.Unmarshal(blogConf); err != nil {
		return err
	}
	return nil
}
