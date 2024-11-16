package logic

import (
	"context"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
	"github.com/google/uuid"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetListenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetListenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetListenLogic {
	return &SetListenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetListenLogic) SetListen(req *types.SetListenReq) error {
	// todo: add your logic here and delete this line
	sid, err := uuid.Parse(req.SessionId)
	if err != nil {
		return err
	}
	nid, err := uuid.Parse(req.NodeId)
	if err != nil {
		return err
	}
	l.svcCtx.LogDispatcher.SetListen(sid, nid, req.Keys)
	return nil
}
