package codesamples

import (
	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils/testconstants"
	"github.com/brunoamancio/NotSolo/constants"
)

const initialWalletFunds = testconstants.InitialWalletFunds

// Default amount of IOTAs to transfer in unit tests.
const transferValueIotas = uint64(100)

const iotaTokensConsumedByRequest = testconstants.IotaTokensConsumedByRequest

const chainAddressBalanceInL1OnChainCreated = constants.DefaultChainStartingBalance + iotaTokensConsumedByRequest
