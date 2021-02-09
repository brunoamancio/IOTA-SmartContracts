package codesamples

import (
	"testing"

	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/accounts"
	"github.com/iotaledger/wasp/packages/vm/core/root"
	"github.com/stretchr/testify/require"
)

/// This is a sample of how to send tokens from the value tangle to a chain, keeping the ownership of the tokens on that chain
func Test_SendTokensToChain_NoContractFees(t *testing.T) {
	env := solo.New(t, false, false)

	require.GreaterOrEqual(t, initialWalletFunds, transferValueIotas)

	// Generates key pairs for sender and receiver wallets and provides both with dummy funds.
	senderWalletKeyPair := env.NewSignatureSchemeWithFunds()
	senderWalletAddress := senderWalletKeyPair.Address()
	senderAgentID := coretypes.NewAgentIDFromAddress(senderWalletAddress)
	require.NotNil(t, senderWalletKeyPair)
	require.NotNil(t, senderWalletAddress)
	require.NotNil(t, senderAgentID)

	// Wallet balance in value tangle -> Before transfer
	env.AssertAddressBalance(senderWalletAddress, balance.ColorIOTA, initialWalletFunds)

	// Generate a dummy chain NOT beloging to sender
	chain := env.NewChain(nil, "myChain")

	// Wallet balance in chain -> Before transfer
	chain.AssertAccountBalance(senderAgentID, balance.ColorIOTA, 0)

	// Transfer from value tangle to the chain
	transferRequest := solo.NewCallParams(accounts.Name, accounts.FuncDeposit).WithTransfer(balance.ColorIOTA, transferValueIotas)
	_, err := chain.PostRequest(transferRequest, senderWalletKeyPair)
	require.NoError(t, err)

	// Wallet balances -> After transfer to chain
	env.AssertAddressBalance(senderWalletAddress, balance.ColorIOTA, initialWalletFunds-transferValueIotas-iotaTokensConsumedByRequest)
	chain.AssertAccountBalance(senderAgentID, balance.ColorIOTA, transferValueIotas+iotaTokensConsumedByRequest)
}

func Test_SendAndReceiveTokens_NoContractFees(t *testing.T) {
	env := solo.New(t, false, false)

	require.GreaterOrEqual(t, initialWalletFunds, transferValueIotas)

	// Generates key pairs for sender wallet and provides it with dummy funds.
	senderWalletKeyPair := env.NewSignatureSchemeWithFunds()
	senderWalletAddress := senderWalletKeyPair.Address()
	senderWalletAgentID := coretypes.NewAgentIDFromAddress(senderWalletAddress)
	require.NotNil(t, senderWalletKeyPair)
	require.NotNil(t, senderWalletAddress)
	env.AssertAddressBalance(senderWalletAddress, balance.ColorIOTA, initialWalletFunds)

	// Generates key pairs for sender wallet.
	receivedWalletKeyPair := env.NewSignatureScheme()
	receiverWalletAddress := receivedWalletKeyPair.Address()
	receiverWalletAgentID := coretypes.NewAgentIDFromAddress(receiverWalletAddress)
	require.NotNil(t, receivedWalletKeyPair)
	require.NotNil(t, receiverWalletAddress)
	require.NotNil(t, receiverWalletAgentID)
	env.AssertAddressBalance(receiverWalletAddress, balance.ColorIOTA, 0)

	// Generate a dummy chain NEITHER belonging to sender NOR receiver
	chain := env.NewChain(nil, "myChain")

	// Transfer within the chain
	transferRequest := solo.NewCallParams(accounts.Name, accounts.FuncDeposit, accounts.ParamAgentID, codec.EncodeAgentID(receiverWalletAgentID)).
		WithTransfer(balance.ColorIOTA, transferValueIotas)

	_, err := chain.PostRequest(transferRequest, senderWalletKeyPair)
	require.NoError(t, err)

	// Wallet balances -> After transfer
	env.AssertAddressBalance(senderWalletAddress, balance.ColorIOTA, initialWalletFunds-transferValueIotas-iotaTokensConsumedByRequest)
	chain.AssertAccountBalance(senderWalletAgentID, balance.ColorIOTA, iotaTokensConsumedByRequest)

	env.AssertAddressBalance(receiverWalletAddress, balance.ColorIOTA, 0)
	chain.AssertAccountBalance(receiverWalletAgentID, balance.ColorIOTA, transferValueIotas)
}

func Test_SendTokensToChain_WithContractFees(t *testing.T) {
	env := solo.New(t, false, false)

	require.GreaterOrEqual(t, initialWalletFunds, transferValueIotas)

	// Generates key pairs for sender and receiver wallets and provides both with dummy funds.
	senderWalletKeyPair := env.NewSignatureSchemeWithFunds()
	senderWalletAddress := senderWalletKeyPair.Address()
	senderAgentID := coretypes.NewAgentIDFromAddress(senderWalletAddress)
	require.NotNil(t, senderWalletKeyPair)
	require.NotNil(t, senderWalletAddress)
	require.NotNil(t, senderAgentID)

	// Wallet balance in value tangle -> Before transfer
	env.AssertAddressBalance(senderWalletAddress, balance.ColorIOTA, initialWalletFunds)

	// Generate a dummy chain NOT beloging to sender
	chain := env.NewChain(nil, "myChain")
	feeColor, ownerFee, validatorFee := chain.GetFeeInfo(accounts.Name)
	require.Equal(t, balance.ColorIOTA, feeColor)
	require.Equal(t, int64(0), ownerFee)
	require.Equal(t, int64(0), validatorFee)

	// Request to chain change fee settings
	const newOwnerFee = int64(100)
	transferRequest := solo.NewCallParams(root.Interface.Name, root.FuncSetContractFee, root.ParamHname, accounts.Interface.Hname(), root.ParamOwnerFee, newOwnerFee)
	_, err := chain.PostRequest(transferRequest, chain.OriginatorSigScheme)
	require.NoError(t, err)

	// Wallet balance in chain -> Before transfer
	chain.AssertAccountBalance(senderAgentID, balance.ColorIOTA, 0)

	// Transfer from value tangle to the chain
	transferRequest = solo.NewCallParams(accounts.Name, accounts.FuncDeposit).WithTransfer(balance.ColorIOTA, transferValueIotas)
	_, err = chain.PostRequest(transferRequest, senderWalletKeyPair)
	require.NoError(t, err)

	// Wallet balances -> After transfer to chain
	env.AssertAddressBalance(senderWalletAddress, balance.ColorIOTA, initialWalletFunds-transferValueIotas-iotaTokensConsumedByRequest) // Transfered tokens are debited from the value tangle
	chain.AssertAccountBalance(senderAgentID, balance.ColorIOTA, transferValueIotas+iotaTokensConsumedByRequest-newOwnerFee)            // His tokens in the chain minus fees
}

func Test_SendTokensToContract_NoContractFees(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "myChain")

	// Uploads wasm of SC and deploys it into chain
	const contractName = "my_iota_sc"
	contractWasmFilePath := "<file path to contract.wasm>" // You can use if file is in SmartContract/pkg testutils.MustGetContractWasmFilePath(contractName)
	err := chain.DeployWasmContract(nil, contractName, contractWasmFilePath)
	require.NoError(t, err)

	// Loads contract information
	contract, err := chain.FindContract(contractName)
	require.NoError(t, err)
	contractID := coretypes.NewContractID(chain.ChainID, contract.Hname())
	contractAgentID := coretypes.NewAgentIDFromContractID(contractID)

	// Generates key pairs for sender wallets, which will send iota tokens to the contract
	senderWalletKeyPair := env.NewSignatureSchemeWithFunds()
	senderWalletAddress := senderWalletKeyPair.Address()
	senderAgentID := coretypes.NewAgentIDFromAddress(senderWalletAddress)
	require.NotNil(t, senderWalletKeyPair)
	require.NotNil(t, senderWalletAddress)
	require.NotNil(t, senderAgentID)

	// Wallet balance in value tangle -> Before transfer
	env.AssertAddressBalance(senderWalletAddress, balance.ColorIOTA, initialWalletFunds)
	require.GreaterOrEqual(t, initialWalletFunds, transferValueIotas)

	// Transfer from value tangle to the contract (in the chain)
	transferRequest := solo.NewCallParams(accounts.Name, accounts.FuncDeposit, accounts.ParamAgentID, contractAgentID).WithTransfer(balance.ColorIOTA, transferValueIotas)
	_, err = chain.PostRequest(transferRequest, senderWalletKeyPair)
	require.NoError(t, err)

	// Wallet balances -> After transfer
	env.AssertAddressBalance(senderWalletAddress, balance.ColorIOTA, initialWalletFunds-transferValueIotas-iotaTokensConsumedByRequest)
	chain.AssertAccountBalance(senderAgentID, balance.ColorIOTA, iotaTokensConsumedByRequest)

	// Contract account balance in the chain -> After transfer
	chain.AssertAccountBalance(contractAgentID, balance.ColorIOTA, transferValueIotas)
}
