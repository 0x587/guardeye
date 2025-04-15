package httpcb

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/common/ent"
	"github.com/0x587/guardeye/common/ent/callback"
)

type IF interface {
	Online(ctx context.Context, cid uuid.UUID)
	Offline(ctx context.Context, cid uuid.UUID)
	Data(ctx context.Context, cid uuid.UUID, topic string, payload string)
}

func New(client *ent.Client) IF {
	return &impl{
		db: client,
		c:  cache.New(time.Second*10, time.Minute),
	}
}

type impl struct {
	db *ent.Client
	c  *cache.Cache
}

func (i *impl) Online(ctx context.Context, cid uuid.UUID) {
	cfg, err := i.getConfig(ctx, cid)
	if err != nil {
		logx.Error(err)
		return
	}
	cfgItem := cfg.Basic.Online
	if !cfgItem.Enabled {
		return
	}
	u, err := url.Parse(cfgItem.Url)
	if err != nil {
		logx.Error(err)
		return
	}
	req := &http.Request{
		URL:    u,
		Method: cfgItem.Method,
		Header: toHeader(cfgItem.Headers),
	}
	if _, err := http.DefaultClient.Do(req); err != nil {
		logx.Error(err)
		return
	}
}

func (i *impl) Offline(ctx context.Context, cid uuid.UUID) {
	cfg, err := i.getConfig(ctx, cid)
	if err != nil {
		logx.Error(err)
		return
	}
	cfgItem := cfg.Basic.Offline
	if !cfgItem.Enabled {
		return
	}
	u, err := url.Parse(cfgItem.Url)
	if err != nil {
		logx.Error(err)
		return
	}
	req := &http.Request{
		URL:    u,
		Method: cfgItem.Method,
		Header: toHeader(cfgItem.Headers),
	}
	if _, err := http.DefaultClient.Do(req); err != nil {
		logx.Error(err)
		return
	}
}

func (i *impl) Data(ctx context.Context, cid uuid.UUID, topic string, payload string) {
	cfg, err := i.getConfig(ctx, cid)
	if err != nil {
		logx.Error(err)
		return
	}
	for _, cfgItem := range cfg.Data {
		if !topicMatch(cfgItem.TopicPattern, topic) {

		}
		if !cfgItem.Enabled {
			continue
		}
		u, err := url.Parse(cfgItem.Url)
		if err != nil {
			logx.Error(err)
			continue
		}
		req := &http.Request{
			URL:    u,
			Method: cfgItem.Method,
			Header: toHeader(cfgItem.Headers),
			Body:   io.NopCloser(strings.NewReader(mixedTopicPayload(topic, payload))),
		}
		if _, err := http.DefaultClient.Do(req); err != nil {
			logx.Error(err)
			continue
		}
	}
}

func (i *impl) getConfig(ctx context.Context, cid uuid.UUID) (*CallbackConfig, error) {
	res, exist := i.c.Get(cid.String())
	if exist {
		i.c.Set(cid.String(), res, cache.DefaultExpiration)
		return res.(*CallbackConfig), nil
	}
	cfg, err := i._getConfig(ctx, cid)
	if err != nil {
		return nil, err
	}
	i.c.Set(cid.String(), cfg, cache.DefaultExpiration)
	return cfg, err
}

func (i *impl) _getConfig(ctx context.Context, cid uuid.UUID) (*CallbackConfig, error) {
	first, err := i.db.Callback.Query().Where(callback.ClientIDEQ(cid)).First(ctx)
	if err != nil {
		return nil, err
	}
	res := &CallbackConfig{}
	if err := json.Unmarshal([]byte(first.Cfg), res); err != nil {
		return nil, err
	}
	return res, nil
}

func toHeader(h map[string]string) http.Header {
	return lo.MapValues(h, func(v, _ string) []string {
		return []string{v}
	})
}

// topicMatch 判断实际 topic 是否匹配 pattern（支持 + 和 # 通配符）
func topicMatch(pattern, topic string) bool {
	pParts := strings.Split(pattern, "/")
	tParts := strings.Split(topic, "/")

	for i := 0; i < len(pParts); i++ {
		if i >= len(tParts) {
			// 只有 # 可以匹配多余部分
			return pParts[i] == "#"
		}

		switch pParts[i] {
		case "#":
			// # 必须是最后一段
			return i == len(pParts)-1
		case "+":
			// + 匹配当前层级，继续匹配下一层
			continue
		default:
			// 精确匹配
			if pParts[i] != tParts[i] {
				return false
			}
		}
	}

	// 如果 topic 还有多余部分，且 pattern 没有 #，则不匹配
	return len(tParts) == len(pParts)
}

func mixedTopicPayload(topic, payload string) string {
	var data any
	err := json.Unmarshal([]byte(payload), &data)
	if err != nil {
		logx.Error(err)
		return payload
	}
	res, err := json.Marshal(map[string]any{
		"topic":   topic,
		"payload": data,
	})
	if err != nil {
		logx.Error(err)
		return payload
	}
	return string(res)
}
