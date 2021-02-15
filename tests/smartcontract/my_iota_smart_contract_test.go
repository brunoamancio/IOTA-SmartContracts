package libtest

import (
	"testing"

	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils"
	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils/testconstants"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/stretchr/testify/require"
)

//  -----------------------------------------------  //
//  See code samples in Tests/testutils/codesamples  //
//  -----------------------------------------------  //

func TestLib(t *testing.T) {
	contractWasmFilePath := testutils.MustGetContractWasmFilePath(t, testconstants.ContractName) // You can use if file is in SmartContract/pkg

	// Name of the SC function to be requested - Defined in lib.rs > add_call > my_sc_function
	functionName := "my_sc_function"

	env := solo.New(t, testconstants.Debug, testconstants.StackTrace)
	chainName := testconstants.ContractName + "Chain"
	chain := env.NewChain(nil, chainName)

	// Uploads wasm of SC and deploys it into chain
	err := chain.DeployWasmContract(nil, testconstants.ContractName, contractWasmFilePath)
	require.NoError(t, err)

	// Defines which contract and function will be called by chain.PostRequest
	req := solo.NewCallParams(testconstants.ContractName, functionName)

	// Calls contract my_iota_sc, function my_sc_function
	_, err = chain.PostRequest(req, nil)
	require.NoError(t, err)
}
