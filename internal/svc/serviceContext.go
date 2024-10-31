package svc

import (
	"gitee.com/unitedrhino/core/service/apisvr/exportMiddleware"
	"gitee.com/unitedrhino/core/service/syssvr/client/log"
	role "gitee.com/unitedrhino/core/service/syssvr/client/rolemanage"
	tenant "gitee.com/unitedrhino/core/service/syssvr/client/tenantmanage"
	user "gitee.com/unitedrhino/core/service/syssvr/client/usermanage"
	"gitee.com/unitedrhino/share/ctxs"
	"gitee.com/unitedrhino/share/eventBus"
	"gitee.com/unitedrhino/share/stores"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/dmsvr/client/deviceinteract"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"lightsvr/internal/config"
	"lightsvr/internal/repo/relationDB"
)

type ServiceContext struct {
	Config         config.Config
	InitCtxsWare   rest.Middleware
	CheckTokenWare rest.Middleware
	DeviceInteract deviceinteract.DeviceInteract
	FastEvent      *eventBus.FastEvent
	Store          kv.Store
	NodeID         int64
}

func NewServiceContext(c config.Config) *ServiceContext {
	var (
		deviceInteract deviceinteract.DeviceInteract
		checkTokenWare rest.Middleware
	)
	stores.InitConn(c.Database)
	err := relationDB.Migrate(c.Database)
	logx.Must(err)
	deviceInteract = deviceinteract.NewDeviceInteract(zrpc.MustNewClient(c.DmRpc.Conf))

	nodeID := utils.GetNodeID(c.CacheRedis, c.Name)

	fastEvent, err := eventBus.NewFastEvent(c.Event, c.Name, nodeID)
	logx.Must(err)
	ur := user.NewUserManage(zrpc.MustNewClient(c.SysRpc.Conf))
	ro := role.NewRoleManage(zrpc.MustNewClient(c.SysRpc.Conf))
	tm := tenant.NewTenantManage(zrpc.MustNewClient(c.SysRpc.Conf))
	lo := log.NewLog(zrpc.MustNewClient(c.SysRpc.Conf))
	checkTokenWare = exportMiddleware.NewCheckTokenWareMiddleware(ur, ro, tm, lo).Handle //如果需要用到对应的服务,则使用这种方式
	checkTokenWare = exportMiddleware.NewCheckTokenWareMiddleware2(c.SysRpc).Handle      //如果不需要用户等rpc,则使用这种方式即可
	return &ServiceContext{
		Config:         c,
		InitCtxsWare:   ctxs.InitMiddleware,
		CheckTokenWare: checkTokenWare,
		FastEvent:      fastEvent,
		DeviceInteract: deviceInteract,
		Store:          kv.NewStore(c.CacheRedis),
	}
}
