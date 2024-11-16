package handler

import (
	"net/http"

	"github.com/0x587/guardeye/api/internal/logic"
	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func WsListenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewWsListenLogic(r.Context(), svcCtx)
		err := l.WsListen(w, r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
