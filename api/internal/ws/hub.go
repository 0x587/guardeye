package ws

import (
	"encoding/json"

	"github.com/0x587/guardeye/api/internal/types"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type hub struct {
	// Registered clients.
	clients map[uuid.UUID]*client
	// Register requests from the clients.
	register chan *client
	// Unregister requests from clients.
	unregister chan *client
}

func newHub() *hub {
	res := &hub{
		register:   make(chan *client),
		unregister: make(chan *client),
		clients:    make(map[uuid.UUID]*client),
	}
	go res.run()
	return res
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			h.clients[c.id] = c
			m := types.WsInitRsp{
				WsMsgBase: types.WsMsgBase{Cmd: "INIT"},
				SessionId: c.id.String(),
			}
			bs := lo.Must(json.Marshal(m))
			h.Send(c.id, bs)
		case c := <-h.unregister:
			if _, ok := h.clients[c.id]; ok {
				delete(h.clients, c.id)
				close(c.send)
			}
		}
	}
}

func (h *hub) Send(cid uuid.UUID, msg []byte) {
	c := h.clients[cid]
	if c == nil {
		return
	}
	select {
	case c.send <- msg:
	//	发送失败 关闭client
	default:
		close(c.send)
		delete(h.clients, cid)
	}
}

func (h *hub) Broadcast(msg []byte) {
	for cid := range h.clients {
		h.Send(cid, msg)
	}
}
