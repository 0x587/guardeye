package downstream

import (
	"net/http"

	"github.com/0x587/guardeye/api/internal/logic/downstream"
	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CallHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownstreamCallReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := downstream.NewCallLogic(r.Context(), svcCtx)
		resp, err := l.Call(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
