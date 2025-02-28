package es

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/samber/lo"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskStatusLogic {
	return &TaskStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskStatusLogic) TaskStatus(req *types.EsTaskStatusReq) (resp *types.EsTaskStatusRsp, err error) {
	redisKey := fmt.Sprintf("report_task_%s", req.TaskId)
	m, err := l.svcCtx.Redis.HgetallCtx(l.ctx, redisKey)
	if err != nil {
		return nil, err
	}
	resp = &types.EsTaskStatusRsp{}
	resp.State = m["status"]
	if m["process"] != "" {
		resp.Process = lo.Must(strconv.Atoi(m["process"]))
	}
	if m["total"] != "" {
		resp.Total = lo.Must(strconv.Atoi(m["total"]))
	}
	resp.Done = m["done"] == "1"
	if resp.Done {
		reqParams := make(url.Values)
		object, err := l.svcCtx.Minio.PresignedGetObject(l.ctx, "export", fmt.Sprintf("%s.csv", req.TaskId), time.Hour, reqParams)
		if err != nil {
			return nil, err
		}
		resp.Link = object.String()
	}
	return resp, nil
}
