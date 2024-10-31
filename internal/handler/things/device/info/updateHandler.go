package info

import (
	"net/http"

	"gitee.com/unitedrhino/share/result"

	"lightsvr/internal/logic/things/device/info"
	"lightsvr/internal/svc"
)

// 更新设备
func UpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := info.NewUpdateLogic(r.Context(), svcCtx)
		err := l.Update()
		result.Http(w, r, nil, err)
	}
}
