package logsubscribe

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/jsonpb"
	"github.com/samber/lo"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/common/cache"
	"github.com/0x587/guardeye/common/ent"
	"github.com/0x587/guardeye/common/ent/subscribe"
	"github.com/0x587/guardeye/common/gql"
	"github.com/0x587/guardeye/common/lifepool"
	"github.com/0x587/guardeye/report/internal/svc"
	"github.com/0x587/guardeye/report/report"
)

func New(ctx context.Context, svcCtx *svc.ServiceContext) kq.ConsumeHandler {
	return &impl{
		svcCtx: svcCtx,
		mqttClients: lifepool.NewLifePool(time.Minute,
			func(url string) mqtt.Client {
				client := mqtt.NewClient(mqtt.NewClientOptions().AddBroker(url))
				if token := client.Connect(); token.Wait() && token.Error() != nil {
					logx.Error(token.Error())
					return nil
				}
				return client
			},
			func(client mqtt.Client) {
				client.Disconnect(250)
			}),
	}
}

type impl struct {
	sync.Once
	svcCtx               *svc.ServiceContext
	fetchSubscribeConfig func(ctx context.Context) ([]*ent.Subscribe, error)
	mqttClients          *lifepool.LifePool[mqtt.Client, string]
}

var ErrSkip = errors.New("skip")

func (i *impl) Consume(ctx context.Context, key, value string) error {
	i.Do(func() {
		i.fetchSubscribeConfig = cache.TimerCacheFunc(i._fetchSubscribeConfig, time.Second)
	})
	config, err := i.fetchSubscribeConfig(ctx)
	if err != nil {
		return err
	}
	for _, s := range config {
		if !s.Enable {
			continue
		}
		res, err := execQuery(s.Query, value)
		if err != nil {
			if errors.Is(err, ErrSkip) {
				return nil
			}
			logx.Error("error on subscribe: ", err)
			return nil
		}
		i.outputValue(ctx, s, res)
	}
	return nil
}

func execQuery(query, value string) (*SubscribeResult, error) {
	tree, errs := gql.ParseQuery(query)
	if len(errs) > 0 {
		return nil, errors.Join(lo.Map(errs, func(e string, _ int) error {
			return errors.New(e)
		})...)
	}
	unmarshaler := jsonpb.Unmarshaler{}
	mqLog := &report.MQLog{}
	if err := unmarshaler.Unmarshal(strings.NewReader(value), mqLog); err != nil {
		return nil, err
	}

	schedule := gql.ScheduleQuery(tree)
	for _, s := range schedule.SourceWhere {
		if !s.Node.Any && s.Node.Nid != mqLog.NodeInfo.GetClientId() {
			return nil, ErrSkip
		}
		// TODO Provider
		for _, needKey := range s.NeedKeys {
			if !strings.Contains(mqLog.Log.Message, needKey) {
				return nil, ErrSkip
			}
		}
	}
	res := make(map[string]any)
	for _, r := range schedule.Result {
		v, err := r.Value.Vf(gql.NewInjector(mqLog.Log.Message))
		if err != nil {
			return nil, err
		}
		res[r.Alias] = v
	}
	return &SubscribeResult{
		Timestamp: mqLog.GetTimestamp(),
		ClientID:  mqLog.GetNodeInfo().GetClientId(),
		Provider:  mqLog.GetLog().GetProvider(),
		Data:      res,
	}, nil
}

type SubscribeResult struct {
	Timestamp string           `json:"timestamp"`
	ClientID  string           `json:"client_id"`
	Provider  *report.Provider `json:"provider"`
	Data      map[string]any   `json:"data"`
}

func (i *impl) outputValue(ctx context.Context, s *ent.Subscribe, res *SubscribeResult) {
	jsonData, err := json.Marshal(res)
	if err != nil {
		logx.Error(err)
		return
	}
	switch s.Type {
	case subscribe.TypeWebHook:
		_, err := http.Post(s.WebHookURL,
			"application/json",
			bytes.NewReader(jsonData))
		if err != nil {
			logx.Error(err)
			return
		}
	case subscribe.TypeMqttPush:
		client := i.mqttClients.Get(s.MqttPushURL)
		token := client.Publish(s.MqttPushTopic, 0, false, jsonData)
		if token.Wait() && token.Error() != nil {
			logx.Error(token.Error())
			return
		}
	case subscribe.TypeWebSocket:
		err := i.svcCtx.WsSubscribePusherClient.KPush(ctx, s.WebSocketKey, string(jsonData))
		if err != nil {
			logx.Error(err)
			return
		}
	}
}

func (i *impl) _fetchSubscribeConfig(ctx context.Context) ([]*ent.Subscribe, error) {
	subscribes, err := i.svcCtx.Db.Subscribe.
		Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return subscribes, nil
}
