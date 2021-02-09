package libtest

import (
	"testing"

	"github.com/iotaledger/wasp/packages/solo"
	"github.com/stretchr/testify/require"
)

//  -----------------------------------------------  //
//  See code samples in Tests/testutils/codesamples  //
//  -----------------------------------------------  //

func TestLib(t *testing.T) {
	// Contract name - Defined in SmartContract/Cargo.toml > package > name
	const contractName = "my_iota_sc"
	contractWasmFilePath := "<file path to contract.wasm>" // You can use if file is in SmartContract/pkg testutils.MustGetContractWasmFilePath(t, contractName)

	// Name of the SC function to be requested - Defined in lib.rs > add_call > my_sc_request
	functionName := "my_sc_request"

	env := solo.New(t, false, false)
	chainName := contractName + "Chain"
	chain := env.NewChain(nil, chainName)

	// Uploads wasm of SC and deploys it into chain
	err := chain.DeployWasmContract(nil, contractName, contractWasmFilePath)
	require.NoError(t, err)

	// Defines which contract and function will be called by chain.PostRequest
	req := solo.NewCallParams(contractName, functionName)

	// Calls contract my_iota_sc, function my_sc_request
	_, err = chain.PostRequest(req, nil)
	require.NoError(t, err)
}
