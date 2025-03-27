package server

import (
	"context"
	"github.com/k1pool/entropyxd/cmd/entropyxwallet/daemon/pb"
	"github.com/k1pool/entropyxd/version"
)

func (s *server) GetVersion(_ context.Context, _ *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return &pb.GetVersionResponse{
		Version: version.Version(),
	}, nil
}
