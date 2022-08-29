package webrtc

import (
	"github.com/fasthttp/websocket"
	"log"
	"sync"
	"videochat-project/pkg/chat"
)

type Room struct {
	Peers *Peers
	Hub   *chat.Hub
}

func RoomConn(c *websocket.Conn, p *Peers) {
	var config webrtc.Configuration
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		log.Println(err)
		return
	}
	newPeer := PeerConnectionState{
		PeerConnection: peerConnection,
		Websocket:      &ThreadSafeWriter{},
		Conn:           c,
		Mutex:          sync.Mutex{},
	}
}
