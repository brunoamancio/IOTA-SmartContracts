package testutils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// MustGetContractWasmFilePath ensures a given smart contract's wasm file exists
func MustGetContractWasmFilePath(t *testing.T, contractName string) string {
	const targetPath = "../smartcontract/rust/pkg/"
	const parentDirectoryLevel = "../"

	filePath := targetPath + contractName + "_bg.wasm"

	contractWasmFilePath := parentDirectoryLevel + filePath
	exists, err := existsFilePath(contractWasmFilePath)
	require.NoError(t, err, "Error trying to find file: "+contractWasmFilePath)

	if !exists {
		contractWasmFilePath = parentDirectoryLevel + parentDirectoryLevel + filePath
		exists, err = existsFilePath(contractWasmFilePath)
		require.NoError(t, err, "Error trying to find file: "+contractWasmFilePath)
	}

	if !exists {
		contractWasmFilePath = parentDirectoryLevel + parentDirectoryLevel + parentDirectoryLevel + filePath
		exists, err = existsFilePath(contractWasmFilePath)
		require.NoError(t, err, "Error trying to find file: "+contractWasmFilePath)
	}

	require.True(t, exists, "File does not exist: "+contractWasmFilePath)
	return contractWasmFilePath
}

func existsFilePath(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
