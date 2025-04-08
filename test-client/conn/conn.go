package conn

import (
	"context"
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/0x587/guardeye/report/report"
	"github.com/0x587/guardeye/report/reportclient"
	"github.com/0x587/guardeye/test-client/config"
	"github.com/0x587/guardeye/test-client/conn/downstream"
	"github.com/0x587/guardeye/test-client/feature"
	"github.com/0x587/guardeye/test-client/feature/featuredelay"
	"github.com/0x587/guardeye/test-client/provider"
	"github.com/0x587/guardeye/test-client/storage"
)

type IF interface {
	Loop(ctx context.Context)
}

func New(reportCli zrpc.Client, c config.Config, providers ...provider.IF) IF {

	res := &impl{
		providers:    providers,
		c:            c,
		reportCli:    reportclient.NewReport(reportCli),
		storage:      storage.New(),
		featureDelay: featuredelay.New(),
		logCh:        make(chan *reportclient.Log, 100),
	}
	return res
}

type impl struct {
	clientID     string
	c            config.Config
	storage      storage.IF
	desc         *reportclient.NodeDescription
	reportCli    reportclient.Report
	featureDelay feature.IF[
		*reportclient.FeatureTransDelayReq,
		*reportclient.FeatureTransDelayRsp,
	]
	providers  []provider.IF
	onceInit   sync.Once
	logCh      chan *reportclient.Log
	downstream downstream.IF
}

func (i *impl) Loop(ctx context.Context) {
	i.onceInit.Do(func() { logx.Must(i.init(ctx)) })
	go i.reportLoop(ctx)
	for _, p := range i.providers {
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case msg := <-p.Get():
					i.logCh <- &reportclient.Log{
						Message:       msg.Message,
						Type:          report.LogType_TEXT,
						Provider:      &msg.Provider,
						ReportAtMilli: uint64(time.Now().UnixMilli()),
					}
				}
			}
		}()
	}
}

func (i *impl) reportLoop(ctx context.Context) {
	for log := range i.logCh {
		err := i.doLogReport(ctx, []*reportclient.Log{log})
		if err != nil {
			logx.Errorf("report error: %v", err)
			if errors.Is(err, context.DeadlineExceeded) {
				go func() {
					i.logCh <- log
				}()
			}
		}
	}
	//for {
	//	logs, _, _, ok := lo.BufferWithTimeout(i.logCh, 10, time.Second)
	//	if !ok {
	//		continue
	//	}
	//	func() {
	//		println(2)
	//		err := i.doLogReport(ctx, logs)
	//		if err != nil {
	//			logx.Errorf("report error: %v", err)
	//			if errors.Is(err, context.DeadlineExceeded) {
	//				go func() {
	//					for _, log := range logs {
	//						i.logCh <- log
	//					}
	//				}()
	//			}
	//		}
	//	}()
	//}
}

func (i *impl) init(ctx context.Context) error {
	cid, err := i.storage.FetchOrSet("clientId", func() []byte {
		req := &reportclient.InitReq{
			NodeDescription: i.getNodeDesc(),
		}
		rsp, err := i.reportCli.Init(ctx, req)
		logx.Must(err)
		clientID := rsp.GetNodeInfo().GetClientId()
		return []byte(clientID)
	})
	if err != nil {
		return err
	}

	i.clientID = string(cid)

	i.downstream, err = downstream.New(lo.Must(uuid.ParseBytes(cid)), i.c)
	if err != nil {
		return err
	}

	logx.Infof("client id: %v", i.clientID)
	heartbeatTicker := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-heartbeatTicker.C:
				_, _ = i.reportCli.Heartbeat(ctx, &reportclient.HeartbeatReq{
					NodeInfo: &reportclient.NodeInfo{
						ClientId:        i.clientID,
						NodeDescription: i.getNodeDesc(),
					},
					SendAtNano: uint64(time.Now().UnixNano()),
				})
			}
		}
	}()
	return nil
}

func (i *impl) doLogReport(ctx context.Context, logs []*reportclient.Log) error {
	req := &reportclient.LogReportReq{
		NodeInfo: &reportclient.NodeInfo{
			ClientId:        i.clientID,
			NodeDescription: i.getNodeDesc(),
		},
		Logs: logs,
		Features: &reportclient.FeaturesReq{
			TransDelay: lo.Must(i.featureDelay.MakeReq()),
		},
	}
	rsp, err := i.reportCli.LogReport(ctx, req)
	if err != nil {
		return err
	}
	lo.Must0(i.featureDelay.HandleRsp(rsp.GetFeatures().GetTransDelay()))
	return nil
}

var bootAt = time.Now()

func (i *impl) getNodeDesc() *reportclient.NodeDescription {
	addr := lo.Must(net.InterfaceAddrs())
	interfaces := lo.Must(net.Interfaces())
	memInfo, _ := mem.VirtualMemory()
	partitions, _ := disk.Partitions(false)
	upTime := time.Now().Sub(bootAt)
	upTime = upTime.Truncate(time.Second)
	return &reportclient.NodeDescription{
		Os:        runtime.GOOS,
		OsVersion: runtime.GOARCH,
		Ips: lo.Map(addr, func(item net.Addr, _ int) string {
			return item.String()
		}),
		Macs: lo.Map(interfaces, func(item net.Interface, _ int) string {
			return item.HardwareAddr.String()
		}),
		Hostname: lo.Must(os.Hostname()),
		Cpu:      strconv.Itoa(runtime.NumCPU()),
		Memory:   strconv.FormatUint(memInfo.Total, 10),
		Disk: strconv.FormatUint(lo.SumBy(partitions, func(item disk.PartitionStat) uint64 {
			usage, _ := disk.Usage(item.Mountpoint)
			return usage.Total
		}), 10),
		Uptime: fmt.Sprintf("%s", upTime),
	}
}
