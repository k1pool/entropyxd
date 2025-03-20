package protowire

import (
	"github.com/k1pool/entropyxd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *EntropyxdMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "EntropyxdMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *EntropyxdMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
