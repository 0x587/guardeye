package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/0x587/guardeye/wsapi/internal/logic"
	"github.com/0x587/guardeye/wsapi/internal/svc"
)

func WsapiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewWsapiLogic(r.Context(), svcCtx)
		err := l.Wsapi(w, r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
