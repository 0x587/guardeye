package ws

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
)

type IF interface {
	WS
	ServeWs(w http.ResponseWriter, r *http.Request)
	ClientIDs() []uuid.UUID
	ClientCount() int
	Close()
	AddHandler(handler func(IF, uuid.UUID, []byte))
}

type WS interface {
	Broadcast(msg []byte)
	Send(cid uuid.UUID, msg []byte)
}

func New() IF {
	res := &impl{
		hub: newHub(),
	}
	go func() {
		for m := range res.hub.received {
			for _, h := range res.handlers {
				h(res, m.cid, m.msg)
			}
		}
	}()
	return res
}

type impl struct {
	*hub
	handlers []func(IF, uuid.UUID, []byte)
}

func (i *impl) AddHandler(handler func(IF, uuid.UUID, []byte)) {
	i.handlers = append(i.handlers, handler)
}

func (i *impl) Close() {
	for _, c := range i.hub.clients {
		c.close()
	}
	i.hub.clients = make(map[uuid.UUID]*client)
}

func (i *impl) ClientIDs() []uuid.UUID {
	return lo.Keys(i.hub.clients)
}

func (i *impl) ClientCount() int {
	return len(i.hub.clients)
}

// ServeWs handles websocket requests from the peer.
func (i *impl) ServeWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logx.Error(err)
		return
	}

	c := newClient(i.hub, conn)
	c.run()
}
