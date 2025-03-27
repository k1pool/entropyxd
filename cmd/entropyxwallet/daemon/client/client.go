package client

import (
	"context"
	"time"

	"github.com/k1pool/entropyxd/cmd/entropyxwallet/daemon/server"

	"github.com/pkg/errors"

	"github.com/k1pool/entropyxd/cmd/entropyxwallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the entropyxwalletd server, and returns the client instance
func Connect(address string) (pb.EntropyxwalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("entropyxwallet daemon is not running, start it with `entropyxwallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewEntropyxwalletdClient(conn), func() {
		conn.Close()
	}, nil
}
