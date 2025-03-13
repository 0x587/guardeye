package logic

import (
	"context"
	"errors"

	"github.com/0x587/guardeye/link/internal/svc"
	"github.com/0x587/guardeye/link/link"

	"github.com/zeromicro/go-zero/core/logx"
)

type TypeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TypeListLogic {
	return &TypeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TypeListLogic) TypeList(in *link.TypeListReq) (*link.TypeListRsp, error) {
	rsp, err := agentConn.Call(l.ctx, in.GetCid(), &link.LinkCommandDownstream{
		PayloadRosList: &link.LinkCommandPayloadRosList{},
	})
	if err != nil {
		return nil, err
	}
	if !rsp.Ok {
		return nil, errors.New(rsp.ErrMsg)
	}
	return rsp.GetTypeListResult(), nil
}
