package webrtc

import (
	"sync"
	chat "videochat-project/pkg/chat"
)


type Peers struct {
	ListLock    sync.RWMutex
	connections []PeerConnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticIP
}

func (p *Peers) DispatchKeyFrame() {

}
