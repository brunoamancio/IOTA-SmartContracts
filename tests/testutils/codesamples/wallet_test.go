package codesamples

import (
	"testing"

	notsolo "github.com/brunoamancio/NotSolo"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
)

// You can generate wallets (signature schemes which are key pairs) to use them in your unit tests
func Test_GenerateWalletWithDummyFunds(t *testing.T) {
	notSolo := notsolo.New(t)

	// Generates a key pair for a wallet and provides it with dummy funds.
	// The amount is defined in Wasp (constant testutil.RequestFundsAmount) and WaspConn plug-in (constant utxodb.RequestFundsAmount)
	walletSigScheme := notSolo.SigScheme.NewSignatureSchemeWithFunds()

	// Uses the walletSigScheme to get the wallet address and ensures the balance is as expected
	notSolo.ValueTangle.RequireBalance(walletSigScheme, balance.ColorIOTA, initialWalletFunds)
}
