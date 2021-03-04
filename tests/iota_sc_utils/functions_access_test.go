package iotascutils

import (
	"testing"

	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils"
	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils/testconstants"
	notsolo "github.com/brunoamancio/NotSolo"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address/signaturescheme"
)

// Test ensures only the expected callers have access to functions
func Test_access_to_functions(t *testing.T) {
	notSolo := notsolo.New(t)

	// Create chain with chainOwnerSigScheme
	chainOwnerSigScheme := notSolo.SigScheme.NewSignatureSchemeWithFunds()
	chain := notSolo.Chain.NewChain(chainOwnerSigScheme, "test_access_chain")

	// Create contractCreator key pair and give it permission to deploy into chain
	contractOriginatorSigScheme := notSolo.SigScheme.NewSignatureSchemeWithFunds()
	notSolo.Chain.GrantDeployPermission(chain, contractOriginatorSigScheme)

	// Deploy contract with contractOwnerSigScheme
	contractFilePath := testutils.MustGetContractWasmFilePath(t, testconstants.ContractName)
	notSolo.Chain.DeployWasmContract(chain, contractOriginatorSigScheme, testconstants.ContractName, contractFilePath)

	// Create random key pair
	randomSigScheme := notSolo.SigScheme.NewSignatureSchemeWithFunds()

	// Map of SC functions and function owners
	functionsToTest := make(map[string]signaturescheme.SignatureScheme)

	// Name of the SC function to be requested and credential required to access it
	functionsToTest["my_sc_function"] = nil                                         // public function
	functionsToTest["contract_creator_only_function"] = contractOriginatorSigScheme // owner-only function
	functionsToTest["chain_owner_only_function"] = chainOwnerSigScheme              // owner-only function

	for functionName, ownerSigScheme := range functionsToTest {
		t.Run(functionName, func(t *testing.T) {
			// Calls SC function as chainOwner
			_, err := notSolo.Request.Post(chainOwnerSigScheme, chain, testconstants.ContractName, functionName)

			// Verifies if access to SC function was given to caller. Fail if unauthorized access.
			testutils.RequireAccess(t, ownerSigScheme, chainOwnerSigScheme, err)

			// Calls SC function as contractCreator
			_, err = notSolo.Request.Post(contractOriginatorSigScheme, chain, testconstants.ContractName, functionName)
			// Verifies if access to SC function was given to caller. Fail if unauthorized access.
			testutils.RequireAccess(t, ownerSigScheme, contractOriginatorSigScheme, err)

			// Calls SC function as anyone else (random)
			_, err = notSolo.Request.Post(randomSigScheme, chain, testconstants.ContractName, functionName)
			// Verifies if access to SC function was given to caller. Fail if unauthorized access.
			testutils.RequireAccess(t, ownerSigScheme, randomSigScheme, err)
		})
	}
}
