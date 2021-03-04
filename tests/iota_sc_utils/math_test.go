package iotascutils

import (
	"testing"

	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils"
	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils/testconstants"
	notsolo "github.com/brunoamancio/NotSolo"
	"github.com/stretchr/testify/require"
)

func Test_math_functions(t *testing.T) {
	notSolo := notsolo.New(t)
	chain := notSolo.Chain.NewChain(nil, "mathChain")

	// Deploy contract with chainOwnerSigScheme
	contractFilePath := testutils.MustGetContractWasmFilePath(t, testconstants.ContractName)
	notSolo.Chain.DeployWasmContract(chain, nil, testconstants.ContractName, contractFilePath)

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

			// Calls SC function as chainOwner
			_, err := notSolo.Request.Post(nil, chain, testconstants.ContractName, functionName)

			// Verifies if SC function call is executed or fails
			if expectSuccess {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
