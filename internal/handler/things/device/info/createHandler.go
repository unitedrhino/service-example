package info

import (
	"net/http"

	"gitee.com/unitedrhino/share/result"

	"lightsvr/internal/logic/things/device/info"
	"lightsvr/internal/svc"
)

// 新增设备
func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := info.NewCreateLogic(r.Context(), svcCtx)
		err := l.Create()
		result.Http(w, r, nil, err)
	}
}
