package codesamples

import (
	"testing"

	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils"
	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils/testconstants"
	notsolo "github.com/brunoamancio/NotSolo"
	"github.com/stretchr/testify/require"
)

func Test_DeploySmartContractIntoChain(t *testing.T) {
	notSolo := notsolo.New(t)
	chain := notSolo.Chain.NewChain(nil, "myChain")

	// Uploads wasm of SC and deploys it into chain
	contractWasmFilePath := testutils.MustGetContractWasmFilePath(t, testconstants.ContractName) // You can use if file is in SmartContract/pkg
	notSolo.Chain.DeployWasmContract(chain, nil, testconstants.ContractName, contractWasmFilePath)

	// Loads contract information
	contract := notSolo.Chain.MustGetContractRecord(chain, testconstants.ContractName)
	require.Equal(t, testconstants.ContractName, contract.Name)
}

func Test_CallSmartContract_PostRequest(t *testing.T) {
	notSolo := notsolo.New(t)
	chain := notSolo.Chain.NewChain(nil, "myChain")

	// Uploads wasm of SC and deploys it into chain
	contractWasmFilePath := testutils.MustGetContractWasmFilePath(t, testconstants.ContractName) // You can use if file is in SmartContract/pkg
	notSolo.Chain.DeployWasmContract(chain, nil, testconstants.ContractName, contractWasmFilePath)

	// Loads contract information
	contract := notSolo.Chain.MustGetContractRecord(chain, testconstants.ContractName)
	require.Equal(t, testconstants.ContractName, contract.Name)

	// Defines which contract and function will be called by chain.PostRequest
	const functionName = "my_sc_function"

	// Calls contract my_iota_sc, function my_sc_function
	notSolo.Request.MustPost(nil, chain, testconstants.ContractName, functionName)
}

func Test_CallSmartContract_CallView(t *testing.T) {
	notSolo := notsolo.New(t)
	chain := notSolo.Chain.NewChain(nil, "myChain")

	// Uploads wasm of SC and deploys it into chain
	contractWasmFilePath := testutils.MustGetContractWasmFilePath(t, testconstants.ContractName) // You can use if file is in SmartContract/pkg
	notSolo.Chain.DeployWasmContract(chain, nil, testconstants.ContractName, contractWasmFilePath)

	// Loads contract information
	contract := notSolo.Chain.MustGetContractRecord(chain, testconstants.ContractName)
	require.Equal(t, testconstants.ContractName, contract.Name)

	// Defines which contract and function will be called by chain.PostRequest
	const functionName = "my_sc_view"

	// Calls contract my_iota_sc, function my_sc_view
	result := notSolo.Request.MustView(chain, testconstants.ContractName, functionName)
	require.NotNil(t, result)
}
