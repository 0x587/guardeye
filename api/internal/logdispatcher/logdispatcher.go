package logdispatcher

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/0x587/guardeye/api/internal/types"
	"github.com/0x587/guardeye/api/internal/ws"
	"github.com/0x587/guardeye/common/tokv"
	"github.com/0x587/guardeye/report/report"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type IF interface {
	Handle(nid uuid.UUID, log *report.Log)
	SetListen(sid, nid uuid.UUID, keys []types.ProviderKeys)
}

func New(ws ws.IF) IF {
	return &impl{
		ws:    ws,
		query: make(map[uuid.UUID][]*Query),
	}
}

type impl struct {
	ws ws.IF
	// nid -> []*Query
	query map[uuid.UUID][]*Query
}

func (i *impl) SetListen(sid, nid uuid.UUID, keys []types.ProviderKeys) {
	for _, q := range i.query[nid] {
		if q.sid == sid {
			q.enable = false
		}
	}
	i.query[nid] = lo.Filter(i.query[nid], func(q *Query, _ int) bool {
		return q.enable
	})
	for _, pk := range keys {
		for _, k := range pk.Keys {
			i.query[nid] = append(i.query[nid], &Query{
				enable:       true,
				providerType: pk.Provider.Ptype,
				providerArgs: pk.Provider.Args,
				dataKey:      k,
				sid:          sid,
			})
		}
	}
	// ------------
	fmt.Printf("\n%#+v", i.query)
	// ------------
}

func (i *impl) Handle(nid uuid.UUID, log *report.Log) {
	var kv tokv.KV
	if log.GetType() == report.LogType_YAML {
		kv = tokv.YamlToKv(log.GetMessage())
	}
	if log.GetType() == report.LogType_JSON {
		kv = tokv.JsonToKv(log.GetMessage())
	}
	ms := make(map[uuid.UUID]*types.MetricRsp)
	for _, q := range i.query[nid] {
		if !q.matchProvider(log.GetProvider()) {
			continue
		}
		pk := types.ProviderKey{
			Provider: types.Provider{
				Ptype: log.GetProvider().GetType(),
				Args:  log.GetProvider().GetArgs(),
				Str:   fmt.Sprintf("%s(%s)", log.GetProvider().GetType(), strings.Join(log.GetProvider().GetArgs(), ",")),
			},
			Key: q.dataKey,
		}
		m := types.Metric{
			Name:  string(lo.Must(json.Marshal(pk))),
			Value: "",
		}
		switch log.GetType() {
		case report.LogType_TEXT:
			m.Value = log.GetMessage()
		case report.LogType_YAML:
			m.Value = kv[q.dataKey]
		case report.LogType_JSON:
			m.Value = kv[q.dataKey]
		}
		if m.Value == "" {
			continue
		}
		if ms[q.sid] == nil {
			ms[q.sid] = &types.MetricRsp{
				WsMsgBase: types.WsMsgBase{Cmd: "METRIC"},
			}
		}
		ms[q.sid].Metrics = append(ms[q.sid].Metrics, m)
	}
	for sid, rsp := range ms {
		bs := lo.Must(json.Marshal(rsp))
		i.ws.Send(sid, bs)
	}
}

type Query struct {
	enable       bool
	providerType string
	providerArgs []string
	dataKey      string
	sid          uuid.UUID
}

func (q Query) matchProvider(p *report.Provider) bool {
	if p.GetType() != q.providerType {
		return false
	}
	slices.Sort(q.providerArgs)
	args := p.GetArgs()
	slices.Sort(args)
	if !slices.Equal(q.providerArgs, args) {
		return false
	}
	return true
}
