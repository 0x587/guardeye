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

type SetCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetCallbackLogic {
	return &SetCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetCallbackLogic) SetCallback(req *types.SetCallbackReq) (resp *types.CallbackConfig, err error) {
	cid, err := uuid.Parse(req.Cid)
	if err != nil {
		return nil, err
	}
	first, err := l.svcCtx.Db.Callback.Query().Where(callback.ClientIDEQ(cid)).First(l.ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}
		first, err = l.svcCtx.Db.Callback.Create().SetClientID(cid).Save(l.ctx)
		if err != nil {
			return nil, err
		}
	}
	cfgBytes, err := json.Marshal(req.Config)
	if err != nil {
		return nil, err
	}
	_, err = first.Update().SetCfg(string(cfgBytes)).Save(l.ctx)
	if err != nil {
		return nil, err
	}
	return &req.Config, nil
}
