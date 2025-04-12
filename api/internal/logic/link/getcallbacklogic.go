package link

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
	"github.com/0x587/guardeye/common/ent"
	"github.com/0x587/guardeye/common/ent/callback"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCallbackLogic {
	return &GetCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var defaultCfg = &types.CallbackConfig{
	Basic: types.BasicCallbackConfig{
		Online: types.CallbackItem{
			Headers: make(map[string]string),
		},
		Offline: types.CallbackItem{
			Headers: make(map[string]string),
		},
	},
	Data: make([]types.DataCallbackItem, 0),
}

func (l *GetCallbackLogic) GetCallback(req *types.GetCallbackReq) (resp *types.CallbackConfig, err error) {
	cid, err := uuid.Parse(req.Cid)
	if err != nil {
		return nil, err
	}
	first, err := l.svcCtx.Db.Callback.Query().Where(callback.ClientIDEQ(cid)).First(l.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return defaultCfg, nil
		}
		return nil, err
	}
	resp = &types.CallbackConfig{}
	if err := json.Unmarshal([]byte(first.Cfg), resp); err != nil {
		return nil, err
	}
	return
}
