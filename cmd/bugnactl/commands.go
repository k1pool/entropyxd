package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/k1pool/entropyxd/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.EntropyxdMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.EntropyxdMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.EntropyxdMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.EntropyxdMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.EntropyxdMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.EntropyxdMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.EntropyxdMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.EntropyxdMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.EntropyxdMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.EntropyxdMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.EntropyxdMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.EntropyxdMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.EntropyxdMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.EntropyxdMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.EntropyxdMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.EntropyxdMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.EntropyxdMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.EntropyxdMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.EntropyxdMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.EntropyxdMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.EntropyxdMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.EntropyxdMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.EntropyxdMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.EntropyxdMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.EntropyxdMessage_BanRequest{}),
	reflect.TypeOf(protowire.EntropyxdMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
