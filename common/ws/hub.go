package ws

import (
	"github.com/google/uuid"
)

type hub struct {
	// Registered clients.
	clients map[uuid.UUID]*client
	// Inbound messages from the clients.
	broadcast chan []byte
	// Register requests from the clients.
	register chan *client
	// Unregister requests from clients.
	unregister chan *client
}

func newHub() *hub {
	res := &hub{
		broadcast:  make(chan []byte),
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
		case c := <-h.unregister:
			if _, ok := h.clients[c.id]; ok {
				delete(h.clients, c.id)
				close(c.send)
			}
		case message := <-h.broadcast:
			for cid, c := range h.clients {
				select {
				case c.send <- message:
				//	发送失败 关闭client
				default:
					close(c.send)
					delete(h.clients, cid)
				}
			}
		}
	}
}
