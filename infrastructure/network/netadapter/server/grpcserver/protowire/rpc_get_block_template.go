package protowire

import (
	"github.com/k1pool/entropyxd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *EntropyxdMessage_GetBlockTemplateRequest) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "EntropyxdMessage_GetBlockTemplateRequest is nil")
	}
	return x.GetBlockTemplateRequest.toAppMessage()
}

func (x *EntropyxdMessage_GetBlockTemplateRequest) fromAppMessage(message *appmessage.GetBlockTemplateRequestMessage) error {
	x.GetBlockTemplateRequest = &GetBlockTemplateRequestMessage{
		PayAddress: message.PayAddress,
		ExtraData:  message.ExtraData,
	}
	return nil
}

func (x *GetBlockTemplateRequestMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "GetBlockTemplateRequestMessage is nil")
	}
	return &appmessage.GetBlockTemplateRequestMessage{
		PayAddress: x.PayAddress,
		ExtraData:  x.ExtraData,
	}, nil
}

func (x *EntropyxdMessage_GetBlockTemplateResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "EntropyxdMessage_GetBlockTemplateResponse is nil")
	}
	return x.GetBlockTemplateResponse.toAppMessage()
}

func (x *EntropyxdMessage_GetBlockTemplateResponse) fromAppMessage(message *appmessage.GetBlockTemplateResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}

	var block *RpcBlock
	if message.Block != nil {
		protoBlock := &RpcBlock{}
		err := protoBlock.fromAppMessage(message.Block)
		if err != nil {
			return err
		}
		block = protoBlock
	}

	x.GetBlockTemplateResponse = &GetBlockTemplateResponseMessage{
		Block:    block,
		IsSynced: message.IsSynced,
		Error:    err,
	}
	return nil
}

func (x *GetBlockTemplateResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "GetBlockTemplateResponseMessage is nil")
	}
	var msgBlock *appmessage.RPCBlock
	if x.Block != nil {
		var err error
		msgBlock, err = x.Block.toAppMessage()
		if err != nil {
			return nil, err
		}
	}
	var rpcError *appmessage.RPCError
	if x.Error != nil {
		var err error
		rpcError, err = x.Error.toAppMessage()
		if err != nil {
			return nil, err
		}
	}
	return &appmessage.GetBlockTemplateResponseMessage{
		Block:    msgBlock,
		IsSynced: x.IsSynced,
		Error:    rpcError,
	}, nil
}
