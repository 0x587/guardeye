package ws

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
)

type WS interface {
	Broadcast(msg []byte)
	Send(cid uuid.UUID, msg []byte)
}

type IF interface {
	WS
	ServeWs(w http.ResponseWriter, r *http.Request)
	ClientIDs() []uuid.UUID
	ClientCount() int
	Close()
}

func New() IF {
	return &impl{
		hub: newHub(),
	}
}

type impl struct {
	*hub
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
