package p2p

// Peer is an interface that represents the remote node.
type Peer interface{}

// Transport is anything that handles the communication
// Between the nodes in the network
// from (TCP, UDP, websockets, ...)
type Transport interface {
	ListenAndAccept() error
}
