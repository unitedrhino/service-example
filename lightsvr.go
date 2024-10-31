package main

import (
	"context"
	"flag"
	"fmt"
	"gitee.com/unitedrhino/share/utils"
	"github.com/zeromicro/go-zero/core/logx"
	_ "github.com/zeromicro/go-zero/core/proc" //开启pprof采集 https://mp.weixin.qq.com/s/yYFM3YyBbOia3qah3eRVQA
	"lightsvr/internal/config"
	"lightsvr/internal/handler"
	"lightsvr/internal/startup"
	"lightsvr/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func main() {
	defer utils.Recover(context.Background())
	flag.Parse()
	logx.DisableStat()
	var c config.Config
	utils.ConfMustLoad("etc/lightsvr.yaml", &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	server.PrintRoutes()
	startup.InitEventBus(ctx)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
