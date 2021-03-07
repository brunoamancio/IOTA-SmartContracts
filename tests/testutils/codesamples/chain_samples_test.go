package codesamples

import (
	"testing"

	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils/testconstants"
	notsolo "github.com/brunoamancio/NotSolo"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/stretchr/testify/require"
)

func Test_CreateChain_chainCreatorSpecified(t *testing.T) {
	notSolo := notsolo.New(t)

	// Create a signature scheme with dummy tokens in it.
	chainOriginatorsSigScheme := notSolo.SigScheme.NewSignatureSchemeWithFunds()

	// Balance in value tangle before chain is created
	notSolo.ValueTangle.RequireBalance(chainOriginatorsSigScheme, balance.ColorIOTA, initialWalletFunds)

	// Create a chain where chainOriginatorsSigScheme is the owner.
	chain := notSolo.Chain.NewChain(chainOriginatorsSigScheme, "myChain")
	require.NotEqual(t, chain.ChainColor, balance.ColorIOTA)

	// Balance in value tangle after chain is created
	notSolo.ValueTangle.RequireBalance(chainOriginatorsSigScheme, balance.ColorIOTA, initialWalletFunds-iotaTokensConsumedByRequest-iotaTokensConsumedByChain)

	// IMPORTANT: When a chain is created, 1 IOTA is colored with the chain's color and sent to the chain's address in the value tangle
	notSolo.ValueTangle.RequireAddressBalance(chain.ChainAddress, chain.ChainColor, iotaTokensConsumedByChain)

	// IMPORTANT: When a chain is created, 1 IOTA is sent from the originator's account in the value tangle their account in the chain
	notSolo.Chain.RequireBalance(chainOriginatorsSigScheme, chain, balance.ColorIOTA, iotaTokensConsumedByRequest)
}

// Sample of how to create chain without specifying a chainOriginator.
// A dummy chain originator is created in the background (by NewChain).
func Test_CreateChain_NoChainCreatorSpecified(t *testing.T) {
	notSolo := notsolo.New(t)

	// Create a chain where chainOriginatorsSigScheme is the owner.
	chain := notSolo.Chain.NewChain(nil, "myChain")
	require.NotEqual(t, chain.ChainColor, balance.ColorIOTA)

	// IMPORTANT: When a chain is created, 1 IOTA is colored with the chain's color and sent to the chain's address in the value tangle
	notSolo.ValueTangle.RequireBalance(chain.ChainSigScheme, chain.ChainColor, iotaTokensConsumedByChain)

	// IMPORTANT: When a chain is created, 1 IOTA is sent from the originator's account in the value tangle their account in the chain
	notSolo.Chain.RequireBalance(chain.OriginatorSigScheme, chain, balance.ColorIOTA, iotaTokensConsumedByRequest)
}

func Test_SetChainFees(t *testing.T) {
	notSolo := notsolo.New(t)

	// Generate a dummy chain beloging to chain.OriginatorSigScheme
	chain := notSolo.Chain.NewChain(nil, "myChain")

	// Initial chain fees
	feeColor, ownerFee, validatorFee := chain.GetFeeInfo(testconstants.AccountsContractName)
	require.Equal(t, balance.ColorIOTA, feeColor)
	require.Equal(t, int64(0), ownerFee)
	require.Equal(t, int64(0), validatorFee)

	// Request to chain - change contract owner fee settings
	const newChainOwnerFee = int64(1000)
	notSolo.Chain.ChangeContractFees(chain.OriginatorSigScheme, chain, testconstants.AccountsContractName, newChainOwnerFee)

	// Request to chain - change validator fee settings
	const newValidatorFee = int64(200)
	notSolo.Chain.ChangeValidatorFees(chain.OriginatorSigScheme, chain, testconstants.AccountsContractName, newValidatorFee)

	// Chain fees after change
	feeColor, ownerFee, validatorFee = chain.GetFeeInfo(testconstants.AccountsContractName)
	require.Equal(t, balance.ColorIOTA, feeColor)
	require.Equal(t, newChainOwnerFee, ownerFee)
	require.Equal(t, newValidatorFee, validatorFee)
}

func Test_SetChainFees_TestCharge(t *testing.T) {
	notSolo := notsolo.New(t)

	// Generate a SigScheme for the chain originator
	chainOriginatorSigScheme := notSolo.SigScheme.NewSignatureSchemeWithFunds()
	notSolo.ValueTangle.RequireBalance(chainOriginatorSigScheme, balance.ColorIOTA, initialWalletFunds)

	// Generate a SigScheme for a wallet using the chain (will be charged for it)
	userSigScheme := notSolo.SigScheme.NewSignatureSchemeWithFunds()
	notSolo.ValueTangle.RequireBalance(userSigScheme, balance.ColorIOTA, initialWalletFunds)

	// Generate a dummy chain beloging to the chain originator
	chain := notSolo.Chain.NewChain(chainOriginatorSigScheme, "myChain")
	chainOriginatorBalanceInValueTangleAfterChainIsCreated := int64(initialWalletFunds - iotaTokensConsumedByRequest - iotaTokensConsumedByChain)
	chainOriginatorBalanceInChainAfterChainIsCreated := int64(iotaTokensConsumedByRequest)
	notSolo.ValueTangle.RequireBalance(chainOriginatorSigScheme, balance.ColorIOTA, chainOriginatorBalanceInValueTangleAfterChainIsCreated)
	notSolo.ValueTangle.RequireAddressBalance(chain.ChainAddress, chain.ChainColor, iotaTokensConsumedByChain)
	notSolo.Chain.RequireBalance(chain.OriginatorSigScheme, chain, balance.ColorIOTA, chainOriginatorBalanceInChainAfterChainIsCreated)

	// Initial chain fees
	feeColor, ownerFee, validatorFee := chain.GetFeeInfo(testconstants.AccountsContractName)
	require.Equal(t, balance.ColorIOTA, feeColor)
	require.Equal(t, int64(0), ownerFee)
	require.Equal(t, int64(0), validatorFee)

	// Request to chain change fee settings
	const newOwnerFee = int64(100)

	notSolo.Chain.ChangeContractFees(chain.OriginatorSigScheme, chain, testconstants.AccountsContractName, newOwnerFee)
	chainOriginatorBalanceInValueTangleAfterFeeIsChanged := chainOriginatorBalanceInValueTangleAfterChainIsCreated - iotaTokensConsumedByRequest
	chainOriginatorBalanceInChainAfterFeeIsChanged := chainOriginatorBalanceInChainAfterChainIsCreated + iotaTokensConsumedByRequest

	// Chain fees after change
	feeColor, ownerFee, validatorFee = chain.GetFeeInfo(testconstants.AccountsContractName)
	require.Equal(t, balance.ColorIOTA, feeColor)
	require.Equal(t, newOwnerFee, ownerFee)
	require.Equal(t, int64(0), validatorFee)

	// User sends a request to the chain (He wants to deposit iotas to his own account in the chain)
	notSolo.ValueTangle.TransferToChainToSelf(userSigScheme, chain, balance.ColorIOTA, transferValueIotas)

	// User gets the iota tokens in his account in the chain
	notSolo.ValueTangle.RequireBalance(userSigScheme, balance.ColorIOTA, initialWalletFunds-transferValueIotas-iotaTokensConsumedByRequest) // Transfered tokens are debited from the value tangle
	// Fee is charged from the funds transfered with the request
	notSolo.Chain.RequireBalance(userSigScheme, chain, balance.ColorIOTA, transferValueIotas+iotaTokensConsumedByRequest-ownerFee) // His tokens in the chain minus fees

	// Chain owner gets fees from the operation
	notSolo.ValueTangle.RequireBalance(chainOriginatorSigScheme, balance.ColorIOTA, chainOriginatorBalanceInValueTangleAfterFeeIsChanged)     // No change in value tangle
	notSolo.Chain.RequireBalance(chainOriginatorSigScheme, chain, balance.ColorIOTA, chainOriginatorBalanceInChainAfterFeeIsChanged+ownerFee) // He gets the fees in the chain
}
