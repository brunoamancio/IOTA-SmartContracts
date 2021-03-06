package codesamples

import (
	"encoding/hex"
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

func Test_ViewMyBoolean(t *testing.T) {
	contractWasmFilePath := testutils.MustGetContractWasmFilePath(t, testconstants.ContractName) // You can use if file is in SmartContract/pkg

	// Name of the SC view to be requested - Defined in lib.rs > add_view > view_my_boolean
	functionName := "view_my_boolean"

	notSolo := notsolo.New(t)

	chainName := testconstants.ContractName + "Chain"
	chain := notSolo.Chain.NewChain(nil, chainName)

	// Uploads wasm of SC and deploys it into chain
	notSolo.Chain.DeployWasmContract(chain, nil, testconstants.ContractName, contractWasmFilePath)

	// Map test case and expected value
	caseToTest := make(map[string]bool)

	// hexadecimal number and the expected result
	caseToTest["01ffc9a7"] = true
	caseToTest["00000000"] = false
	caseToTest["11111111"] = false

	for hexadecimalString, expectedResult := range caseToTest {

		hexadecimalAsBytes, err := hex.DecodeString(hexadecimalString)
		require.NoError(t, err)

		t.Run(functionName, func(t *testing.T) {
			// Input parameter to sc view
			const ParamHexadecimal = "hexadecimal"
			// Output parameter from sc view
			const MatchesExpected = "matches_expected"

			// Call contract 'my_iota_sc', function 'view_my_boolean'
			response := notSolo.Request.MustView(chain, testconstants.ContractName, functionName, ParamHexadecimal, hexadecimalAsBytes)

			// Get output parameter MatchesExpected
			viewMyBooleanResponse := notSolo.Data.MustGetBool(response[MatchesExpected])

			// Ensure it matches in the put
			require.Equal(t, expectedResult, viewMyBooleanResponse)
		})
	}
}
