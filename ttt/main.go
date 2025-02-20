package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"

	report2 "github.com/0x587/guardeye/report/report"
	"github.com/0x587/guardeye/report/reportclient"
)

func main() {
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		client, err := zrpc.NewClient(zrpc.RpcClientConf{
			Etcd: discov.EtcdConf{},
			Endpoints: []string{
				"10.0.4.112:38080",
			},
		})
		if err != nil {
			panic(err)
		}
		payload := map[string]any{
			"temperature": 100,
		}
		if time.Now().Second() < 20 {
			payload["temperature"] = 75
		}
		marshal, err := json.Marshal(payload)
		if err != nil {
			panic(err)
		}
		report, err := reportclient.NewReport(client).LogReport(
			context.Background(), &reportclient.LogReportReq{
				NodeInfo: &reportclient.NodeInfo{
					ClientId:        "3b4b8a63-96af-4a57-810e-85b43deb650c",
					NodeDescription: nil,
				},
				Logs: []*reportclient.Log{
					{
						Message: string(marshal),
						Provider: &reportclient.Provider{
							Type: "mqtt",
							Args: nil,
						},
						Type:          report2.LogType_JSON,
						ReportAtMilli: uint64(time.Now().UnixMilli()),
					},
				},
			})
		if err != nil {
			panic(err)
		}
		fmt.Print(report)
	}
}
