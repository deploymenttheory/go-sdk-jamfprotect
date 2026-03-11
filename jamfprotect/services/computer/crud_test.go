package computer_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/computer"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/computer/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testUUID = "aaaaaaaa-bbbb-4ccc-8ddd-eeeeeeeeeeee"

func setupMockService(t *testing.T) (*computer.Service, *mocks.ComputerMock) {
	t.Helper()
	mock := mocks.NewComputerMock()
	return computer.NewService(mock), mock
}

func TestComputerService_GetComputer(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterGetComputerMock()

	result, _, err := service.GetComputer(context.Background(), testUUID)

	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, result.UUID)
	assert.Equal(t, testUUID, *result.UUID)
	assert.NotNil(t, result.HostName)
}

func TestComputerService_ListComputers(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListComputersMock()

	result, _, err := service.ListComputers(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	require.NotNil(t, result[0].UUID)
	assert.Equal(t, testUUID, *result[0].UUID)
}

func TestComputerService_ListComputers_EmptyResult(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/graphql", "listComputers", 200, "list_computers_empty.json")

	result, _, err := service.ListComputers(context.Background())

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

func TestComputerService_GetComputer_NilResult(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/graphql", "getComputer", 200, "get_computer_empty.json")

	result, _, err := service.GetComputer(context.Background(), "550e8400-e29b-41d4-a716-446655440000")

	require.NoError(t, err)
	assert.Nil(t, result)
}

func TestComputerService_ValidationErrors(t *testing.T) {
	service, _ := setupMockService(t)

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{
			name: "GetComputer empty uuid",
			fn: func() error {
				_, _, err := service.GetComputer(context.Background(), "")
				return err
			},
			wantErr: "uuid is required",
		},
		{
			name: "GetComputer invalid uuid",
			fn: func() error {
				_, _, err := service.GetComputer(context.Background(), "not-a-uuid")
				return err
			},
			wantErr: "uuid must be a valid UUID",
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
