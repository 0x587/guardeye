package ws

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
)

type IF[T any] interface {
	ServeWs(w http.ResponseWriter, r *http.Request)
	ClientIDs() []uuid.UUID
	ClientCount() int
	Close()
	Broadcast(msg T) error
}

func New[T any]() IF[T] {
	return &impl[T]{
		hub: newHub(),
	}
}

type impl[T any] struct {
	hub *hub
}

func (i *impl[T]) Broadcast(msg T) error {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	i.hub.broadcast <- bytes
	return nil
}

func (i *impl[T]) Close() {
	for _, c := range i.hub.clients {
		c.close()
	}
	i.hub.clients = make(map[uuid.UUID]*client)
}

func (i *impl[T]) ClientIDs() []uuid.UUID {
	return lo.Keys(i.hub.clients)
}

func (i *impl[T]) ClientCount() int {
	return len(i.hub.clients)
}

// ServeWs handles websocket requests from the peer.
func (i *impl[T]) ServeWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logx.Error(err)
		return
	}

	c := newClient(i.hub, conn)
	c.run()
}
