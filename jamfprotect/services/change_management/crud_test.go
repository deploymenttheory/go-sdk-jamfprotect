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
