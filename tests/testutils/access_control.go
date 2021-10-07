package testutils

import (
	"testing"

	"github.com/iotaledger/hive.go/crypto/ed25519"
	"github.com/stretchr/testify/require"
)

// RequireAccess fails a unit test if unauthorized access is given to caller
func RequireAccess(t *testing.T, ownerKeyPair *ed25519.KeyPair, callerKeyPair *ed25519.KeyPair, err error) {
	unauthozizedAcess := ownerKeyPair != nil && ownerKeyPair != callerKeyPair
	if unauthozizedAcess {
		require.Error(t, err, "Access given to unauthorized key pair")
	} else {
		require.NoError(t, err)
	}
}
