package logic

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/0x587/guardeye/common/async"
	"github.com/0x587/guardeye/common/foxglovetopb"
	"github.com/0x587/guardeye/link/internal/svc"
	"github.com/0x587/guardeye/link/link"

	"github.com/zeromicro/go-zero/core/logx"
)

type TypeGenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTypeGenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TypeGenLogic {
	return &TypeGenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TypeGenLogic) TypeGen(in *link.TypeGenReq) (*link.TypeGenRsp, error) {
	listRsp, err := NewTypeListLogic(l.ctx, l.svcCtx).TypeList(&link.TypeListReq{Cid: in.GetCid()})
	if err != nil {
		return nil, err
	}
	var tasks []func() error
	var mutex sync.Mutex
	env := foxglovetopb.New()
	for _, topic := range listRsp.GetMessages() {
		tasks = append(tasks, func() error {
			rsp, err := agentConn.Call(l.ctx, in.GetCid(), &link.LinkCommandDownstream{
				PayloadRosType: &link.LinkCommandPayloadRosType{
					Type:     link.RosTypeGenType_MESSAGE,
					RosTopic: topic,
				},
			})
			if err != nil {
				return err
			}
			if !rsp.Ok {
				return errors.New(fmt.Sprintf("topic: %s %s", topic, rsp.ErrMsg))
			}
			mutex.Lock()
			defer mutex.Unlock()
			err = env.ParseMsg(topic, rsp.GetTypeGenResult().GetName(), rsp.GetTypeGenResult().GetReq())
			if err != nil {
				return err
			}
			return nil
		})
	}
	for _, topic := range listRsp.GetServices() {
		tasks = append(tasks, func() error {
			rsp, err := agentConn.Call(l.ctx, in.GetCid(), &link.LinkCommandDownstream{
				PayloadRosType: &link.LinkCommandPayloadRosType{
					Type:     link.RosTypeGenType_SERVICE,
					RosTopic: topic,
				},
			})
			if err != nil {
				return err
			}
			if !rsp.Ok {
				return errors.New(fmt.Sprintf("topic: %s %s", topic, rsp.ErrMsg))
			}
			mutex.Lock()
			defer mutex.Unlock()
			err = env.ParseSrv(topic, rsp.GetTypeGenResult().GetName(), rsp.GetTypeGenResult().GetReq(), rsp.GetTypeGenResult().GetRsp())
			if err != nil {
				return err
			}
			return nil
		})
	}
	if err := async.GoAndWait(l.ctx, tasks...); err != nil {
		return nil, err
	}
	return &link.TypeGenRsp{Pb: env.Output()}, nil
}
