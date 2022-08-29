package chat

import (
	"github.com/fasthttp/websocket"
	"time"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

type Client struct {
	Hub  *Hub
	conn *websocket.Conn
	Send chan []byte
}

func (c *Client) readPump() {

}
func (c *Client) writePump() {

}
func PeerChatConn() {

}
