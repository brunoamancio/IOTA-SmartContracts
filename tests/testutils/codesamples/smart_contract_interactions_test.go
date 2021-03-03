package codesamples

import (
	"testing"

	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils"
	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils/testconstants"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/stretchr/testify/require"
)

func Test_DeploySmartContractIntoChain(t *testing.T) {
	env := solo.New(t, testconstants.Debug, testconstants.StackTrace)
	chain := env.NewChain(nil, "myChain")

	// Uploads wasm of SC and deploys it into chain
	contractWasmFilePath := testutils.MustGetContractWasmFilePath(t, testconstants.ContractName) // You can use if file is in SmartContract/pkg
	err := chain.DeployWasmContract(nil, testconstants.ContractName, contractWasmFilePath)
	require.NoError(t, err)

	// Loads contract information
	contract, err := chain.FindContract(testconstants.ContractName)
	require.NoError(t, err)
	require.NotNil(t, contract)
	require.Equal(t, testconstants.ContractName, contract.Name)
}

func Test_CallSmartContract_PostRequest(t *testing.T) {
	env := solo.New(t, testconstants.Debug, testconstants.StackTrace)
	chain := env.NewChain(nil, "myChain")

	// Uploads wasm of SC and deploys it into chain
	contractWasmFilePath := testutils.MustGetContractWasmFilePath(t, testconstants.ContractName) // You can use if file is in SmartContract/pkg
	err := chain.DeployWasmContract(nil, testconstants.ContractName, contractWasmFilePath)
	require.NoError(t, err)

	// Loads contract information
	contract, err := chain.FindContract(testconstants.ContractName)
	require.NoError(t, err)
	require.NotNil(t, contract)
	require.Equal(t, testconstants.ContractName, contract.Name)

	// Defines which contract and function will be called by chain.PostRequest
	const functionName = "my_sc_function"
	req := solo.NewCallParams(testconstants.ContractName, functionName)

	// Calls contract my_iota_sc, function my_sc_function
	_, err = chain.PostRequestSync(req, nil)
	require.NoError(t, err)
}

func Test_CallSmartContract_CallView(t *testing.T) {
	env := solo.New(t, testconstants.Debug, testconstants.StackTrace)
	chain := env.NewChain(nil, "myChain")

	// Uploads wasm of SC and deploys it into chain
	contractWasmFilePath := testutils.MustGetContractWasmFilePath(t, testconstants.ContractName) // You can use if file is in SmartContract/pkg
	err := chain.DeployWasmContract(nil, testconstants.ContractName, contractWasmFilePath)
	require.NoError(t, err)

	// Loads contract information
	contract, err := chain.FindContract(testconstants.ContractName)
	require.NoError(t, err)
	require.NotNil(t, contract)
	require.Equal(t, testconstants.ContractName, contract.Name)

	// Defines which contract and function will be called by chain.PostRequest
	const functionName = "my_sc_view"

	// Calls contract my_iota_sc, function my_sc_view
	result, err := chain.CallView(testconstants.ContractName, functionName)
	require.NoError(t, err)
	require.NotNil(t, result)
}
