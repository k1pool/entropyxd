package server

import (
	"testing"

	"github.com/k1pool/entropyxd/cmd/entropyxwallet/libentropyxwallet/serialization"

	"github.com/k1pool/entropyxd/cmd/entropyxwallet/keys"
	"github.com/k1pool/entropyxd/util/txmass"

	"github.com/k1pool/entropyxd/domain/dagconfig"

	"github.com/k1pool/entropyxd/domain/consensus/model/externalapi"
	"github.com/k1pool/entropyxd/domain/consensus/utils/consensushashing"
	"github.com/k1pool/entropyxd/domain/consensus/utils/txscript"
	"github.com/k1pool/entropyxd/domain/consensus/utils/utxo"

	"github.com/k1pool/entropyxd/cmd/entropyxwallet/libentropyxwallet"
	"github.com/k1pool/entropyxd/domain/consensus"
	"github.com/k1pool/entropyxd/domain/consensus/utils/testutils"
)

func TestEstimateComputeMassAfterSignatures(t *testing.T) {
	testutils.ForAllNets(t, true, func(t *testing.T, consensusConfig *consensus.Config) {
		unsignedTransaction, mnemonics, params, teardown := testEstimateMassIncreaseForSignaturesSetUp(t, consensusConfig)
		defer teardown(false)

		serverInstance := &server{
			params:           params,
			keysFile:         &keys.File{MinimumSignatures: 2},
			shutdown:         make(chan struct{}),
			addressSet:       make(walletAddressSet),
			txMassCalculator: txmass.NewCalculator(params.MassPerTxByte, params.MassPerScriptPubKeyByte, params.MassPerSigOp),
		}

		estimatedMassAfterSignatures, err := serverInstance.estimateComputeMassAfterSignatures(unsignedTransaction)
		if err != nil {
			t.Fatalf("Error from estimateComputeMassAfterSignatures: %s", err)
		}

		unsignedTransactionBytes, err := serialization.SerializePartiallySignedTransaction(unsignedTransaction)
		if err != nil {
			t.Fatalf("Error deserializing unsignedTransaction: %s", err)
		}

		signedTxStep1Bytes, err := libentropyxwallet.Sign(params, mnemonics[:1], unsignedTransactionBytes, false)
		if err != nil {
			t.Fatalf("Sign: %+v", err)
		}

		signedTxStep2Bytes, err := libentropyxwallet.Sign(params, mnemonics[1:2], signedTxStep1Bytes, false)
		if err != nil {
			t.Fatalf("Sign: %+v", err)
		}

		extractedSignedTx, err := libentropyxwallet.ExtractTransaction(signedTxStep2Bytes, false)
		if err != nil {
			t.Fatalf("ExtractTransaction: %+v", err)
		}

		actualMassAfterSignatures := serverInstance.txMassCalculator.CalculateTransactionMass(extractedSignedTx)

		if estimatedMassAfterSignatures != actualMassAfterSignatures {
			t.Errorf("Estimated mass after signatures: %d but actually got %d",
				estimatedMassAfterSignatures, actualMassAfterSignatures)
		}
	})
}

func TestEstimateMassAfterSignatures(t *testing.T) {
	testutils.ForAllNets(t, true, func(t *testing.T, consensusConfig *consensus.Config) {
		unsignedTransaction, mnemonics, params, teardown := testEstimateMassIncreaseForSignaturesSetUp(t, consensusConfig)
		defer teardown(false)

		for i := range unsignedTransaction.Tx.Inputs {
			unsignedTransaction.Tx.Inputs[i].UTXOEntry = utxo.NewUTXOEntry(1, &externalapi.ScriptPublicKey{}, false, 0)
			unsignedTransaction.PartiallySignedInputs[i].PrevOutput = &externalapi.DomainTransactionOutput{
				Value:           1,
				ScriptPublicKey: &externalapi.ScriptPublicKey{},
			}
		}

		serverInstance := &server{
			params:           params,
			keysFile:         &keys.File{MinimumSignatures: 2},
			shutdown:         make(chan struct{}),
			addressSet:       make(walletAddressSet),
			txMassCalculator: txmass.NewCalculator(params.MassPerTxByte, params.MassPerScriptPubKeyByte, params.MassPerSigOp),
		}

		estimatedMassAfterSignatures, err := serverInstance.estimateMassAfterSignatures(unsignedTransaction)
		if err != nil {
			t.Fatalf("Error from estimateMassAfterSignatures: %s", err)
		}

		unsignedTransactionBytes, err := serialization.SerializePartiallySignedTransaction(unsignedTransaction)
		if err != nil {
			t.Fatalf("Error deserializing unsignedTransaction: %s", err)
		}

		signedTxStep1Bytes, err := libentropyxwallet.Sign(params, mnemonics[:1], unsignedTransactionBytes, false)
		if err != nil {
			t.Fatalf("Sign: %+v", err)
		}

		signedTxStep2Bytes, err := libentropyxwallet.Sign(params, mnemonics[1:2], signedTxStep1Bytes, false)
		if err != nil {
			t.Fatalf("Sign: %+v", err)
		}

		extractedSignedTx, err := libentropyxwallet.ExtractTransaction(signedTxStep2Bytes, false)
		if err != nil {
			t.Fatalf("ExtractTransaction: %+v", err)
		}

		for i := range extractedSignedTx.Inputs {
			extractedSignedTx.Inputs[i].UTXOEntry = utxo.NewUTXOEntry(1, &externalapi.ScriptPublicKey{}, false, 0)
		}

		actualMassAfterSignatures := serverInstance.txMassCalculator.CalculateTransactionOverallMass(extractedSignedTx)

		if estimatedMassAfterSignatures != actualMassAfterSignatures {
			t.Errorf("Estimated mass after signatures: %d but actually got %d",
				estimatedMassAfterSignatures, actualMassAfterSignatures)
		}
	})
}

func testEstimateMassIncreaseForSignaturesSetUp(t *testing.T, consensusConfig *consensus.Config) (
	*serialization.PartiallySignedTransaction, []string, *dagconfig.Params, func(keepDataDir bool)) {

	consensusConfig.BlockCoinbaseMaturity = 0
	params := &consensusConfig.Params

	tc, teardown, err := consensus.NewFactory().NewTestConsensus(consensusConfig, "TestMultisig")
	if err != nil {
		t.Fatalf("Error setting up tc: %+v", err)
	}

	const numKeys = 3
	mnemonics := make([]string, numKeys)
	publicKeys := make([]string, numKeys)
	for i := 0; i < numKeys; i++ {
		var err error
		mnemonics[i], err = libentropyxwallet.CreateMnemonic()
		if err != nil {
			t.Fatalf("CreateMnemonic: %+v", err)
		}

		publicKeys[i], err = libentropyxwallet.MasterPublicKeyFromMnemonic(&consensusConfig.Params, mnemonics[i], true)
		if err != nil {
			t.Fatalf("MasterPublicKeyFromMnemonic: %+v", err)
		}
	}

	const minimumSignatures = 2
	path := "m/1/2/3"
	address, err := libentropyxwallet.Address(params, publicKeys, minimumSignatures, path, false)
	if err != nil {
		t.Fatalf("Address: %+v", err)
	}

	scriptPublicKey, err := txscript.PayToAddrScript(address)
	if err != nil {
		t.Fatalf("PayToAddrScript: %+v", err)
	}

	coinbaseData := &externalapi.DomainCoinbaseData{
		ScriptPublicKey: scriptPublicKey,
		ExtraData:       nil,
	}

	fundingBlockHash, _, err := tc.AddBlock([]*externalapi.DomainHash{consensusConfig.GenesisHash}, coinbaseData, nil)
	if err != nil {
		t.Fatalf("AddBlock: %+v", err)
	}

	block1Hash, _, err := tc.AddBlock([]*externalapi.DomainHash{fundingBlockHash}, nil, nil)
	if err != nil {
		t.Fatalf("AddBlock: %+v", err)
	}

	block1, _, err := tc.GetBlock(block1Hash)
	if err != nil {
		t.Fatalf("GetBlock: %+v", err)
	}

	block1Tx := block1.Transactions[0]
	block1TxOut := block1Tx.Outputs[0]
	selectedUTXOs := []*libentropyxwallet.UTXO{
		{
			Outpoint: &externalapi.DomainOutpoint{
				TransactionID: *consensushashing.TransactionID(block1.Transactions[0]),
				Index:         0,
			},
			UTXOEntry:      utxo.NewUTXOEntry(block1TxOut.Value, block1TxOut.ScriptPublicKey, true, 0),
			DerivationPath: path,
		},
	}

	unsignedTransaction, err := libentropyxwallet.CreateUnsignedTransaction(publicKeys, minimumSignatures,
		[]*libentropyxwallet.Payment{{
			Address: address,
			Amount:  10,
		}}, selectedUTXOs)
	if err != nil {
		t.Fatalf("CreateUnsignedTransactions: %+v", err)
	}

	return unsignedTransaction, mnemonics, params, teardown
}
