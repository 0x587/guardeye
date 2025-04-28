package helper

import (
	"net/http"

	"github.com/0x587/guardeye/api/internal/logic/helper"
	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func StartListenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := helper.NewStartListenLogic(r.Context(), svcCtx)
		err := l.StartListen(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
