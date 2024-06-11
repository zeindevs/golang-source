package p2p

import "errors"

// ErrInvalidHandshake is returned if the hanshake between
// the local and remote node could not be established
var ErrInvalidHandshake = errors.New("invalid handshake")

// HandshakerFunc... ?
type HandshakerFunc func(Peer) error

func NewHandshakerFunc(Peer) error { return nil }

func NOPHandshakeFunc(Peer) error { return nil }
