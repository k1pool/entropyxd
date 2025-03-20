package protowire

import (
	"github.com/k1pool/entropyxd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *EntropyxdMessage_RequestNextHeaders) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "EntropyxdMessage_RequestNextHeaders is nil")
	}
	return &appmessage.MsgRequestNextHeaders{}, nil
}

func (x *EntropyxdMessage_RequestNextHeaders) fromAppMessage(_ *appmessage.MsgRequestNextHeaders) error {
	return nil
}
