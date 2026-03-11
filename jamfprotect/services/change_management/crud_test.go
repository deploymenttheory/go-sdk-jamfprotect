package changemanagement_test

import (
	"context"
	"testing"

	changemanagement "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/change_management"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/change_management/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*changemanagement.Service, *mocks.ChangeManagementMock) {
	t.Helper()
	mock := mocks.NewChangeManagementMock()
	return changemanagement.NewService(mock), mock
}

func TestChangeManagementService_GetConfigFreeze(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterGetConfigFreezeMock()

	result, _, err := service.GetConfigFreeze(context.Background())

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, false, result.ConfigFreeze)
}

func TestChangeManagementService_UpdateConfigFreeze(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterUpdateConfigFreezeMock()

	result, _, err := service.UpdateConfigFreeze(context.Background(), true)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, true, result.ConfigFreeze)
}

func TestChangeManagementService_GetConfigFreeze_NilResult(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/graphql", "getConfigFreeze", 200, "get_config_freeze_empty.json")

	result, _, err := service.GetConfigFreeze(context.Background())

	require.NoError(t, err)
	assert.Nil(t, result)
}

func TestChangeManagementService_UpdateConfigFreeze_NilResult(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/app", "updateOrganizationConfigFreeze", 200, "update_config_freeze_empty.json")

	result, _, err := service.UpdateConfigFreeze(context.Background(), false)

	require.NoError(t, err)
	assert.Nil(t, result)
}

func TestChangeManagementService_GetConfigFreeze_Error(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterError("/graphql", "getConfigFreeze", 500, "", "Internal Server Error")

	result, _, err := service.GetConfigFreeze(context.Background())

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to get config freeze")
}

func TestChangeManagementService_UpdateConfigFreeze_Error(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterError("/app", "updateOrganizationConfigFreeze", 500, "", "Internal Server Error")

	result, _, err := service.UpdateConfigFreeze(context.Background(), true)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to update config freeze")
}
