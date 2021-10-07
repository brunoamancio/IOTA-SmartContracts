package iotascutils

import (
	"testing"

	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils"
	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils/testconstants"
	notsolo "github.com/brunoamancio/NotSolo"
	"github.com/iotaledger/hive.go/crypto/ed25519"
)

// Test ensures only the expected callers have access to functions
func Test_access_to_functions(t *testing.T) {
	notSolo := notsolo.New(t)

	// Create chain with chainOwnerKeyPair
	chainOwnerKeyPair := notSolo.KeyPair.NewKeyPairWithFunds()
	chain := notSolo.Chain.NewChain(chainOwnerKeyPair, "test_access_chain")

	// Create contractCreator key pair and give it permission to deploy into chain
	contractOriginatorKeyPair := notSolo.KeyPair.NewKeyPairWithFunds()
	notSolo.Chain.GrantDeployPermission(chain, contractOriginatorKeyPair)

	// Deploy contract with contractOwnerKeyPair
	contractFilePath := testutils.MustGetContractWasmFilePath(t, testconstants.ContractName)
	notSolo.Chain.DeployWasmContract(chain, contractOriginatorKeyPair, testconstants.ContractName, contractFilePath)

	// Create random key pair
	randomKeyPair := notSolo.KeyPair.NewKeyPairWithFunds()

	// Map of SC functions and function owners
	functionsToTest := make(map[string]*ed25519.KeyPair)

	// Name of the SC function to be requested and credential required to access it
	functionsToTest["my_sc_function"] = nil                                       // public function
	functionsToTest["contract_creator_only_function"] = contractOriginatorKeyPair // owner-only function
	functionsToTest["chain_owner_only_function"] = chainOwnerKeyPair              // owner-only function

	for functionName, ownerKeyPair := range functionsToTest {
		t.Run(functionName, func(t *testing.T) {
			// Calls SC function as chainOwner
			_, err := notSolo.Request.Post(chainOwnerKeyPair, chain, testconstants.ContractName, functionName)

			// Verifies if access to SC function was given to caller. Fail if unauthorized access.
			testutils.RequireAccess(t, ownerKeyPair, chainOwnerKeyPair, err)

			// Calls SC function as contractCreator
			_, err = notSolo.Request.Post(contractOriginatorKeyPair, chain, testconstants.ContractName, functionName)
			// Verifies if access to SC function was given to caller. Fail if unauthorized access.
			testutils.RequireAccess(t, ownerKeyPair, contractOriginatorKeyPair, err)

			// Calls SC function as anyone else (random)
			_, err = notSolo.Request.Post(randomKeyPair, chain, testconstants.ContractName, functionName)
			// Verifies if access to SC function was given to caller. Fail if unauthorized access.
			testutils.RequireAccess(t, ownerKeyPair, randomKeyPair, err)
		})
	}
}
