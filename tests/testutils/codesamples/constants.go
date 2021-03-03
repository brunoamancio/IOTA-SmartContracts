package codesamples

// Used to fund address in NewSignatureSchemeWithFunds. // Defined in iotaledger/wasp/packages/testutiltestutil.RequestFundsAmount.
const initialWalletFunds = 1337

// Default amount of IOTAs to transfer in unit tests.
const transferValueIotas = 1000

/* INTERESTING FACT: Calls to a smart contract require 1 EXTRA iota token to be sent to the chain it is located in.
It is colored with the chain's color and corresponds to the request. That is how the protocol locates the backlog of
requests to be processed. Basically, it works as a flag. After the request is processed, the token is uncolored
and sent to the chain owner's account in the chain.
*/
const iotaTokensConsumedByRequest = 1

/* INTERESTING FACT: Creating a chain requires 2 iota tokens. They are colored with the chain's color,
1 is sent to the chain's address in the value tangle, and the other is used exactly as iotaTokensConsumedByRequest.
*/
const iotaTokensConsumedByChain = 1
