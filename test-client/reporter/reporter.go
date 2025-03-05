package reporter

import (
	"context"
	"net"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/report/report"
	"github.com/0x587/guardeye/report/reportclient"
	"github.com/0x587/guardeye/test-client/feature"
	"github.com/0x587/guardeye/test-client/feature/featuredelay"
	"github.com/0x587/guardeye/test-client/http"
	"github.com/0x587/guardeye/test-client/provider"
	"github.com/0x587/guardeye/test-client/reporter/downstream"
	"github.com/0x587/guardeye/test-client/storage"
)

type IF interface {
	Loop(ctx context.Context)
}

func New(reportBaseurl string, providers ...provider.IF) IF {
	res := &impl{
		providers: providers,
		storage:   storage.New(),
		//reportClient: reportclient.NewReport(cli),
		reportBaseurl: reportBaseurl,
		featureDelay:  featuredelay.New(),
		logCh:         make(chan *reportclient.Log, 100),
	}
	return res
}

type impl struct {
	clientID string
	storage  storage.IF
	desc     *reportclient.NodeDescription
	//reportClient reportclient.Report
	reportBaseurl string
	featureDelay  feature.IF[
		*reportclient.FeatureTransDelayReq,
		*reportclient.FeatureTransDelayRsp,
	]
	providers []provider.IF
	onceInit  sync.Once
	logCh     chan *reportclient.Log
	mqttCli   downstream.MqttCli
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
	for {
		logs, _, _, ok := lo.BufferWithTimeout(i.logCh, 10, time.Second)
		if !ok {
			continue
		}
		func() {
			err := i.doLogReport(ctx, logs)
			if err != nil {
				logx.Errorf("report error: %v", err)
				if errors.Is(err, context.DeadlineExceeded) {
					go func() {
						for _, log := range logs {
							i.logCh <- log
						}
					}()
				}
			}
		}()
	}
}

func (i *impl) init(ctx context.Context) error {
	cid, err := i.storage.FetchOrSet("clientId", func() []byte {
		req := &reportclient.InitReq{
			NodeDescription: i.getNodeDesc(),
		}
		rsp := &reportclient.InitRsp{}
		err := http.Post(ctx, i.reportBaseurl+"/init", req, rsp)
		logx.Must(err)
		clientID := rsp.GetNodeInfo().GetClientId()
		return []byte(clientID)
	})
	if err != nil {
		return err
	}
	i.clientID = string(cid)
	i.mqttCli, err = downstream.NewMqtt(lo.Must(uuid.ParseBytes(cid)))
	if err != nil {
		return err
	}
	logx.Infof("client id: %v", i.clientID)
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
	rsp := &reportclient.LogReportRsp{}
	err := http.Post(ctx, i.reportBaseurl+"/log_report", req, rsp)
	if err != nil {
		return err
	}
	lo.Must0(i.featureDelay.HandleRsp(rsp.GetFeatures().GetTransDelay()))
	return nil
}

func (i *impl) getNodeDesc() *reportclient.NodeDescription {
	addr := lo.Must(net.InterfaceAddrs())
	interfaces := lo.Must(net.Interfaces())
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
	}
}
