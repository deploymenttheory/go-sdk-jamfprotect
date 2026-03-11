package exceptionset_test

import (
	"context"
	"testing"

	exceptionset "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/exception_set"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/exception_set/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testUUID = "aaaaaaaa-bbbb-4ccc-8ddd-eeeeeeeeeeee"

func setupMockService(t *testing.T) (*exceptionset.Service, *mocks.ExceptionSetMock) {
	t.Helper()
	mock := mocks.NewExceptionSetMock()
	return exceptionset.NewService(mock), mock
}

func TestExceptionSetService_CreateExceptionSet(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterCreateExceptionSetMock()

	req := &exceptionset.CreateExceptionSetRequest{
		Name:        "Test Exception Set",
		Description: "A test exception set",
	}

	result, _, err := service.CreateExceptionSet(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testUUID, result.UUID)
	assert.Equal(t, "Test Exception Set", result.Name)
}

func TestExceptionSetService_GetExceptionSet(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterGetExceptionSetMock()

	result, _, err := service.GetExceptionSet(context.Background(), testUUID)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testUUID, result.UUID)
	assert.Equal(t, "Test Exception Set", result.Name)
}

func TestExceptionSetService_UpdateExceptionSet(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterUpdateExceptionSetMock()

	req := &exceptionset.UpdateExceptionSetRequest{
		Name:        "Updated Exception Set",
		Description: "An updated exception set",
	}

	result, _, err := service.UpdateExceptionSet(context.Background(), testUUID, req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testUUID, result.UUID)
	assert.Equal(t, "Updated Exception Set", result.Name)
}

func TestExceptionSetService_DeleteExceptionSet(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterDeleteExceptionSetMock()

	_, err := service.DeleteExceptionSet(context.Background(), testUUID)

	require.NoError(t, err)
}

func TestExceptionSetService_ListExceptionSets(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListExceptionSetsMock()

	result, _, err := service.ListExceptionSets(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, testUUID, result[0].UUID)
}

func TestExceptionSetService_ListExceptionSetNames(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListExceptionSetNamesMock()

	result, _, err := service.ListExceptionSetNames(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Test Exception Set", result[0])
}

func TestExceptionSetService_ValidationErrors(t *testing.T) {
	service, _ := setupMockService(t)

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{
			name: "CreateExceptionSet nil request",
			fn: func() error {
				_, _, err := service.CreateExceptionSet(context.Background(), nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "CreateExceptionSet missing name",
			fn: func() error {
				_, _, err := service.CreateExceptionSet(context.Background(), &exceptionset.CreateExceptionSetRequest{})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "GetExceptionSet empty uuid",
			fn: func() error {
				_, _, err := service.GetExceptionSet(context.Background(), "")
				return err
			},
			wantErr: "uuid is required",
		},
		{
			name: "DeleteExceptionSet empty uuid",
			fn: func() error {
				_, err := service.DeleteExceptionSet(context.Background(), "")
				return err
			},
			wantErr: "uuid is required",
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
