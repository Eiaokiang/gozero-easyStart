package main

import (
	"easyStart/pkg/store"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"github.com/zeromicro/go-zero/core/logx"

	"easyStart/internal/config"
	"easyStart/internal/handler"
	"easyStart/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/easystart-api.yaml", "the config file")

func main() {
	flag.Parse()

	//配置文件
	var c config.Config
	conf.MustLoad(*configFile, &c)
	//日志设置 需要写在MustNewServer 之前 否则无法生效
	logx.MustSetup(logx.LogConf{
		Encoding:   viper.GetString("logx.encoding"),
		TimeFormat: viper.GetString("logx.timeFormat"),
		Level:      viper.GetString("logx.level"),
	})

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	//数据库
	dbStore := store.NewDbStore()
	//redis
	redisStore := store.NewRedisStore()

	ctx := svc.NewServiceContext(c, dbStore, redisStore)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
