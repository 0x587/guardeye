package logdispatcher

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/0x587/guardeye/api/internal/types"
	"github.com/0x587/guardeye/common/ws"
	"github.com/0x587/guardeye/report/report"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"gopkg.in/yaml.v3"
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
	for _, q := range i.query[nid] {
		if !q.matchProvider(log.GetProvider()) {
			continue
		}
		pk := types.ProviderKey{
			Provider: types.Provider{
				Ptype: log.GetProvider().GetType(),
				Args:  log.GetProvider().GetArgs(),
			},
			Key: q.dataKey,
		}
		m := ws.MetricRsp{
			MsgBase: ws.MsgBase{Cmd: "METRIC"},
			Name:    string(lo.Must(json.Marshal(pk))),
			Value:   "",
		}
		switch log.GetType() {
		case report.LogType_TEXT:
			m.Value = log.GetMessage()
		case report.LogType_YAML:
			var obj map[string]interface{}
			err := yaml.Unmarshal([]byte(log.GetMessage()), &obj)
			if err != nil {
				continue
			}
			m.Value = getValueFromPath(obj, strings.Split(q.dataKey, "."))
		}
		bs := lo.Must(json.Marshal(m))
		i.ws.Send(q.sid, bs)
	}
}

func getValueFromPath(obj interface{}, path []string) string {
	if len(path) == 0 {
		return fmt.Sprintf("%v", obj)
	}
	switch v := obj.(type) {
	case map[string]interface{}:
		return getValueFromPath(v[path[0]], path[1:])
	case []interface{}:
		i, err := strconv.ParseInt(path[0], 10, 64)
		if err != nil {
			return ""
		}
		if int(i) >= len(v) {
			return ""
		}
		return getValueFromPath(v[i], path[1:])
	default:
		return ""
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
