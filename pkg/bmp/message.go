package bmp

import "github.com/sbezverk/gobmp/pkg/base"

// Message defines a message used to transfer BMP messages for further processing
// for BMP messages which do not carry PerPeerHeader, it will be set to nil.
type Message struct {
	L3p        base.L3Pkt
	PeerHeader *PerPeerHeader
	Payload    interface{}
}
