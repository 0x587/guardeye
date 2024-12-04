package reporter

import (
	"bytes"
	"context"
	"io"
	"net"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/report/report"
	"github.com/0x587/guardeye/report/reportclient"
	"github.com/0x587/guardeye/test-client/feature"
	"github.com/0x587/guardeye/test-client/feature/featuredelay"
	"github.com/0x587/guardeye/test-client/provider"
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
		err := post(i, ctx, "/init", req, rsp)
		logx.Must(err)
		clientID := rsp.GetNodeInfo().GetClientId()
		return []byte(clientID)
	})
	if err != nil {
		return err
	}
	i.clientID = string(cid)
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
	err := post(i, ctx, "/log_report", req, rsp)
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

func post[ReqT, RspT proto.Message](i *impl, ctx context.Context, url string, req ReqT, rsp RspT) error {
	m := jsonpb.Marshaler{}
	s, err := m.MarshalToString(req)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	request, err := http.NewRequestWithContext(
		ctx,
		"POST",
		i.reportBaseurl+url,
		bytes.NewBuffer([]byte(s)),
	)
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.Errorf("report http error: %v", response.Status)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = jsonpb.UnmarshalString(string(body), rsp)
	if err != nil {
		return err
	}
	return nil
}
