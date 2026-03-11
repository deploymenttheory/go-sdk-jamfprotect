package dataretention_test

import (
	"context"
	"testing"

	dataretention "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/data_retention"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/data_retention/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*dataretention.Service, *mocks.DataRetentionMock) {
	t.Helper()
	mock := mocks.NewDataRetentionMock()
	return dataretention.NewService(mock), mock
}

func TestDataRetentionService_GetDataRetention(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterGetDataRetentionMock()

	result, _, err := service.GetDataRetention(context.Background())

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, int64(90), result.Database.Log.NumberOfDays)
	assert.Equal(t, int64(365), result.Database.Alert.NumberOfDays)
	assert.Equal(t, int64(730), result.Cold.Alert.NumberOfDays)
}

func TestDataRetentionService_UpdateDataRetention(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterUpdateDataRetentionMock()

	req := &dataretention.UpdateDataRetentionRequest{
		DatabaseLogDays:   30,
		DatabaseAlertDays: 180,
		ColdAlertDays:     365,
	}

	result, _, err := service.UpdateDataRetention(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, int64(30), result.Database.Log.NumberOfDays)
	assert.Equal(t, int64(180), result.Database.Alert.NumberOfDays)
}

func TestDataRetentionService_GetDataRetention_EmptyResult(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/graphql", "getDataRetention", 200, "get_data_retention_empty.json")

	result, _, err := service.GetDataRetention(context.Background())

	require.NoError(t, err)
	assert.Nil(t, result)
}

func TestDataRetentionService_UpdateDataRetention_EmptyResult(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/app", "updateOrganizationRetention", 200, "update_data_retention_empty.json")

	req := &dataretention.UpdateDataRetentionRequest{
		DatabaseLogDays:   30,
		DatabaseAlertDays: 180,
		ColdAlertDays:     365,
	}

	result, _, err := service.UpdateDataRetention(context.Background(), req)

	require.NoError(t, err)
	assert.Nil(t, result)
}

func TestDataRetentionService_GetDataRetention_Error(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterError("/graphql", "getDataRetention", 500, "", "Internal Server Error")

	result, _, err := service.GetDataRetention(context.Background())

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to get data retention")
}

func TestDataRetentionService_UpdateDataRetention_Error(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterError("/app", "updateOrganizationRetention", 500, "", "Internal Server Error")

	req := &dataretention.UpdateDataRetentionRequest{
		DatabaseLogDays:   30,
		DatabaseAlertDays: 180,
		ColdAlertDays:     365,
	}

	result, _, err := service.UpdateDataRetention(context.Background(), req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to update data retention")
}

func TestDataRetentionService_ValidationErrors(t *testing.T) {
	service, _ := setupMockService(t)

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{
			name: "UpdateDataRetention nil request",
			fn: func() error {
				_, _, err := service.UpdateDataRetention(context.Background(), nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "UpdateDataRetention zero databaseLogDays",
			fn: func() error {
				_, _, err := service.UpdateDataRetention(context.Background(), &dataretention.UpdateDataRetentionRequest{
					DatabaseLogDays:   0,
					DatabaseAlertDays: 180,
					ColdAlertDays:     365,
				})
				return err
			},
			wantErr: "databaseLogDays must be a positive integer",
		},
		{
			name: "UpdateDataRetention zero databaseAlertDays",
			fn: func() error {
				_, _, err := service.UpdateDataRetention(context.Background(), &dataretention.UpdateDataRetentionRequest{
					DatabaseLogDays:   30,
					DatabaseAlertDays: 0,
					ColdAlertDays:     365,
				})
				return err
			},
			wantErr: "databaseAlertDays must be a positive integer",
		},
		{
			name: "UpdateDataRetention zero coldAlertDays",
			fn: func() error {
				_, _, err := service.UpdateDataRetention(context.Background(), &dataretention.UpdateDataRetentionRequest{
					DatabaseLogDays:   30,
					DatabaseAlertDays: 180,
					ColdAlertDays:     0,
				})
				return err
			},
			wantErr: "coldAlertDays must be a positive integer",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fn()
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}
