package helper

import (
	"context"
	"strings"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StateListenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStateListenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StateListenLogic {
	return &StateListenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StateListenLogic) StateListen() (resp *types.ListenState, err error) {
	state := make(map[string]map[string]bool)
	listenState.Range(func(key, _ any) bool {
		parts := strings.SplitN(key.(string), "_", 2)
		cid, topic := parts[0], parts[1]
		if state[cid] == nil {
			state[cid] = make(map[string]bool)
		}
		state[cid][topic] = true
		return true
	})
	resp = &types.ListenState{
		State: state,
	}
	return
}
