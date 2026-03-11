package apiclient_test

import (
	"context"
	"testing"

	apiclient "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/api_client"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/api_client/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testClientID = "test-client-id-1234"

func setupMockService(t *testing.T) (*apiclient.Service, *mocks.ApiClientMock) {
	t.Helper()
	mock := mocks.NewApiClientMock()
	return apiclient.NewService(mock), mock
}

func TestApiClientService_CreateApiClient(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterCreateApiClientMock()

	req := &apiclient.CreateApiClientRequest{
		Name:    "Test API Client",
		RoleIDs: []string{"role-id-1"},
	}

	result, _, err := service.CreateApiClient(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testClientID, result.ClientID)
	assert.Equal(t, "Test API Client", result.Name)
	assert.NotEmpty(t, result.Password)
}

func TestApiClientService_GetApiClient(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterGetApiClientMock()

	result, _, err := service.GetApiClient(context.Background(), testClientID)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testClientID, result.ClientID)
	assert.Equal(t, "Test API Client", result.Name)
}

func TestApiClientService_UpdateApiClient(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterUpdateApiClientMock()

	req := &apiclient.UpdateApiClientRequest{
		Name:    "Updated API Client",
		RoleIDs: []string{"role-id-1"},
	}

	result, _, err := service.UpdateApiClient(context.Background(), testClientID, req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testClientID, result.ClientID)
	assert.Equal(t, "Updated API Client", result.Name)
}

func TestApiClientService_DeleteApiClient(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterDeleteApiClientMock()

	_, err := service.DeleteApiClient(context.Background(), testClientID)

	require.NoError(t, err)
}

func TestApiClientService_ListApiClients(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListApiClientsMock()

	result, _, err := service.ListApiClients(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, testClientID, result[0].ClientID)
	assert.Equal(t, "Test API Client", result[0].Name)
}

func TestApiClientService_ValidationErrors(t *testing.T) {
	service, _ := setupMockService(t)

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{
			name: "CreateApiClient nil request",
			fn: func() error {
				_, _, err := service.CreateApiClient(context.Background(), nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "CreateApiClient missing name",
			fn: func() error {
				_, _, err := service.CreateApiClient(context.Background(), &apiclient.CreateApiClientRequest{})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "GetApiClient empty clientId",
			fn: func() error {
				_, _, err := service.GetApiClient(context.Background(), "")
				return err
			},
			wantErr: "clientId is required",
		},
		{
			name: "UpdateApiClient empty clientId",
			fn: func() error {
				_, _, err := service.UpdateApiClient(context.Background(), "", &apiclient.UpdateApiClientRequest{Name: "test"})
				return err
			},
			wantErr: "clientId is required",
		},
		{
			name: "UpdateApiClient nil request",
			fn: func() error {
				_, _, err := service.UpdateApiClient(context.Background(), testClientID, nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "DeleteApiClient empty clientId",
			fn: func() error {
				_, err := service.DeleteApiClient(context.Background(), "")
				return err
			},
			wantErr: "clientId is required",
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
