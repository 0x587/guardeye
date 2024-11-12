package ws

import (
	"bytes"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// client is a middleman between the websocket connection and the hub.
type client struct {
	id  uuid.UUID
	hub *hub
	// The websocket connection.
	conn *websocket.Conn
	// Buffered channel of outbound messages.
	send chan []byte
}

func newClient(h *hub, conn *websocket.Conn) *client {
	return &client{
		id:   uuid.New(),
		hub:  h,
		conn: conn,
		send: make(chan []byte, bufSize),
	}
}

func (c *client) run() {
	c.hub.register <- c
	go c.readLoop()
	go c.writeLoop()
}

func (c *client) close() {
	_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
	_ = c.conn.Close()
	close(c.send)
}

func (c *client) readLoop() {
	defer func() {
		c.hub.unregister <- c
		_ = c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { _ = c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logx.Errorf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(message)
		c.hub.broadcast <- message
	}
}

func (c *client) writeLoop() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		_ = c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, _ = w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
