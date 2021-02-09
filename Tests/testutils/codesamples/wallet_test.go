package codesamples

import (
	"testing"

	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/testutil"
	"github.com/stretchr/testify/require"
)

// You can generate wallets to use them in your unit tests
func Test_GenerateWalletWithDummyFunds(t *testing.T) {
	env := solo.New(t, false, false)

	// Generates a key pair for a wallet and provides it with dummy funds.
	// The amount is defined in Wasp (constant testutil.RequestFundsAmount) and WaspConn plug-in (constant utxodb.RequestFundsAmount)
	walletKeyPair := env.NewSignatureSchemeWithFunds()
	require.NotNil(t, walletKeyPair)

	// Wallet address can be used to send and receive funds in your tests
	walletAddress := walletKeyPair.Address()
	require.NotNil(t, walletAddress)

	// Wallet balance
	require.NotNil(t, walletAddress)
	env.AssertAddressBalance(walletAddress, balance.ColorIOTA, testutil.RequestFundsAmount)
}
