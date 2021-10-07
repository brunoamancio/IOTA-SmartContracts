package codesamples

import (
	"testing"

	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils"
	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils/testconstants"
	notsolo "github.com/brunoamancio/NotSolo"
	"github.com/iotaledger/wasp/packages/iscp/colored"
	"github.com/stretchr/testify/require"
)

/// This is a sample of how to send tokens from the value tangle to a chain, keeping the ownership of the tokens on that chain
func Test_SendTokensToChain_NoContractFees(t *testing.T) {
	notSolo := notsolo.New(t)

	require.GreaterOrEqual(t, initialWalletFunds, transferValueIotas)

	// Generates key pairs for sender and receiver wallets and provides both with dummy funds.
	senderWalletKeyPair := notSolo.KeyPair.NewKeyPairWithFunds()

	// Wallet balance in value tangle -> Before transfer
	notSolo.L1.RequireBalance(senderWalletKeyPair, colored.IOTA, initialWalletFunds)

	// Generate a dummy chain NOT beloging to sender
	chain := notSolo.Chain.NewChain(nil, "myChain")

	// Wallet balance in chain -> Before transfer
	notSolo.Chain.RequireBalance(senderWalletKeyPair, chain, colored.IOTA, 0)

	// Transfer from sender's address in the value tangle to his account in 'chain'
	notSolo.L1.MustTransferToChainToSelf(senderWalletKeyPair, chain, colored.IOTA, transferValueIotas)

	// Wallet balances -> After transfer to chain
	notSolo.L1.RequireBalance(senderWalletKeyPair, colored.IOTA, initialWalletFunds-transferValueIotas)
	notSolo.Chain.RequireBalance(senderWalletKeyPair, chain, colored.IOTA, transferValueIotas)
}

func Test_SendAndReceiveTokens_NoContractFees(t *testing.T) {
	notSolo := notsolo.New(t)

	require.GreaterOrEqual(t, initialWalletFunds, transferValueIotas)

	// Generates key pairs for sender and receiver wallets and provides both with dummy funds.
	senderWalletKeyPair := notSolo.KeyPair.NewKeyPairWithFunds()

	// Wallet balance in value tangle -> Before transfer
	notSolo.L1.RequireBalance(senderWalletKeyPair, colored.IOTA, initialWalletFunds)

	// Generates key pairs for sender wallet.
	receiverWalletKeyPair := notSolo.KeyPair.NewKeyPair()
	notSolo.L1.RequireBalance(receiverWalletKeyPair, colored.IOTA, 0)

	// Generate a dummy chain NEITHER belonging to sender NOR receiver
	chain := notSolo.Chain.NewChain(nil, "myChain")

	// Transfer from sender's address in the value tangle to the receiver's account in 'chain'
	notSolo.L1.MustTransferToChain(senderWalletKeyPair, chain, colored.IOTA, transferValueIotas, receiverWalletKeyPair)

	// Wallet balances -> After transfer
	notSolo.L1.RequireBalance(senderWalletKeyPair, colored.IOTA, initialWalletFunds-transferValueIotas)
	notSolo.Chain.RequireBalance(senderWalletKeyPair, chain, colored.IOTA, 0)

	notSolo.L1.RequireBalance(receiverWalletKeyPair, colored.IOTA, 0)
	notSolo.Chain.RequireBalance(receiverWalletKeyPair, chain, colored.IOTA, transferValueIotas)
}

func Test_SendTokensToChain_WithContractFees(t *testing.T) {
	notSolo := notsolo.New(t)

	require.GreaterOrEqual(t, initialWalletFunds, transferValueIotas)

	// Generates key pairs for sender and receiver wallets and provides both with dummy funds.
	senderWalletKeyPair := notSolo.KeyPair.NewKeyPairWithFunds()

	// Wallet balance in value tangle -> Before transfer
	notSolo.L1.RequireBalance(senderWalletKeyPair, colored.IOTA, initialWalletFunds)

	// Generate a dummy chain NOT beloging to sender
	chain := notSolo.Chain.NewChain(nil, "myChain")
	contractName := "accounts" // This is a root contract, present in every chain

	// Check contract fees before changing them
	feeColor, ownerFee, validatorFee := chain.GetFeeInfo(contractName)
	require.Equal(t, colored.IOTA, feeColor)
	require.Equal(t, uint64(0), ownerFee)
	require.Equal(t, uint64(0), validatorFee)

	// Request to change the contract fee settings in 'chain'
	const newContractOwnerFee = uint64(2)
	notSolo.Chain.ChangeContractFees(chain.OriginatorKeyPair, chain, contractName, newContractOwnerFee)

	// Check contract fees after changing them
	feeColor, ownerFee, validatorFee = chain.GetFeeInfo(contractName)
	require.Equal(t, colored.IOTA, feeColor)
	require.Equal(t, newContractOwnerFee, ownerFee)
	require.Equal(t, uint64(0), validatorFee)

	// Wallet balance in chain -> Before transfer
	notSolo.Chain.RequireBalance(senderWalletKeyPair, chain, colored.IOTA, 0)

	// Transfer from sender's address in the value tangle to his account in 'chain'
	notSolo.L1.MustTransferToChainToSelf(senderWalletKeyPair, chain, colored.IOTA, transferValueIotas)

	// Balances -> After transfer to chain
	notSolo.L1.RequireBalance(senderWalletKeyPair, colored.IOTA, initialWalletFunds-transferValueIotas)            // Transfered tokens are debited from the value tangle
	notSolo.Chain.RequireBalance(senderWalletKeyPair, chain, colored.IOTA, transferValueIotas-newContractOwnerFee) // His tokens in the chain minus fees
}

func Test_SendTokensToContract_NoContractFees(t *testing.T) {
	notSolo := notsolo.New(t)

	// Generate a dummy chain NOT beloging to sender
	chain := notSolo.Chain.NewChain(nil, "myChain")

	// Uploads wasm of SC and deploys it into chain
	contractWasmFilePath := testutils.MustGetContractWasmFilePath(t, testconstants.ContractName) // You can use if file is in SmartContract/pkg
	notSolo.Chain.DeployWasmContract(chain, nil, testconstants.ContractName, contractWasmFilePath)

	// Generates key pairs for sender wallets, which will send iota tokens to the contract
	senderWalletKeyPair := notSolo.KeyPair.NewKeyPairWithFunds()

	// Balance in value tangle -> Before transfer
	notSolo.L1.RequireBalance(senderWalletKeyPair, colored.IOTA, initialWalletFunds)
	require.GreaterOrEqual(t, initialWalletFunds, transferValueIotas)

	// Transfer from sender's address in the value tangle to the contract's account in 'chain'
	notSolo.L1.MustTransferToContract(senderWalletKeyPair, chain, colored.IOTA, transferValueIotas, testconstants.ContractName)

	// Balances -> After transfer
	notSolo.L1.RequireBalance(senderWalletKeyPair, colored.IOTA, initialWalletFunds-transferValueIotas)
	notSolo.Chain.RequireBalance(senderWalletKeyPair, chain, colored.IOTA, 0)

	// Contract's account balance in the chain -> After transfer
	notSolo.Chain.RequireContractBalance(chain, testconstants.ContractName, colored.IOTA, transferValueIotas)
}
