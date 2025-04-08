package logtometric

//import (
//	"net/http"
//
//	"github.com/0x587/guardeye/api/internal/logic/logtometric"
//	"github.com/0x587/guardeye/api/internal/svc"
//	"github.com/0x587/guardeye/api/internal/types"
//	"github.com/zeromicro/go-zero/rest/httpx"
//)
//
//func GetQueriesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var req types.GetQueriesReq
//		if err := httpx.Parse(r, &req); err != nil {
//			httpx.ErrorCtx(r.Context(), w, err)
//			return
//		}
//
//		l := logtometric.NewGetQueriesLogic(r.Context(), svcCtx)
//		resp, err := l.GetQueries(&req)
//		if err != nil {
//			httpx.ErrorCtx(r.Context(), w, err)
//		} else {
//			httpx.OkJsonCtx(r.Context(), w, resp)
//		}
//	}
//}
