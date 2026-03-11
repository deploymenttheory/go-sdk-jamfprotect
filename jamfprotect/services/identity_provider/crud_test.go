package identityprovider_test

import (
	"context"
	"testing"

	identityprovider "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/identity_provider"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/identity_provider/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testConnectionID = "conn-id-1234"

func setupMockService(t *testing.T) (*identityprovider.Service, *mocks.IdentityProviderMock) {
	t.Helper()
	mock := mocks.NewIdentityProviderMock()
	return identityprovider.NewService(mock), mock
}

func TestIdentityProviderService_ListConnections(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListConnectionsMock()

	result, _, err := service.ListConnections(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, testConnectionID, result[0].ID)
	assert.Equal(t, "Test Connection", result[0].Name)
	assert.Equal(t, "samlp", result[0].Strategy)
	assert.Equal(t, "saml", result[0].Source)
	assert.True(t, result[0].GroupsSupport)
}

func TestIdentityProviderService_ListConnections_EmptyResult(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/graphql", "listConnections", 200, "list_connections_empty.json")

	result, _, err := service.ListConnections(context.Background())

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

func TestIdentityProviderService_ListConnections_Error(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterError("/graphql", "listConnections", 500, "", "connection error")

	_, _, err := service.ListConnections(context.Background())

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to list connections")
}
