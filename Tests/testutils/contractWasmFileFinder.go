package testutils

import (
	"testing"

	"github.com/drand/drand/fs"
	"github.com/stretchr/testify/require"
)

// MustGetContractWasmFilePath ensures a given smart contract's wasm file exists
func MustGetContractWasmFilePath(t *testing.T, contractName string) string {
	contractWasmFilePath := "../SmartContract/pkg/" + contractName + "_bg.wasm"
	exists, err := fs.Exists(contractWasmFilePath)
	require.NoError(t, err, "Error trying to find file: "+contractWasmFilePath)

	if !exists {
		contractWasmFilePath = "../../SmartContract/pkg/" + contractName + "_bg.wasm"
		exists, err = fs.Exists(contractWasmFilePath)
		require.NoError(t, err, "Error trying to find file: "+contractWasmFilePath)
	}

	if !exists {
		contractWasmFilePath = "../../../SmartContract/pkg/" + contractName + "_bg.wasm"
		exists, err = fs.Exists(contractWasmFilePath)
		require.NoError(t, err, "Error trying to find file: "+contractWasmFilePath)
	}

	require.True(t, exists, "File does not exist: "+contractWasmFilePath)
	return contractWasmFilePath
}
