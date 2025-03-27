package protowire

import (
	"github.com/k1pool/entropyxd/app/appmessage"
	"github.com/pkg/errors"
)

type converter interface {
	toAppMessage() (appmessage.Message, error)
}

// ToAppMessage converts a EntropyxdMessage to its appmessage.Message representation
func (x *EntropyxdMessage) ToAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "EntropyxdMessage is nil")
	}
	converter, ok := x.Payload.(converter)
	if !ok {
		return nil, errors.Errorf("received invalid message")
	}
	appMessage, err := converter.toAppMessage()
	if err != nil {
		return nil, err
	}
	return appMessage, nil
}

// FromAppMessage creates a EntropyxdMessage from a appmessage.Message
func FromAppMessage(message appmessage.Message) (*EntropyxdMessage, error) {
	payload, err := toPayload(message)
	if err != nil {
		return nil, err
	}
	return &EntropyxdMessage{
		Payload: payload,
	}, nil
}

func toPayload(message appmessage.Message) (isEntropyxdMessage_Payload, error) {
	p2pPayload, err := toP2PPayload(message)
	if err != nil {
		return nil, err
	}
	if p2pPayload != nil {
		return p2pPayload, nil
	}

	rpcPayload, err := toRPCPayload(message)
	if err != nil {
		return nil, err
	}
	if rpcPayload != nil {
		return rpcPayload, nil
	}

	return nil, errors.Errorf("unknown message type %T", message)
}

func toP2PPayload(message appmessage.Message) (isEntropyxdMessage_Payload, error) {
	switch message := message.(type) {
	case *appmessage.MsgAddresses:
		payload := new(EntropyxdMessage_Addresses)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgBlock:
		payload := new(EntropyxdMessage_Block)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestBlockLocator:
		payload := new(EntropyxdMessage_RequestBlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgBlockLocator:
		payload := new(EntropyxdMessage_BlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestAddresses:
		payload := new(EntropyxdMessage_RequestAddresses)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestIBDBlocks:
		payload := new(EntropyxdMessage_RequestIBDBlocks)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestNextHeaders:
		payload := new(EntropyxdMessage_RequestNextHeaders)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgDoneHeaders:
		payload := new(EntropyxdMessage_DoneHeaders)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestRelayBlocks:
		payload := new(EntropyxdMessage_RequestRelayBlocks)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestTransactions:
		payload := new(EntropyxdMessage_RequestTransactions)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgTransactionNotFound:
		payload := new(EntropyxdMessage_TransactionNotFound)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgInvRelayBlock:
		payload := new(EntropyxdMessage_InvRelayBlock)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgInvTransaction:
		payload := new(EntropyxdMessage_InvTransactions)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPing:
		payload := new(EntropyxdMessage_Ping)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPong:
		payload := new(EntropyxdMessage_Pong)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgTx:
		payload := new(EntropyxdMessage_Transaction)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgVerAck:
		payload := new(EntropyxdMessage_Verack)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgVersion:
		payload := new(EntropyxdMessage_Version)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgReject:
		payload := new(EntropyxdMessage_Reject)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestPruningPointUTXOSet:
		payload := new(EntropyxdMessage_RequestPruningPointUTXOSet)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPruningPointUTXOSetChunk:
		payload := new(EntropyxdMessage_PruningPointUtxoSetChunk)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgUnexpectedPruningPoint:
		payload := new(EntropyxdMessage_UnexpectedPruningPoint)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDBlockLocator:
		payload := new(EntropyxdMessage_IbdBlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDBlockLocatorHighestHash:
		payload := new(EntropyxdMessage_IbdBlockLocatorHighestHash)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDBlockLocatorHighestHashNotFound:
		payload := new(EntropyxdMessage_IbdBlockLocatorHighestHashNotFound)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.BlockHeadersMessage:
		payload := new(EntropyxdMessage_BlockHeaders)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestNextPruningPointUTXOSetChunk:
		payload := new(EntropyxdMessage_RequestNextPruningPointUtxoSetChunk)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgDonePruningPointUTXOSetChunks:
		payload := new(EntropyxdMessage_DonePruningPointUtxoSetChunks)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgBlockWithTrustedData:
		payload := new(EntropyxdMessage_BlockWithTrustedData)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestPruningPointAndItsAnticone:
		payload := new(EntropyxdMessage_RequestPruningPointAndItsAnticone)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgDoneBlocksWithTrustedData:
		payload := new(EntropyxdMessage_DoneBlocksWithTrustedData)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDBlock:
		payload := new(EntropyxdMessage_IbdBlock)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestHeaders:
		payload := new(EntropyxdMessage_RequestHeaders)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPruningPoints:
		payload := new(EntropyxdMessage_PruningPoints)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestPruningPointProof:
		payload := new(EntropyxdMessage_RequestPruningPointProof)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPruningPointProof:
		payload := new(EntropyxdMessage_PruningPointProof)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgReady:
		payload := new(EntropyxdMessage_Ready)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgTrustedData:
		payload := new(EntropyxdMessage_TrustedData)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgBlockWithTrustedDataV4:
		payload := new(EntropyxdMessage_BlockWithTrustedDataV4)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestNextPruningPointAndItsAnticoneBlocks:
		payload := new(EntropyxdMessage_RequestNextPruningPointAndItsAnticoneBlocks)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestIBDChainBlockLocator:
		payload := new(EntropyxdMessage_RequestIBDChainBlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDChainBlockLocator:
		payload := new(EntropyxdMessage_IbdChainBlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestAnticone:
		payload := new(EntropyxdMessage_RequestAnticone)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	default:
		return nil, nil
	}
}

func toRPCPayload(message appmessage.Message) (isEntropyxdMessage_Payload, error) {
	switch message := message.(type) {
	case *appmessage.GetCurrentNetworkRequestMessage:
		payload := new(EntropyxdMessage_GetCurrentNetworkRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetCurrentNetworkResponseMessage:
		payload := new(EntropyxdMessage_GetCurrentNetworkResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitBlockRequestMessage:
		payload := new(EntropyxdMessage_SubmitBlockRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitBlockResponseMessage:
		payload := new(EntropyxdMessage_SubmitBlockResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockTemplateRequestMessage:
		payload := new(EntropyxdMessage_GetBlockTemplateRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockTemplateResponseMessage:
		payload := new(EntropyxdMessage_GetBlockTemplateResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyBlockAddedRequestMessage:
		payload := new(EntropyxdMessage_NotifyBlockAddedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyBlockAddedResponseMessage:
		payload := new(EntropyxdMessage_NotifyBlockAddedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.BlockAddedNotificationMessage:
		payload := new(EntropyxdMessage_BlockAddedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetPeerAddressesRequestMessage:
		payload := new(EntropyxdMessage_GetPeerAddressesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetPeerAddressesResponseMessage:
		payload := new(EntropyxdMessage_GetPeerAddressesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetSelectedTipHashRequestMessage:
		payload := new(EntropyxdMessage_GetSelectedTipHashRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetSelectedTipHashResponseMessage:
		payload := new(EntropyxdMessage_GetSelectedTipHashResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntryRequestMessage:
		payload := new(EntropyxdMessage_GetMempoolEntryRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntryResponseMessage:
		payload := new(EntropyxdMessage_GetMempoolEntryResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetConnectedPeerInfoRequestMessage:
		payload := new(EntropyxdMessage_GetConnectedPeerInfoRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetConnectedPeerInfoResponseMessage:
		payload := new(EntropyxdMessage_GetConnectedPeerInfoResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.AddPeerRequestMessage:
		payload := new(EntropyxdMessage_AddPeerRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.AddPeerResponseMessage:
		payload := new(EntropyxdMessage_AddPeerResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitTransactionRequestMessage:
		payload := new(EntropyxdMessage_SubmitTransactionRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitTransactionResponseMessage:
		payload := new(EntropyxdMessage_SubmitTransactionResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualSelectedParentChainChangedRequestMessage:
		payload := new(EntropyxdMessage_NotifyVirtualSelectedParentChainChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualSelectedParentChainChangedResponseMessage:
		payload := new(EntropyxdMessage_NotifyVirtualSelectedParentChainChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.VirtualSelectedParentChainChangedNotificationMessage:
		payload := new(EntropyxdMessage_VirtualSelectedParentChainChangedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockRequestMessage:
		payload := new(EntropyxdMessage_GetBlockRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockResponseMessage:
		payload := new(EntropyxdMessage_GetBlockResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetSubnetworkRequestMessage:
		payload := new(EntropyxdMessage_GetSubnetworkRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetSubnetworkResponseMessage:
		payload := new(EntropyxdMessage_GetSubnetworkResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetVirtualSelectedParentChainFromBlockRequestMessage:
		payload := new(EntropyxdMessage_GetVirtualSelectedParentChainFromBlockRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetVirtualSelectedParentChainFromBlockResponseMessage:
		payload := new(EntropyxdMessage_GetVirtualSelectedParentChainFromBlockResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlocksRequestMessage:
		payload := new(EntropyxdMessage_GetBlocksRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlocksResponseMessage:
		payload := new(EntropyxdMessage_GetBlocksResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockCountRequestMessage:
		payload := new(EntropyxdMessage_GetBlockCountRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockCountResponseMessage:
		payload := new(EntropyxdMessage_GetBlockCountResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockDAGInfoRequestMessage:
		payload := new(EntropyxdMessage_GetBlockDagInfoRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockDAGInfoResponseMessage:
		payload := new(EntropyxdMessage_GetBlockDagInfoResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.ResolveFinalityConflictRequestMessage:
		payload := new(EntropyxdMessage_ResolveFinalityConflictRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.ResolveFinalityConflictResponseMessage:
		payload := new(EntropyxdMessage_ResolveFinalityConflictResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyFinalityConflictsRequestMessage:
		payload := new(EntropyxdMessage_NotifyFinalityConflictsRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyFinalityConflictsResponseMessage:
		payload := new(EntropyxdMessage_NotifyFinalityConflictsResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.FinalityConflictNotificationMessage:
		payload := new(EntropyxdMessage_FinalityConflictNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.FinalityConflictResolvedNotificationMessage:
		payload := new(EntropyxdMessage_FinalityConflictResolvedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntriesRequestMessage:
		payload := new(EntropyxdMessage_GetMempoolEntriesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntriesResponseMessage:
		payload := new(EntropyxdMessage_GetMempoolEntriesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.ShutDownRequestMessage:
		payload := new(EntropyxdMessage_ShutDownRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.ShutDownResponseMessage:
		payload := new(EntropyxdMessage_ShutDownResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetHeadersRequestMessage:
		payload := new(EntropyxdMessage_GetHeadersRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetHeadersResponseMessage:
		payload := new(EntropyxdMessage_GetHeadersResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyUTXOsChangedRequestMessage:
		payload := new(EntropyxdMessage_NotifyUtxosChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyUTXOsChangedResponseMessage:
		payload := new(EntropyxdMessage_NotifyUtxosChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.UTXOsChangedNotificationMessage:
		payload := new(EntropyxdMessage_UtxosChangedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.StopNotifyingUTXOsChangedRequestMessage:
		payload := new(EntropyxdMessage_StopNotifyingUtxosChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.StopNotifyingUTXOsChangedResponseMessage:
		payload := new(EntropyxdMessage_StopNotifyingUtxosChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetUTXOsByAddressesRequestMessage:
		payload := new(EntropyxdMessage_GetUtxosByAddressesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetUTXOsByAddressesResponseMessage:
		payload := new(EntropyxdMessage_GetUtxosByAddressesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBalanceByAddressRequestMessage:
		payload := new(EntropyxdMessage_GetBalanceByAddressRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBalanceByAddressResponseMessage:
		payload := new(EntropyxdMessage_GetBalanceByAddressResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetVirtualSelectedParentBlueScoreRequestMessage:
		payload := new(EntropyxdMessage_GetVirtualSelectedParentBlueScoreRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetVirtualSelectedParentBlueScoreResponseMessage:
		payload := new(EntropyxdMessage_GetVirtualSelectedParentBlueScoreResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualSelectedParentBlueScoreChangedRequestMessage:
		payload := new(EntropyxdMessage_NotifyVirtualSelectedParentBlueScoreChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualSelectedParentBlueScoreChangedResponseMessage:
		payload := new(EntropyxdMessage_NotifyVirtualSelectedParentBlueScoreChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.VirtualSelectedParentBlueScoreChangedNotificationMessage:
		payload := new(EntropyxdMessage_VirtualSelectedParentBlueScoreChangedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.BanRequestMessage:
		payload := new(EntropyxdMessage_BanRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.BanResponseMessage:
		payload := new(EntropyxdMessage_BanResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.UnbanRequestMessage:
		payload := new(EntropyxdMessage_UnbanRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.UnbanResponseMessage:
		payload := new(EntropyxdMessage_UnbanResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetInfoRequestMessage:
		payload := new(EntropyxdMessage_GetInfoRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetInfoResponseMessage:
		payload := new(EntropyxdMessage_GetInfoResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyPruningPointUTXOSetOverrideRequestMessage:
		payload := new(EntropyxdMessage_NotifyPruningPointUTXOSetOverrideRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyPruningPointUTXOSetOverrideResponseMessage:
		payload := new(EntropyxdMessage_NotifyPruningPointUTXOSetOverrideResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.PruningPointUTXOSetOverrideNotificationMessage:
		payload := new(EntropyxdMessage_PruningPointUTXOSetOverrideNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.StopNotifyingPruningPointUTXOSetOverrideRequestMessage:
		payload := new(EntropyxdMessage_StopNotifyingPruningPointUTXOSetOverrideRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.EstimateNetworkHashesPerSecondRequestMessage:
		payload := new(EntropyxdMessage_EstimateNetworkHashesPerSecondRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.EstimateNetworkHashesPerSecondResponseMessage:
		payload := new(EntropyxdMessage_EstimateNetworkHashesPerSecondResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualDaaScoreChangedRequestMessage:
		payload := new(EntropyxdMessage_NotifyVirtualDaaScoreChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualDaaScoreChangedResponseMessage:
		payload := new(EntropyxdMessage_NotifyVirtualDaaScoreChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.VirtualDaaScoreChangedNotificationMessage:
		payload := new(EntropyxdMessage_VirtualDaaScoreChangedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBalancesByAddressesRequestMessage:
		payload := new(EntropyxdMessage_GetBalancesByAddressesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBalancesByAddressesResponseMessage:
		payload := new(EntropyxdMessage_GetBalancesByAddressesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyNewBlockTemplateRequestMessage:
		payload := new(EntropyxdMessage_NotifyNewBlockTemplateRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyNewBlockTemplateResponseMessage:
		payload := new(EntropyxdMessage_NotifyNewBlockTemplateResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NewBlockTemplateNotificationMessage:
		payload := new(EntropyxdMessage_NewBlockTemplateNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntriesByAddressesRequestMessage:
		payload := new(EntropyxdMessage_GetMempoolEntriesByAddressesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntriesByAddressesResponseMessage:
		payload := new(EntropyxdMessage_GetMempoolEntriesByAddressesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetCoinSupplyRequestMessage:
		payload := new(EntropyxdMessage_GetCoinSupplyRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetCoinSupplyResponseMessage:
		payload := new(EntropyxdMessage_GetCoinSupplyResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetFeeEstimateRequestMessage:
		payload := new(EntropyxdMessage_GetFeeEstimateRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitTransactionReplacementRequestMessage:
		payload := new(EntropyxdMessage_SubmitTransactionReplacementRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitTransactionReplacementResponseMessage:
		payload := new(EntropyxdMessage_SubmitTransactionReplacementResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	default:
		return nil, nil
	}
}
