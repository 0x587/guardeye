package handler

import (
	"net/http"

	"github.com/0x587/guardeye/api/internal/logic"
	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func NodeWsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NodeWsReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error(err)
			logx.Error(r.URL.String())
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewNodeWsLogic(r.Context(), svcCtx)
		err := l.NodeWs(&req, w, r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
