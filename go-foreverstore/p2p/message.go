package p2p

import "net"

// Message holds any arbirary data that is being send over the
// each transport between two nodes in the network
type Message struct {
	From    net.Addr
	Payload []byte
}
