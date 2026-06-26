package beta_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/beta"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/beta/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*beta.Service, *mocks.BetaMock) {
	t.Helper()
	mock := mocks.NewBetaMock()
	return beta.NewService(mock), mock
}

func TestBetaService_GetBetaAcceptanceStatus(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterGetBetaAcceptanceStatusMock()

	result, _, err := service.GetBetaAcceptanceStatus(context.Background())

	require.NoError(t, err)
	require.Len(t, result, 1)
	assert.Equal(t, "NGTP_BETA", result[0].BetaName)
	assert.Equal(t, "2024-01-01T00:00:00Z", result[0].AcceptedTimestamp)
	assert.Equal(t, "admin@example.com", result[0].AcceptedUser)
}

func TestBetaService_GetBetaAcceptanceStatus_Empty(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/app", "getBetaAcceptanceStatus", 200, "get_beta_acceptance_status_empty.json")

	result, _, err := service.GetBetaAcceptanceStatus(context.Background())

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

func TestBetaService_UpdateBetaAcceptanceStatus(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterUpdateBetaAcceptanceStatusMock()

	result, _, err := service.UpdateBetaAcceptanceStatus(context.Background(), beta.BetaNameNGTP)

	require.NoError(t, err)
	require.Len(t, result, 1)
	assert.Equal(t, "NGTP_BETA", result[0].BetaName)
}

func TestBetaService_UpdateBetaAcceptanceStatus_InvalidInput(t *testing.T) {
	service, _ := setupMockService(t)

	result, _, err := service.UpdateBetaAcceptanceStatus(context.Background(), "")

	require.Error(t, err)
	assert.Nil(t, result)
}
