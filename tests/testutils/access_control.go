package testutils

import (
	"testing"

	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address/signaturescheme"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/root"
	"github.com/stretchr/testify/require"
)

// GrantDeployPermission gives permission to the specified agentID to deploy SCs into the chain
func GrantDeployPermission(chain *solo.Chain, sigScheme signaturescheme.SignatureScheme, deployerAgentID coretypes.AgentID, params ...interface{}) error {
	if sigScheme == nil {
		sigScheme = chain.OriginatorSigScheme
	}

	par := []interface{}{root.ParamDeployer, deployerAgentID}
	par = append(par, params...)
	req := solo.NewCallParams(root.Interface.Name, root.FuncGrantDeploy, par...)
	_, err := chain.PostRequest(req, sigScheme)
	return err
}

// RevokeDeployPermission removes permission of the specified agentID to deploy SCs into the chain
func RevokeDeployPermission(chain *solo.Chain, sigScheme signaturescheme.SignatureScheme, deployerAgentID coretypes.AgentID, params ...interface{}) error {
	if sigScheme == nil {
		sigScheme = chain.OriginatorSigScheme
	}

	par := []interface{}{root.ParamDeployer, deployerAgentID}
	par = append(par, params...)
	req := solo.NewCallParams(root.Interface.Name, root.FuncRevokeDeploy, par...)
	_, err := chain.PostRequest(req, sigScheme)
	return err
}

// RequireAccess fails a unit test if unauthorized access is given to caller
func RequireAccess(t *testing.T, ownerKeyPair signaturescheme.SignatureScheme, callerKeyPair signaturescheme.SignatureScheme, err error) {
	unauthozizedAcess := ownerKeyPair != nil && ownerKeyPair != callerKeyPair
	if unauthozizedAcess {
		require.Error(t, err, "Access given to unauthorized key pair")
	} else {
		require.NoError(t, err)
	}
}
