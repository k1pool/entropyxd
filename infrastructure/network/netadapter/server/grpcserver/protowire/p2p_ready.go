package protowire

import (
	"github.com/k1pool/entropyxd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *EntropyxdMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "EntropyxdMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *EntropyxdMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
