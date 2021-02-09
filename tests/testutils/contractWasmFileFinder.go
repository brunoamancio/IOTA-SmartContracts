package testutils

import (
	"testing"

	"github.com/drand/drand/fs"
	"github.com/stretchr/testify/require"
)

// MustGetContractWasmFilePath ensures a given smart contract's wasm file exists
func MustGetContractWasmFilePath(t *testing.T, contractName string) string {
	const targetPath = "smartContracs/target/"
	const parentDirectoryLevel = "../"

	filePath := targetPath + contractName + "_bg.wasm"

	contractWasmFilePath := parentDirectoryLevel + filePath
	exists, err := fs.Exists(contractWasmFilePath)
	require.NoError(t, err, "Error trying to find file: "+contractWasmFilePath)

	if !exists {
		contractWasmFilePath := parentDirectoryLevel + parentDirectoryLevel + filePath
		exists, err = fs.Exists(contractWasmFilePath)
		require.NoError(t, err, "Error trying to find file: "+contractWasmFilePath)
	}

	if !exists {
		contractWasmFilePath := parentDirectoryLevel + parentDirectoryLevel + parentDirectoryLevel + filePath
		exists, err = fs.Exists(contractWasmFilePath)
		require.NoError(t, err, "Error trying to find file: "+contractWasmFilePath)
	}

	require.True(t, exists, "File does not exist: "+contractWasmFilePath)
	return contractWasmFilePath
}
