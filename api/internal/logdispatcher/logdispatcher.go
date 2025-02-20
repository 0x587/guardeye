package logdispatcher

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/google/uuid"

	"github.com/0x587/guardeye/api/internal/types"
	"github.com/0x587/guardeye/api/internal/ws"
	"github.com/0x587/guardeye/report/report"
)

type IF interface {
	Handle(nid uuid.UUID, log *report.Log)
	SetListen(sid uuid.UUID, query []types.LogQuery)
}

func New(ws ws.IF) IF {
	return &impl{
		ws:    ws,
		query: make(map[uuid.UUID][]types.LogQuery),
	}
}

type impl struct {
	ws ws.IF
	// sid -> []*Query
	query map[uuid.UUID][]types.LogQuery
}

func (i *impl) SetListen(sid uuid.UUID, queries []types.LogQuery) {
	for _, q := range queries {
		if q.Provider.Str == "" {
			q.Provider.Str = fmt.Sprintf("%s(%s)", q.Provider.Ptype, strings.Join(q.Provider.Args, ","))
		}
	}
	i.query[sid] = queries
}

func (i *impl) Handle(nid uuid.UUID, log *report.Log) {
	//var kv tokv.KV
	//if log.GetType() == report.LogType_YAML {
	//	kv = tokv.YamlToKv(log.GetMessage())
	//}
	//if log.GetType() == report.LogType_JSON {
	//	kv = tokv.JsonToKv(log.GetMessage())
	//}
	logs := make(map[uuid.UUID][]string)
	logp := log.GetProvider()
	for sid, queries := range i.query {
		for _, query := range queries {
			slices.Sort(query.Provider.Args)
			slices.Sort(logp.GetArgs())
			if !(query.Provider.Ptype == logp.GetType() && slices.Equal(query.Provider.Args, logp.GetArgs())) {
				continue
			}
			logs[sid] = append(logs[sid], log.GetMessage())
		}
		//	TODO Send Metric
	}
	for sid, logs := range logs {
		rsp := types.WsLogRsp{
			WsMsgBase: types.WsMsgBase{Cmd: "LOG"},
			Logs:      logs,
		}
		bytes, err := json.Marshal(rsp)
		if err != nil {
			continue
		}
		i.ws.Send(sid, bytes)
	}

	//for _, q := range i.query[nid] {
	//	if !q.matchProvider(log.GetProvider()) {
	//		continue
	//	}
	//	pk := types.ProviderKey{
	//		Provider: types.Provider{
	//			Ptype: log.GetProvider().GetType(),
	//			Args:  log.GetProvider().GetArgs(),
	//			Str:   fmt.Sprintf("%s(%s)", log.GetProvider().GetType(), strings.Join(log.GetProvider().GetArgs(), ",")),
	//		},
	//		Key: q.dataKey,
	//	}
	//	m := types.Metric{
	//		Name:  string(lo.Must(json.Marshal(pk))),
	//		Value: "",
	//	}
	//	switch log.GetType() {
	//	case report.LogType_TEXT:
	//		m.Value = log.GetMessage()
	//	case report.LogType_YAML:
	//		m.Value = kv[q.dataKey]
	//	case report.LogType_JSON:
	//		m.Value = kv[q.dataKey]
	//	}
	//	if m.Value == "" {
	//		continue
	//	}
	//	if ms[q.sid] == nil {
	//		ms[q.sid] = &types.MetricRsp{
	//			WsMsgBase: types.WsMsgBase{Cmd: "METRIC"},
	//		}
	//	}
	//	ms[q.sid].Metrics = append(ms[q.sid].Metrics, m)
	//}
	//for sid, rsp := range ms {
	//	bs := lo.Must(json.Marshal(rsp))
	//	i.ws.Send(sid, bs)
	//}
}
