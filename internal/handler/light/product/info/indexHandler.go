package info

import (
	"gitee.com/unitedrhino/share/errors"
	"gitee.com/unitedrhino/share/result"
	"github.com/zeromicro/go-zero/rest/httpx"
	"lightsvr/internal/logic/light/product/info"
	"lightsvr/internal/svc"
	"lightsvr/internal/types"
	"net/http"
)

// 获取产品详情列表
func IndexHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductInfoIndexReq
		if err := httpx.Parse(r, &req); err != nil {
			result.Http(w, r, nil, errors.Parameter.WithMsg("入参不正确:"+err.Error()))
			return
		}

		l := info.NewIndexLogic(r.Context(), svcCtx)
		resp, err := l.Index(&req)
		result.Http(w, r, resp, err)
	}
}
