package testconstants

import (
	"github.com/brunoamancio/NotSolo/constants"
)

const (
	// ContractName is defined in smartcontract/rust/Cargo.toml > package > name
	ContractName = "my_iota_sc"
	// Debug is used by Solo. 'true' for logging level 'debug', otherwise 'info'
	Debug = false
	// StackTrace is used by Solo. 'true' if stack trace must be printed in case of errors
	StackTrace = false

	/* INTERESTING FACT: Calls to a smart contract require 1 EXTRA iota token to be sent to the chain it is located in.
	   It is colored with the chain's color and corresponds to the request. That is how the protocol locates the backlog of
	   requests to be processed. Basically, it works as a flag. After the request is processed, the token is uncolored
	   and sent to the chain owner's account in the chain.
	*/
	IotaTokensConsumedByRequest = constants.IotaTokensConsumedByRequest

	// Used to fund address in NewKeyPairWithFunds. // Defined in iotaledger/wasp/packages/testutiltestutil.RequestFundsAmount.
	InitialWalletFunds = uint64(1_000_000)

	// AccountsContractName sets the name of the Accounts contract, which is a root contract present in every chain
	AccountsContractName = "accounts"
	BlobContractName     = "blob"
)
