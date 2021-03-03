package codesamples

import (
	"testing"

	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils"
	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils/testconstants"
	notsolo "github.com/brunoamancio/NotSolo"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/stretchr/testify/require"
)

/// This is a sample of how to send tokens from the value tangle to a chain, keeping the ownership of the tokens on that chain
func Test_SendTokensToChain_NoContractFees(t *testing.T) {
	notSolo := notsolo.New(t)

	require.GreaterOrEqual(t, initialWalletFunds, transferValueIotas)

	// Generates key pairs for sender and receiver wallets and provides both with dummy funds.
	senderWalletSigScheme := notSolo.SigScheme.NewSignatureSchemeWithFunds()

	// Wallet balance in value tangle -> Before transfer
	notSolo.ValueTangle.RequireBalance(senderWalletSigScheme, balance.ColorIOTA, initialWalletFunds)

	// Generate a dummy chain NOT beloging to sender
	chain := notSolo.Chain.NewChain(nil, "myChain")

	// Wallet balance in chain -> Before transfer
	notSolo.Chain.RequireBalance(senderWalletSigScheme, chain, balance.ColorIOTA, 0)

	// Transfer from sender's address in the value tangle to his account in 'chain'
	notSolo.ValueTangle.MustTransferToChainToSelf(senderWalletSigScheme, chain, balance.ColorIOTA, transferValueIotas)

	// Wallet balances -> After transfer to chain
	notSolo.ValueTangle.RequireBalance(senderWalletSigScheme, balance.ColorIOTA, initialWalletFunds-transferValueIotas-iotaTokensConsumedByRequest)
	notSolo.Chain.RequireBalance(senderWalletSigScheme, chain, balance.ColorIOTA, transferValueIotas+iotaTokensConsumedByRequest)
}

func Test_SendAndReceiveTokens_NoContractFees(t *testing.T) {
	notSolo := notsolo.New(t)

	require.GreaterOrEqual(t, initialWalletFunds, transferValueIotas)

	// Generates key pairs for sender and receiver wallets and provides both with dummy funds.
	senderWalletSigScheme := notSolo.SigScheme.NewSignatureSchemeWithFunds()

	// Wallet balance in value tangle -> Before transfer
	notSolo.ValueTangle.RequireBalance(senderWalletSigScheme, balance.ColorIOTA, initialWalletFunds)

	// Generates key pairs for sender wallet.
	receiverWalletSigScheme := notSolo.SigScheme.NewSignatureScheme()
	notSolo.ValueTangle.RequireBalance(receiverWalletSigScheme, balance.ColorIOTA, 0)

	// Generate a dummy chain NEITHER belonging to sender NOR receiver
	chain := notSolo.Chain.NewChain(nil, "myChain")

	// Transfer from sender's address in the value tangle to the receiver's account in 'chain'
	notSolo.ValueTangle.MustTransferToChain(senderWalletSigScheme, chain, balance.ColorIOTA, transferValueIotas, receiverWalletSigScheme)

	// Wallet balances -> After transfer
	notSolo.ValueTangle.RequireBalance(senderWalletSigScheme, balance.ColorIOTA, initialWalletFunds-transferValueIotas-iotaTokensConsumedByRequest)
	notSolo.Chain.RequireBalance(senderWalletSigScheme, chain, balance.ColorIOTA, iotaTokensConsumedByRequest)

	notSolo.ValueTangle.RequireBalance(receiverWalletSigScheme, balance.ColorIOTA, 0)
	notSolo.Chain.RequireBalance(receiverWalletSigScheme, chain, balance.ColorIOTA, transferValueIotas)
}

func Test_SendTokensToChain_WithContractFees(t *testing.T) {
	notSolo := notsolo.New(t)

	require.GreaterOrEqual(t, initialWalletFunds, transferValueIotas)

	// Generates key pairs for sender and receiver wallets and provides both with dummy funds.
	senderWalletSigScheme := notSolo.SigScheme.NewSignatureSchemeWithFunds()

	// Wallet balance in value tangle -> Before transfer
	notSolo.ValueTangle.RequireBalance(senderWalletSigScheme, balance.ColorIOTA, initialWalletFunds)

	// Generate a dummy chain NOT beloging to sender
	chain := notSolo.Chain.NewChain(nil, "myChain")
	contractName := "accounts" // This is a root contract, present in every chain

	// Check contract fees before changing them
	feeColor, ownerFee, validatorFee := chain.GetFeeInfo(contractName)
	require.Equal(t, balance.ColorIOTA, feeColor)
	require.Equal(t, int64(0), ownerFee)
	require.Equal(t, int64(0), validatorFee)

	// Request to change the contract fee settings in 'chain'
	const newContractOwnerFee = int64(100)
	notSolo.Chain.ChangeContractFees(chain.OriginatorSigScheme, chain, contractName, newContractOwnerFee)

	// Check contract fees after changing them
	feeColor, ownerFee, validatorFee = chain.GetFeeInfo(contractName)
	require.Equal(t, balance.ColorIOTA, feeColor)
	require.Equal(t, newContractOwnerFee, ownerFee)
	require.Equal(t, int64(0), validatorFee)

	// Wallet balance in chain -> Before transfer
	notSolo.Chain.RequireBalance(senderWalletSigScheme, chain, balance.ColorIOTA, 0)

	// Transfer from sender's address in the value tangle to his account in 'chain'
	notSolo.ValueTangle.MustTransferToChainToSelf(senderWalletSigScheme, chain, balance.ColorIOTA, transferValueIotas)

	// Balances -> After transfer to chain
	notSolo.ValueTangle.RequireBalance(senderWalletSigScheme, balance.ColorIOTA, initialWalletFunds-transferValueIotas-iotaTokensConsumedByRequest)   // Transfered tokens are debited from the value tangle
	notSolo.Chain.RequireBalance(senderWalletSigScheme, chain, balance.ColorIOTA, transferValueIotas+iotaTokensConsumedByRequest-newContractOwnerFee) // His tokens in the chain minus fees
}

func Test_SendTokensToContract_NoContractFees(t *testing.T) {
	notSolo := notsolo.New(t)

	// Generate a dummy chain NOT beloging to sender
	chain := notSolo.Chain.NewChain(nil, "myChain")

	// Uploads wasm of SC and deploys it into chain
	contractWasmFilePath := testutils.MustGetContractWasmFilePath(t, testconstants.ContractName) // You can use if file is in SmartContract/pkg
	notSolo.Chain.DeployWasmContract(chain, nil, testconstants.ContractName, contractWasmFilePath)

	// Generates key pairs for sender wallets, which will send iota tokens to the contract
	senderWalletSigScheme := notSolo.SigScheme.NewSignatureSchemeWithFunds()

	// Balance in value tangle -> Before transfer
	notSolo.ValueTangle.RequireBalance(senderWalletSigScheme, balance.ColorIOTA, initialWalletFunds)
	require.GreaterOrEqual(t, initialWalletFunds, transferValueIotas)

	// Transfer from sender's address in the value tangle to the contract's account in 'chain'
	notSolo.ValueTangle.MustTransferToContract(senderWalletSigScheme, chain, balance.ColorIOTA, transferValueIotas, testconstants.ContractName)

	// Balances -> After transfer
	notSolo.ValueTangle.RequireBalance(senderWalletSigScheme, balance.ColorIOTA, initialWalletFunds-transferValueIotas-iotaTokensConsumedByRequest)
	notSolo.Chain.RequireBalance(senderWalletSigScheme, chain, balance.ColorIOTA, iotaTokensConsumedByRequest)

	// Contract's account balance in the chain -> After transfer
	notSolo.Chain.RequireContractBalance(chain, testconstants.ContractName, balance.ColorIOTA, transferValueIotas)
}
