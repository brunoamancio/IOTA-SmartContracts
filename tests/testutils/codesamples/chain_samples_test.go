package codesamples

import (
	"testing"

	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils/testconstants"
	notsolo "github.com/brunoamancio/NotSolo"
	"github.com/brunoamancio/NotSolo/constants"
	"github.com/iotaledger/wasp/packages/iscp/colored"
	"github.com/iotaledger/wasp/packages/vm/core/blob"
	"github.com/stretchr/testify/require"
)

func Test_CreateChain_chainCreatorSpecified(t *testing.T) {
	notSolo := notsolo.New(t)

	// Create a key pair with dummy tokens in it.
	chainOriginatorsKeyPair := notSolo.KeyPair.NewKeyPairWithFunds()

	// Balance in value tangle before chain is created
	notSolo.L1.RequireBalance(chainOriginatorsKeyPair, colored.IOTA, initialWalletFunds)

	// Create a chain where chainOriginatorsKeyPair is the owner.
	chain := notSolo.Chain.NewChain(chainOriginatorsKeyPair, "myChain")

	// IMPORTANT: When a chain is created >>> USING SOLO <<<, a default amount of IOTA is sent to ChainID in L1
	// Another IOTA is consumed by the request and also sent to ChainID in L1
	expectedChainIdBalance := constants.DefaultChainStartingBalance + constants.IotaTokensConsumedByRequest
	notSolo.L1.RequireAddressBalance(chain.ChainID.AsAddress(), colored.IOTA, expectedChainIdBalance)

	// IMPORTANT: Originator has no balance in the chain
	notSolo.Chain.RequireBalance(chainOriginatorsKeyPair, chain, colored.IOTA, 0)

	// IMPORTANT: Originator has initial balance - the amount transfered from L1
	notSolo.L1.RequireBalance(chainOriginatorsKeyPair, colored.IOTA, initialWalletFunds-expectedChainIdBalance)
}

// Sample of how to create chain without specifying a chainOriginator.
// A dummy chain originator is created in the background (by NewChain).
func Test_CreateChain_NoChainCreatorSpecified(t *testing.T) {
	notSolo := notsolo.New(t)

	// Create a chain where chainOriginatorsKeyPair is the owner.
	chain := notSolo.Chain.NewChain(nil, "myChain")

	// IMPORTANT: When a chain is created >>> USING SOLO <<<, a default amount of IOTA is sent to ChainID in L1
	// Another IOTA is consumed by the request and also sent to ChainID in L1
	notSolo.L1.RequireAddressBalance(chain.ChainID.AsAddress(), colored.IOTA, chainAddressBalanceInL1OnChainCreated)

	// IMPORTANT: Originator has no balance in the chain
	notSolo.Chain.RequireBalance(chain.OriginatorKeyPair, chain, colored.IOTA, 0)

	// IMPORTANT: Originator has initial balance - the amount transfered from L1
	notSolo.L1.RequireBalance(chain.OriginatorKeyPair, colored.IOTA, initialWalletFunds-chainAddressBalanceInL1OnChainCreated)
}

func Test_SetChainFees(t *testing.T) {
	notSolo := notsolo.New(t)

	// Generate a dummy chain beloging to chain.OriginatorKeyPair
	chain := notSolo.Chain.NewChain(nil, "myChain")

	// Initial chain fees
	feeColor, ownerFee, validatorFee := chain.GetFeeInfo(testconstants.AccountsContractName)
	require.Equal(t, colored.IOTA, feeColor)
	require.Equal(t, uint64(0), ownerFee)
	require.Equal(t, uint64(0), validatorFee)

	// Request to chain - change contract owner fee settings
	const newChainOwnerFee = uint64(1000)
	notSolo.Chain.ChangeContractFees(chain.OriginatorKeyPair, chain, testconstants.AccountsContractName, newChainOwnerFee)

	// Request to chain - change validator fee settings
	const newValidatorFee = uint64(200)
	notSolo.Chain.ChangeValidatorFees(chain.OriginatorKeyPair, chain, testconstants.AccountsContractName, newValidatorFee)

	// Chain fees after change
	feeColor, ownerFee, validatorFee = chain.GetFeeInfo(testconstants.AccountsContractName)
	require.Equal(t, colored.IOTA, feeColor)
	require.Equal(t, newChainOwnerFee, ownerFee)
	require.Equal(t, newValidatorFee, validatorFee)
}

func Test_SetChainFees_TestCharge(t *testing.T) {
	notSolo := notsolo.New(t)

	// Generate a KeyPair for the chain originator
	chainOriginatorKeyPair := notSolo.KeyPair.NewKeyPairWithFunds()
	notSolo.L1.RequireBalance(chainOriginatorKeyPair, colored.IOTA, initialWalletFunds)

	// Generate a KeyPair for a wallet using the chain (will be charged for it)
	userKeyPair := notSolo.KeyPair.NewKeyPairWithFunds()
	notSolo.L1.RequireBalance(userKeyPair, colored.IOTA, initialWalletFunds)

	// Generate a dummy chain beloging to the chain originator
	chain := notSolo.Chain.NewChain(chainOriginatorKeyPair, "myChain")
	// Request cost is credited to chain's agentID: 'iotaTokensConsumedByRequest'
	expectedChainBalance := uint64(iotaTokensConsumedByRequest)
	notSolo.Chain.RequireChainBalance(chain, colored.IOTA, expectedChainBalance)

	chainOriginatorBalanceInL1AfterChainIsCreated := uint64(initialWalletFunds) - chainAddressBalanceInL1OnChainCreated
	chainOriginatorBalanceInChainAfterChainIsCreated := uint64(0)
	notSolo.L1.RequireBalance(chainOriginatorKeyPair, colored.IOTA, chainOriginatorBalanceInL1AfterChainIsCreated)
	notSolo.Chain.RequireBalance(chainOriginatorKeyPair, chain, colored.IOTA, chainOriginatorBalanceInChainAfterChainIsCreated)

	// Initial chain fees
	feeColor, ownerFee, validatorFee := chain.GetFeeInfo(testconstants.BlobContractName)
	require.Equal(t, colored.IOTA, feeColor)
	require.Equal(t, uint64(0), ownerFee)
	require.Equal(t, uint64(0), validatorFee)

	// Request to chain change fee settings
	const newOwnerFee = uint64(100)
	notSolo.Chain.ChangeContractFees(chainOriginatorKeyPair, chain, testconstants.BlobContractName, newOwnerFee)
	// IMPORTANT: calls to change contracts must send at least 1 token with the request (NotSolo sends "iotaTokensConsumedByRequest")
	chainOriginatorBalanceInL1AfterFeeIsChanged := chainOriginatorBalanceInL1AfterChainIsCreated - iotaTokensConsumedByRequest
	// Request cost is credited to chain's agentID: 'iotaTokensConsumedByRequest'
	expectedChainBalance += iotaTokensConsumedByRequest
	notSolo.L1.RequireBalance(chainOriginatorKeyPair, colored.IOTA, chainOriginatorBalanceInL1AfterFeeIsChanged)
	notSolo.Chain.RequireBalance(chainOriginatorKeyPair, chain, colored.IOTA, chainOriginatorBalanceInChainAfterChainIsCreated)

	notSolo.Chain.RequireChainBalance(chain, colored.IOTA, expectedChainBalance)

	// Chain fees after change
	feeColor, ownerFee, validatorFee = chain.GetFeeInfo(testconstants.BlobContractName)
	require.Equal(t, colored.IOTA, feeColor)
	require.Equal(t, newOwnerFee, ownerFee)
	require.Equal(t, uint64(0), validatorFee)

	// User sends a request to the contract (which needs to pay the fee)
	_, err := chain.UploadBlob(userKeyPair,
		blob.VarFieldVMType, "dummyType",
		blob.VarFieldProgramBinary, "dummyBinary",
	)
	require.NoError(t, err)

	// Fee is charged from the user's wallet
	notSolo.L1.RequireBalance(userKeyPair, colored.IOTA, initialWalletFunds-ownerFee) // His tokens minus fees
	notSolo.Chain.RequireBalance(userKeyPair, chain, colored.IOTA, 0)

	// Owner fee is credited to chain's agentID (upload request does not charge 'iotaTokensConsumedByRequest')
	expectedChainBalance += ownerFee
	notSolo.Chain.RequireChainBalance(chain, colored.IOTA, expectedChainBalance)

	// No change in the chain owner addresses neither in the chain nor in L1
	notSolo.L1.RequireBalance(chainOriginatorKeyPair, colored.IOTA, chainOriginatorBalanceInL1AfterFeeIsChanged)
	notSolo.Chain.RequireBalance(chainOriginatorKeyPair, chain, colored.IOTA, chainOriginatorBalanceInChainAfterChainIsCreated)

	// Chain originator harvests fees from chain to his account in the chain
	notSolo.Chain.Harvest(chain, colored.IOTA, newOwnerFee)
	// Request cost is credited to and hardvested amount is debited from chain's agentID
	expectedChainBalance = expectedChainBalance + iotaTokensConsumedByRequest - newOwnerFee

	notSolo.L1.RequireBalance(chainOriginatorKeyPair, colored.IOTA, chainOriginatorBalanceInL1AfterFeeIsChanged-iotaTokensConsumedByRequest)
	notSolo.Chain.RequireBalance(chainOriginatorKeyPair, chain, colored.IOTA, newOwnerFee)
	notSolo.Chain.RequireChainBalance(chain, colored.IOTA, expectedChainBalance)
}
