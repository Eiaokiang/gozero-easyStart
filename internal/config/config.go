package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
}

func init() {
	configFile := "etc/easystart-api.yaml"
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic("配置文件加载失败：" + configFile)
	}
	fmt.Println("加载配置文件：" + configFile)
	// 监控配置文件变化
	viper.WatchConfig()
}
