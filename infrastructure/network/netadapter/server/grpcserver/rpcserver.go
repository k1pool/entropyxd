package grpcserver

import (
	"github.com/k1pool/entropyxd/infrastructure/network/netadapter/server"
	"github.com/k1pool/entropyxd/infrastructure/network/netadapter/server/grpcserver/protowire"
	"github.com/k1pool/entropyxd/util/panics"
)

type rpcServer struct {
	protowire.UnimplementedRPCServer
	gRPCServer
}

// RPCMaxMessageSize is the max message size for the RPC server to send and receive
const RPCMaxMessageSize = 1024 * 1024 * 1024 // 1 GB

// NewRPCServer creates a new RPCServer
func NewRPCServer(listeningAddresses []string, rpcMaxInboundConnections int) (server.Server, error) {
	gRPCServer := newGRPCServer(listeningAddresses, RPCMaxMessageSize, rpcMaxInboundConnections, "RPC")
	rpcServer := &rpcServer{gRPCServer: *gRPCServer}
	protowire.RegisterRPCServer(gRPCServer.server, rpcServer)
	return rpcServer, nil
}

func (r *rpcServer) MessageStream(stream protowire.RPC_MessageStreamServer) error {
	defer panics.HandlePanic(log, "rpcServer.MessageStream", nil)

	return r.handleInboundConnection(stream.Context(), stream)
}
