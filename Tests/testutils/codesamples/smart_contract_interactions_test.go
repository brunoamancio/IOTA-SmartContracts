package codesamples

import (
	"testing"

	"github.com/iotaledger/wasp/packages/solo"
	"github.com/stretchr/testify/require"
)

func Test_DeploySmartContractIntoChain(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "myChain")

	// Uploads wasm of SC and deploys it into chain
	const contractName = "my_iota_sc"
	const contractWasmFilePath = "<file path to contract.wasm>" // You can use if file is in SmartContract/pkg testutils.MustGetContractWasmFilePath(t, contractName)
	err := chain.DeployWasmContract(nil, contractName, contractWasmFilePath)
	require.NoError(t, err)

	// Loads contract information
	contract, err := chain.FindContract(contractName)
	require.NoError(t, err)
	require.NotNil(t, contract)
	require.Equal(t, contractName, contract.Name)
}

func Test_CallSmartContract_PostRequest(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "myChain")

	// Uploads wasm of SC and deploys it into chain
	const contractName = "my_iota_sc"
	const contractWasmFilePath = "<file path to contract.wasm>" // You can use if file is in SmartContract/pkg testutils.MustGetContractWasmFilePath(t, contractName)
	err := chain.DeployWasmContract(nil, contractName, contractWasmFilePath)
	require.NoError(t, err)

	// Loads contract information
	contract, err := chain.FindContract(contractName)
	require.NoError(t, err)
	require.NotNil(t, contract)
	require.Equal(t, contractName, contract.Name)

	// Defines which contract and function will be called by chain.PostRequest
	const functionName = "my_sc_request"
	req := solo.NewCallParams(contractName, functionName)

	// Calls contract my_iota_sc, function my_sc_request
	_, err = chain.PostRequest(req, nil)
	require.NoError(t, err)
}

func Test_CallSmartContract_CallView(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "myChain")

	// Uploads wasm of SC and deploys it into chain
	const contractName = "my_iota_sc"
	const contractWasmFilePath = "<file path to contract.wasm>" // You can use if file is in SmartContract/pkg testutils.MustGetContractWasmFilePath(t, contractName)
	err := chain.DeployWasmContract(nil, contractName, contractWasmFilePath)
	require.NoError(t, err)

	// Loads contract information
	contract, err := chain.FindContract(contractName)
	require.NoError(t, err)
	require.NotNil(t, contract)
	require.Equal(t, contractName, contract.Name)

	// Defines which contract and function will be called by chain.PostRequest
	const functionName = "my_sc_view"

	// Calls contract my_iota_sc, function my_sc_view
	result, err := chain.CallView(contractName, functionName)
	require.NoError(t, err)
	require.NotNil(t, result)
}
