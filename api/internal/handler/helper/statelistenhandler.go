package helper

import (
	"net/http"

	"github.com/0x587/guardeye/api/internal/logic/helper"
	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func StateListenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := helper.NewStateListenLogic(r.Context(), svcCtx)
		resp, err := l.StateListen()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
