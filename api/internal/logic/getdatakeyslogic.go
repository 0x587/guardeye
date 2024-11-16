package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
	"github.com/0x587/guardeye/report/report"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDataKeysLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDataKeysLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDataKeysLogic {
	return &GetDataKeysLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDataKeysLogic) GetDataKeys(req *types.GetDataKeysReq) (resp *types.GetDataKeysRsp, err error) {
	keys, err := l.svcCtx.DataKeyRedisClient.GetKeys(l.ctx, &report.NodeInfo{ClientId: req.Id})
	if err != nil {
		return nil, err
	}
	ps := make(map[string][]string)
	for _, k := range keys {
		p := &types.Provider{
			Ptype: k.Provider.GetType(),
			Args:  k.Provider.GetArgs(),
			Str:   fmt.Sprintf("%s(%s)", k.Provider.GetType(), strings.Join(k.Provider.GetArgs(), ",")),
		}
		pb := lo.Must(json.Marshal(p))
		ps[string(pb)] = append(ps[string(pb)], k.Key)
	}
	rs := lo.MapToSlice(ps, func(k string, v []string) types.ProviderKeys {
		var p types.Provider
		logc.Must(json.Unmarshal([]byte(k), &p))
		slices.Sort(v)
		return types.ProviderKeys{
			Provider: p,
			Keys:     v,
		}
	})
	return &types.GetDataKeysRsp{
		Keys: rs,
	}, nil
}
