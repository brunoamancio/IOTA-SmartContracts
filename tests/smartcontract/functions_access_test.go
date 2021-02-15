package libtest

import (
	"testing"

	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils"
	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils/testconstants"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address/signaturescheme"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/stretchr/testify/require"
)

// Contract name - Defined in SmartContract/Cargo.toml > package > name
const contractName = "my_iota_sc"

// Test ensures only the expected callers have access to functions
func Test_access_to_functions(t *testing.T) {
	env := solo.New(t, testconstants.Debug, testconstants.StackTrace)

	// Create chain with chainOwnerKeyPair
	chainOwnerKeyPair := env.NewSignatureSchemeWithFunds()
	chain := env.NewChain(chainOwnerKeyPair, "test_access_chain")

	// Create contractCreator key pair and give it permission to deploy into chain
	contractCreatorKeyPair := env.NewSignatureSchemeWithFunds()
	contractCreatorAgentID := coretypes.NewAgentIDFromAddress(contractCreatorKeyPair.Address())
	testutils.GrantDeployPermission(chain, chainOwnerKeyPair, contractCreatorAgentID)

	// Deploy contract with contractOwnerKeyPair
	contractFilePath := testutils.MustGetContractWasmFilePath(t, contractName)
	err := chain.DeployWasmContract(contractCreatorKeyPair, contractName, contractFilePath)
	require.NoError(t, err)

	// Map of SC functions and function owners
	functionsToTest := make(map[string]signaturescheme.SignatureScheme)

	// Name of the SC function to be requested and credential required to acess it
	functionsToTest["my_sc_function"] = nil                                    // public function
	functionsToTest["contract_creator_only_function"] = contractCreatorKeyPair // owner-only function
	functionsToTest["chain_owner_only_function"] = chainOwnerKeyPair           // owner-only function

	// Create random key pair
	randomKeyPair := env.NewSignatureSchemeWithFunds()

	for functionName, ownerKeyPair := range functionsToTest {
		t.Run(functionName, func(t *testing.T) {

			// Defines which contract and function will be called by chain.PostRequest
			reqParams := solo.NewCallParams(contractName, functionName)

			// Calls SC function as chainOwner
			_, err = chain.PostRequest(reqParams, chainOwnerKeyPair)
			// Verifies if access to SC function was given to caller. Fail if unauthorized acess.
			testutils.RequireAccess(t, ownerKeyPair, chainOwnerKeyPair, err)

			// Calls SC function as contractCreator
			_, err = chain.PostRequest(reqParams, contractCreatorKeyPair)
			// Verifies if access to SC function was given to caller. Fail if unauthorized acess.
			testutils.RequireAccess(t, ownerKeyPair, contractCreatorKeyPair, err)

			// Calls SC function as anyone else (random)
			_, err = chain.PostRequest(reqParams, randomKeyPair)
			// Verifies if access to SC function was given to caller. Fail if unauthorized acess.
			testutils.RequireAccess(t, ownerKeyPair, randomKeyPair, err)
		})
	}
}
