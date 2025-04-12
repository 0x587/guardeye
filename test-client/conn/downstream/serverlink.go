package downstream

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/syncx"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/0x587/guardeye/link/link"
	"github.com/0x587/guardeye/link/linkclient"
)

type serverLink interface {
	Init(callback) error
	Close()
}
type callback func(req *linkclient.LinkCommandDownstream) (*linkclient.LinkCommandUpstream, error)

type mqttServerLink struct {
	cid uuid.UUID
	c   mqtt.Client
}

func newMqttServerLink(cid uuid.UUID) serverLink {
	opts := mqtt.NewClientOptions().AddBroker("tcp://emqxtcp.guardeye.shawnsiu.site:58701").
		SetClientID(fmt.Sprintf("agent_%s", cid.String())).
		SetKeepAlive(5 * time.Second).
		SetPingTimeout(1 * time.Second).
		SetAutoReconnect(true)

	c := mqtt.NewClient(opts)

	return &mqttServerLink{
		c:   c,
		cid: cid,
	}
}

func (l *mqttServerLink) Init(cb callback) error {
	if token := l.c.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	if token := l.c.Subscribe(fmt.Sprintf("command/req/%s", l.cid.String()), 2,
		func(client mqtt.Client, message mqtt.Message) {
			req := &linkclient.LinkCommandDownstream{}
			if err := json.Unmarshal(message.Payload(), req); err != nil {
				logx.Error(err)
				return
			}
			rsp, err := cb(req)
			if err != nil {
				logx.Error(err)
			}
			rspBuf, err := json.Marshal(rsp)
			if err != nil {
				logx.Error(err)
				return
			}
			token := client.Publish(fmt.Sprintf("command/rsp/%s", l.cid.String()), 2, false, rspBuf)
			if token.Wait(); token.Error() != nil {
				logx.Error(token.Error())
			}
		}); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (l *mqttServerLink) Close() {
	l.c.Disconnect(250)
}

type rpcServerLink struct {
	cli linkclient.Link
	cid uuid.UUID
}

func newRpcServerLink(cid uuid.UUID, linkEndpoint string) (serverLink, error) {
	flag.Parse()
	client, err := zrpc.NewClient(zrpc.RpcClientConf{
		Target: linkEndpoint,
	})
	if err != nil {
		return nil, err
	}
	cli := linkclient.NewLink(client)
	return &rpcServerLink{
		cli: cli,
		cid: cid,
	}, nil
}

func (l *rpcServerLink) Init(cb callback) error {
	limit := syncx.NewLimit(3)
	syncx.NewLockedCalls()
	c := cache.New(time.Minute, time.Minute*2)
	ticker := time.NewTicker(time.Second)
	makeLink := func() (link.Link_LinkClient, error) {
		// 创建连接
		link_, err := l.cli.Link(context.Background())
		if err != nil {
			return nil, err
		}
		err = link_.Send(&linkclient.LinkCommandUpstream{
			Cid: l.cid.String(),
		})
		if err != nil {
			return nil, err
		}
		return link_, err
	}
	handleLink := func(link_ link.Link_LinkClient) error {
		go func() {
			ticker := time.NewTicker(time.Second)
			// 心跳
			for {
				select {
				case <-ticker.C:
					err := link_.Send(&linkclient.LinkCommandUpstream{})
					if err != nil {
						return
					}
				}
			}
		}()
		for {
			recv, err := link_.Recv()
			if err != nil {
				return err
			}

			_, ok := c.Get(recv.GetId())
			if ok {
				continue
			}
			c.Set(recv.GetId(), true, cache.DefaultExpiration)
			rsp, err := cb(recv)
			if err != nil {
				return err
			}
			err = link_.Send(rsp)
			if err != nil {
				return err
			}
		}
	}

	go func() {
		//持续创建连接
		for {
			select {
			case <-ticker.C:
				if ok := limit.TryBorrow(); !ok {
					continue
				}
				go func() {
					err := func() error {
						// 创建连接
						logx.Info("make link")
						link_, err := makeLink()
						if err != nil {
							return err
						}
						err = handleLink(link_)
						if err != nil {
							return err
						}
						return nil
					}()
					if err != nil {
						logx.Error("link fail ", err)
						_ = limit.Return()
					}
				}()
			}
		}
	}()

	return nil
}

func (l *rpcServerLink) Close() {
	return
}
