package logic

import (
	"context"

	"github.com/google/uuid"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/link/internal/svc"
	"github.com/0x587/guardeye/link/link"
)

type LinkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LinkLogic {
	return &LinkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LinkLogic) Link(stream link.Link_LinkServer) error {
	var cid string
	for {
		recv, err := stream.Recv()
		if err != nil {
			return err
		}
		if recv.GetCid() == "" {
			continue
		}
		cid = recv.GetCid()
		break
	}

	linkId := uuid.New()
	logx.Infof("accept link %s %s", cid, linkId.String())
	linkRpcPoll.Accept(cid, stream)
	logx.Infof("kill link %s %s", cid, linkId.String())

	return nil
}
