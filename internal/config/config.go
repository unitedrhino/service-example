package config

import (
	"gitee.com/unitedrhino/share/conf"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	CacheRedis cache.ClusterConf
	Event      conf.EventConf     //和things内部交互的设置
	DmRpc      conf.RpcClientConf `json:",optional"`
	SysRpc     conf.RpcClientConf `json:",optional"`
	Database   conf.Database
}
