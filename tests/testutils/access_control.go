package testutils

import (
	"testing"

	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address/signaturescheme"
	"github.com/stretchr/testify/require"
)

// RequireAccess fails a unit test if unauthorized access is given to caller
func RequireAccess(t *testing.T, ownerSigScheme signaturescheme.SignatureScheme, callerSigScheme signaturescheme.SignatureScheme, err error) {
	unauthozizedAcess := ownerSigScheme != nil && ownerSigScheme != callerSigScheme
	if unauthozizedAcess {
		require.Error(t, err, "Access given to unauthorized key pair")
	} else {
		require.NoError(t, err)
	}
}
