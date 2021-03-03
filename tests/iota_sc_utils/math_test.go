package iotascutils

import (
	"testing"

	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils"
	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils/testconstants"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/stretchr/testify/require"
)

func Test_math_functions(t *testing.T) {
	env := solo.New(t, testconstants.Debug, testconstants.StackTrace)
	chain := env.NewChain(nil, "mathChain")

	// Deploy contract with chainOwnerKeyPair
	contractFilePath := testutils.MustGetContractWasmFilePath(t, testconstants.ContractName)
	err := chain.DeployWasmContract(nil, testconstants.ContractName, contractFilePath)
	require.NoError(t, err)

	typesToTest := []string{"u8", "u16", "u32", "u64", "usize", "i8", "i16", "i32", "i64", "isize"}
	operationsToTest := []string{"add", "sub", "mul", "div"}

	// Map of SC functions and expect success (true if yes)
	functionsToTest := make(map[string]bool)
	for _, operationToTest := range operationsToTest {
		for _, typeToTest := range typesToTest {
			// Name of the SC function to be requested and credential required to access it
			functionsToTest[typeToTest+"_safe_"+operationToTest+"_no_overflow_function"] = true    // expect success
			functionsToTest[typeToTest+"_safe_"+operationToTest+"_with_overflow_function"] = false // expect failure
		}
	}

	for functionName, expectSuccess := range functionsToTest {
		t.Run(functionName, func(t *testing.T) {
			// Defines which contract and function will be called by chain.PostRequest.
			reqParams := solo.NewCallParams(testconstants.ContractName, functionName)

			// Calls SC function as chainOwner
			_, err := chain.PostRequestSync(reqParams, nil)

			// Verifies if SC function call is executed or fails
			if expectSuccess {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
