package codesamples

import (
	"testing"

	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/accounts"
	"github.com/iotaledger/wasp/packages/vm/core/root"
	"github.com/stretchr/testify/require"
)

func Test_CreateChain_chainCreatorSpecified(t *testing.T) {
	env := solo.New(t, false, false)

	// Create address with dummy tokens in it.
	chainOriginatorsWalletKeyPair := env.NewSignatureSchemeWithFunds()
	require.NotNil(t, chainOriginatorsWalletKeyPair)

	// Wallet addresses
	chainOriginatorsWalletAddress := chainOriginatorsWalletKeyPair.Address()
	require.NotNil(t, chainOriginatorsWalletAddress)

	// Wallet balance in value tangle before chain is created
	env.AssertAddressBalance(chainOriginatorsWalletAddress, balance.ColorIOTA, initialWalletFunds)

	// Create a chain where chainCreatorsWalletKeyPair is the owner.
	chain := env.NewChain(chainOriginatorsWalletKeyPair, "myChain")
	require.NotNil(t, chain)
	require.NotEqual(t, chain.ChainColor, balance.ColorIOTA)

	// Wallet balance in value tangle after chain is created
	env.AssertAddressBalance(chainOriginatorsWalletAddress, balance.ColorIOTA, initialWalletFunds-iotaTokensConsumedByRequest-iotaTokensConsumedByChain)

	// IMPORTANT: When a chain is created, 1 IOTA is colored with the chain's color and sent to the chain's address in the value tangle
	env.AssertAddressBalance(chain.ChainAddress, chain.ChainColor, iotaTokensConsumedByChain)

	// AgentID of the wallet (also, chain.OriginatorAgentID)
	chainOriginatorsAgentID := coretypes.NewAgentIDFromAddress(chainOriginatorsWalletAddress)

	// IMPORTANT: When a chain is created, 1 IOTA is sent from the originator's account in the value tangle their account in the chain
	chain.AssertAccountBalance(chainOriginatorsAgentID, balance.ColorIOTA, iotaTokensConsumedByRequest)
}

// Sample of how to create chain without specifying a chainOriginator.
// A dummy chain originator is created in the background (by NewChain).
func Test_CreateChain_NoChainCreatorSpecified(t *testing.T) {
	env := solo.New(t, false, false)

	// Create a chain where chainCreatorsWalletKeyPair is the owner.
	chain := env.NewChain(nil, "myChain")
	require.NotNil(t, chain)
	require.NotEqual(t, chain.ChainColor, balance.ColorIOTA)

	// IMPORTANT: When a chain is created, 1 IOTA is colored with the chain's color and sent to the chain's address in the value tangle
	env.AssertAddressBalance(chain.ChainAddress, chain.ChainColor, iotaTokensConsumedByChain)

	// IMPORTANT: When a chain is created, 1 IOTA is sent from the originator's account in the value tangle their account in the chain
	chain.AssertAccountBalance(chain.OriginatorAgentID, balance.ColorIOTA, iotaTokensConsumedByRequest)
}

func Test_SetChainFees(t *testing.T) {
	env := solo.New(t, false, false)

	// Generate a dummy chain beloging to chain.OriginatorSigScheme
	chain := env.NewChain(nil, "myChain")

	// Initial chain fees
	feeColor, ownerFee, validatorFee := chain.GetFeeInfo(accounts.Name)
	require.Equal(t, balance.ColorIOTA, feeColor)
	require.Equal(t, int64(0), ownerFee)
	require.Equal(t, int64(0), validatorFee)

	// Request to chain - change owner fee settings
	const newChainOwnerFee = int64(1000)
	transferRequest := solo.NewCallParams(root.Interface.Name, root.FuncSetContractFee, root.ParamHname, accounts.Interface.Hname(), root.ParamOwnerFee, newChainOwnerFee)
	_, err := chain.PostRequest(transferRequest, chain.OriginatorSigScheme)
	require.NoError(t, err)

	// Request to chain - change owner fee settings
	const newValidatorFee = int64(200)
	transferRequest = solo.NewCallParams(root.Interface.Name, root.FuncSetContractFee, root.ParamHname, accounts.Interface.Hname(), root.ParamValidatorFee, newValidatorFee)
	_, err = chain.PostRequest(transferRequest, chain.OriginatorSigScheme)
	require.NoError(t, err)

	// Chain fees after change
	feeColor, ownerFee, validatorFee = chain.GetFeeInfo(accounts.Name)
	require.Equal(t, balance.ColorIOTA, feeColor)
	require.Equal(t, newChainOwnerFee, ownerFee)
	require.Equal(t, newValidatorFee, validatorFee)
}

func Test_SetChainFees_TestCharge(t *testing.T) {
	env := solo.New(t, false, false)

	// Generate a keypair for the chain originator
	chainOriginatorKeyPair := env.NewSignatureSchemeWithFunds()
	chainOriginatorAddress := chainOriginatorKeyPair.Address()
	chainOriginatorAgentID := coretypes.NewAgentIDFromAddress(chainOriginatorAddress)
	require.NotNil(t, chainOriginatorKeyPair)
	env.AssertAddressBalance(chainOriginatorAddress, balance.ColorIOTA, initialWalletFunds)

	// Generate a keypair for a wallet using the chain (will be charged for it)
	userKeyPair := env.NewSignatureSchemeWithFunds()
	userAddress := userKeyPair.Address()
	userAgentID := coretypes.NewAgentIDFromAddress(userAddress)
	require.NotNil(t, userKeyPair)
	env.AssertAddressBalance(userAddress, balance.ColorIOTA, initialWalletFunds)

	// Generate a dummy chain beloging to the chain originator
	chain := env.NewChain(chainOriginatorKeyPair, "myChain")
	chainOriginatorBalanceInValueTangleAfterChainIsCreated := int64(initialWalletFunds - iotaTokensConsumedByRequest - iotaTokensConsumedByChain)
	chainOriginatorBalanceInChainAfterChainIsCreated := int64(iotaTokensConsumedByRequest)
	env.AssertAddressBalance(chainOriginatorAddress, balance.ColorIOTA, chainOriginatorBalanceInValueTangleAfterChainIsCreated)
	env.AssertAddressBalance(chain.ChainAddress, chain.ChainColor, iotaTokensConsumedByChain)
	chain.AssertAccountBalance(chainOriginatorAgentID, balance.ColorIOTA, chainOriginatorBalanceInChainAfterChainIsCreated)

	// Initial chain fees
	feeColor, ownerFee, validatorFee := chain.GetFeeInfo(accounts.Name)
	require.Equal(t, balance.ColorIOTA, feeColor)
	require.Equal(t, int64(0), ownerFee)
	require.Equal(t, int64(0), validatorFee)

	// Request to chain change fee settings
	const newOwnerFee = int64(100)
	transferRequest := solo.NewCallParams(root.Interface.Name, root.FuncSetContractFee, root.ParamHname, accounts.Interface.Hname(), root.ParamOwnerFee, newOwnerFee)
	_, err := chain.PostRequest(transferRequest, chainOriginatorKeyPair)
	require.NoError(t, err)
	chainOriginatorBalanceInValueTangleAfterFeeIsChanged := chainOriginatorBalanceInValueTangleAfterChainIsCreated - iotaTokensConsumedByRequest
	chainOriginatorBalanceInChainAfterFeeIsChanged := chainOriginatorBalanceInChainAfterChainIsCreated + iotaTokensConsumedByRequest

	// Chain fees after change
	feeColor, ownerFee, validatorFee = chain.GetFeeInfo(accounts.Name)
	require.Equal(t, balance.ColorIOTA, feeColor)
	require.Equal(t, newOwnerFee, ownerFee)
	require.Equal(t, int64(0), validatorFee)

	// User sends a request to the chain (He wants to deposit iotas to his own account in the chain)
	transferRequest = solo.NewCallParams(accounts.Name, accounts.FuncDeposit).WithTransfer(balance.ColorIOTA, transferValueIotas)
	_, err = chain.PostRequest(transferRequest, userKeyPair)
	require.NoError(t, err)

	// User gets the iota tokens in his account in the chain
	env.AssertAddressBalance(userAddress, balance.ColorIOTA, initialWalletFunds-transferValueIotas-iotaTokensConsumedByRequest) // Transfered tokens are debited from the value tangle
	// Fee is charged from the funds transfered with the request
	chain.AssertAccountBalance(userAgentID, balance.ColorIOTA, transferValueIotas+iotaTokensConsumedByRequest-ownerFee) // His tokens in the chain minus fees

	// Chain owner gets fees from the operation
	env.AssertAddressBalance(chainOriginatorAddress, balance.ColorIOTA, chainOriginatorBalanceInValueTangleAfterFeeIsChanged)      // No change in value tangle
	chain.AssertAccountBalance(chainOriginatorAgentID, balance.ColorIOTA, chainOriginatorBalanceInChainAfterFeeIsChanged+ownerFee) // He gets the fees in the chain
}
