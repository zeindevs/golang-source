package node

import (
	"context"
	"fmt"
	"net"

	"github.com/zeindevs/go-blocker/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type Node struct {
	proto.UnimplementedNodeServer
	peers   map[net.Addr]*grpc.ClientConn
	version string
}

func NewNode() *Node {
	return &Node{
		version: "blocker-0.1",
	}
}
func (n *Node) Handshake(ctx context.Context, v *proto.Version) (*proto.Version, error) {
	ourVersion := &proto.Version{
		Version: n.version,
		Height:  100,
	}

	p, _ := peer.FromContext(ctx)

	fmt.Printf("received version form %s : %+v\n", v, p.Addr)

	return ourVersion, nil
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {
	peer, _ := peer.FromContext(ctx)
	fmt.Println("received tx from:", peer)
	return &proto.Ack{}, nil
}
