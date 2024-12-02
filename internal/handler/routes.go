// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.1

package handler

import (
	"net/http"

	lightdeviceinfo "lightsvr/internal/handler/light/device/info"
	lightproductinfo "lightsvr/internal/handler/light/product/info"
	"lightsvr/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare},
			[]rest.Route{
				{
					// 设备控制
					Method:  http.MethodPost,
					Path:    "/control",
					Handler: lightdeviceinfo.ControlHandler(serverCtx),
				},
				{
					// 新增设备
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: lightdeviceinfo.CreateHandler(serverCtx),
				},
				{
					// 删除设备
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: lightdeviceinfo.DeleteHandler(serverCtx),
				},
				{
					// 更新设备
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: lightdeviceinfo.UpdateHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/light/device/info"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare},
			[]rest.Route{
				{
					// 添加产品详情
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: lightproductinfo.CreateHandler(serverCtx),
				},
				{
					// 删除产品详情
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: lightproductinfo.DeleteHandler(serverCtx),
				},
				{
					// 获取产品详情列表
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: lightproductinfo.IndexHandler(serverCtx),
				},
				{
					// 获取产品详情详情
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: lightproductinfo.ReadHandler(serverCtx),
				},
				{
					// 更新产品详情
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: lightproductinfo.UpdateHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/light/product/info"),
	)
}
