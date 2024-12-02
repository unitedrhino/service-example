package info

import (
	"net/http"

	"gitee.com/unitedrhino/share/result"

	"lightsvr/internal/logic/light/device/info"
	"lightsvr/internal/svc"
)

// 删除设备
func DeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := info.NewDeleteLogic(r.Context(), svcCtx)
		err := l.Delete()
		result.Http(w, r, nil, err)
	}
}
