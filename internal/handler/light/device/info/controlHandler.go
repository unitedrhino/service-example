package info

import (
	"net/http"

	"gitee.com/unitedrhino/share/result"

	"lightsvr/internal/logic/light/device/info"
	"lightsvr/internal/svc"
)

// 设备控制
func ControlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := info.NewControlLogic(r.Context(), svcCtx)
		err := l.Control()
		result.Http(w, r, nil, err)
	}
}
