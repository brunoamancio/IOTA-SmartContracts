package libtest

import (
	"testing"

	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils"
	"github.com/brunoamancio/IOTA-SmartContracts/tests/testutils/testconstants"
	notsolo "github.com/brunoamancio/NotSolo"
)

//  -----------------------------------------------  //
//  See code samples in Tests/testutils/codesamples  //
//  -----------------------------------------------  //

func TestLib(t *testing.T) {
	contractWasmFilePath := testutils.MustGetContractWasmFilePath(t, testconstants.ContractName) // You can use if file is in SmartContract/pkg

	// Name of the SC function to be requested - Defined in lib.rs > add_call > my_sc_function
	functionName := "my_sc_function"

	notSolo := notsolo.New(t)

	chainName := testconstants.ContractName + "Chain"
	chain := notSolo.Chain.NewChain(nil, chainName)

	// Uploads wasm of SC and deploys it into chain
	notSolo.Chain.DeployWasmContract(chain, nil, testconstants.ContractName, contractWasmFilePath)

	// Call contract 'my_iota_sc', function 'my_sc_function'
	notSolo.Request.MustPost(nil, chain, testconstants.ContractName, functionName)
}
