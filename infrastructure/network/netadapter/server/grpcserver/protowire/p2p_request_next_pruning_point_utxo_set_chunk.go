package protowire

import (
	"github.com/k1pool/entropyxd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *EntropyxdMessage_RequestNextPruningPointUtxoSetChunk) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "EntropyxdMessage_RequestNextPruningPointUtxoSetChunk is nil")
	}
	return &appmessage.MsgRequestNextPruningPointUTXOSetChunk{}, nil
}

func (x *EntropyxdMessage_RequestNextPruningPointUtxoSetChunk) fromAppMessage(_ *appmessage.MsgRequestNextPruningPointUTXOSetChunk) error {
	x.RequestNextPruningPointUtxoSetChunk = &RequestNextPruningPointUtxoSetChunkMessage{}
	return nil
}
