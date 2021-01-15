package msgswap

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative follow.proto

import (
	"fmt"

	"github.com/Mrcrypt/5words/protoio"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/protocol"
)

//FollowID ...
var FollowID = protocol.ID("/follow/1.0.0")

//FollowHandleStream ...
func FollowHandleStream(callback func(*NewFollower)) network.StreamHandler {
	return func(stream network.Stream) {
		fmt.Println("Got a new stream!")

		reader := protoio.NewDelimitedReader(stream, 8*1024)

		m := NewFollower{}
		reader.ReadMsg(&m)

		fmt.Printf("You Recieved a Follower: %s\n", string(m.Name))
		callback(&m)
	}
}
