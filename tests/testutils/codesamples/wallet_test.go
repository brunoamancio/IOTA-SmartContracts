package codesamples

import (
	"testing"

	notsolo "github.com/brunoamancio/NotSolo"
	"github.com/iotaledger/wasp/packages/iscp/colored"
)

// You can generate wallets (key pairs which are key pairs) to use them in your unit tests
func Test_GenerateWalletWithDummyFunds(t *testing.T) {
	notSolo := notsolo.New(t)

	// Generates a key pair for a wallet and provides it with dummy funds.
	// The amount is defined in Wasp (constant testutil.RequestFundsAmount) and WaspConn plug-in (constant utxodb.RequestFundsAmount)
	walletKeyPair := notSolo.KeyPair.NewKeyPairWithFunds()

	// Uses the walletKeyPair to get the wallet address and ensures the balance is as expected
	notSolo.L1.RequireBalance(walletKeyPair, colored.IOTA, initialWalletFunds)
}
