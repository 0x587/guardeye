package logic

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/common/ent/agent"
	"github.com/0x587/guardeye/common/rediskey"
	"github.com/0x587/guardeye/report/internal/svc"
	"github.com/0x587/guardeye/report/report"
)

type HeartbeatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

const stateCheckerTimeout = time.Second * 10

var stateCheckerCache = cache.New(time.Minute, time.Minute*2)

func NewHeartbeatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HeartbeatLogic {
	return &HeartbeatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HeartbeatLogic) Heartbeat(in *report.HeartbeatReq) (*report.Empty, error) {
	l.stateCheck(in)
	nodeInfo := in.GetNodeInfo()
	cid, err := uuid.Parse(nodeInfo.GetClientId())
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.Db.Agent.Update().
		SetIps(nodeInfo.GetNodeDescription().GetIps()).
		SetMacs(nodeInfo.GetNodeDescription().GetMacs()).
		SetOs(nodeInfo.GetNodeDescription().GetOs()).
		SetOsVersion(nodeInfo.GetNodeDescription().GetOsVersion()).
		SetHostname(nodeInfo.GetNodeDescription().GetHostname()).
		SetCPU(nodeInfo.GetNodeDescription().GetCpu()).
		SetMemory(nodeInfo.GetNodeDescription().GetMemory()).
		SetDisk(nodeInfo.GetNodeDescription().GetDisk()).
		SetUptime(nodeInfo.GetNodeDescription().GetUptime()).
		Where(agent.ClientIDEQ(cid)).
		Save(l.ctx)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.RedisClient.SetCtx(l.ctx, rediskey.SeeAtKey(nodeInfo), strconv.FormatInt(time.Now().UnixNano(), 10))
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.RedisClient.SetCtx(l.ctx, rediskey.LatencyKey(nodeInfo),
		strconv.FormatUint(uint64(in.GetLatency()), 10))
	if err != nil {
		return nil, err
	}
	return &report.Empty{}, nil
}

func (l *HeartbeatLogic) stateCheck(in *report.HeartbeatReq) {
	nodeInfo := in.GetNodeInfo()
	state, err := l.svcCtx.RedisClient.Get(rediskey.StateKey(nodeInfo))
	if err != nil {
		logx.Error(err)
		return
	}
	cid, _ := uuid.Parse(nodeInfo.GetClientId())
	if state != "online" {
		_ = l.svcCtx.RedisClient.Set(rediskey.StateKey(nodeInfo), "online")
		l.svcCtx.HttpCallback.Online(context.Background(), cid)
	}
	var c *stateChecker
	cc, exits := stateCheckerCache.Get(nodeInfo.GetClientId())
	if !exits {
		c = newStateChecker(stateCheckerTimeout, func() {
			_ = l.svcCtx.RedisClient.Set(rediskey.StateKey(nodeInfo), "offline")
			l.svcCtx.HttpCallback.Offline(context.Background(), cid)
		})
		stateCheckerCache.Set(nodeInfo.GetClientId(), c, cache.DefaultExpiration)
	}
	c, ok := cc.(*stateChecker)
	if ok {
		c.Heartbeat()
	}
}

type stateChecker struct {
	mu       sync.Mutex
	duration time.Duration
	ticker   *time.Ticker
	f        func()
	lastSee  time.Time
}

func newStateChecker(duration time.Duration, f func()) *stateChecker {
	return &stateChecker{
		duration: duration,
		f:        f,
	}
}

func (s *stateChecker) Heartbeat() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lastSee = time.Now()
	go func() {
		time.Sleep(s.duration)
		s.mu.Lock()
		defer s.mu.Unlock()
		if time.Since(s.lastSee) >= s.duration {
			s.f()
		}
	}()
}
