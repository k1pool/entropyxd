package main

import (
	"context"
	"fmt"
	"github.com/k1pool/entropyxd/cmd/entropyxwallet/daemon/client"
	"github.com/k1pool/entropyxd/cmd/entropyxwallet/daemon/pb"
)

func newAddress(conf *newAddressConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()

	response, err := daemonClient.NewAddress(ctx, &pb.NewAddressRequest{})
	if err != nil {
		return err
	}

	fmt.Printf("New address:\n%s\n", response.Address)
	return nil
}
