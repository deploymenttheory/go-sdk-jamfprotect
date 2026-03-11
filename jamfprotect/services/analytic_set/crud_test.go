package analyticset_test

import (
	"context"
	"testing"

	analyticset "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/analytic_set"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/analytic_set/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testUUID = "aaaaaaaa-bbbb-4ccc-8ddd-eeeeeeeeeeee"

func setupMockService(t *testing.T) (*analyticset.Service, *mocks.AnalyticSetMock) {
	t.Helper()
	mock := mocks.NewAnalyticSetMock()
	return analyticset.NewService(mock), mock
}

func TestAnalyticSetService_CreateAnalyticSet(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterCreateAnalyticSetMock()

	req := &analyticset.CreateAnalyticSetRequest{
		Name:        "Test Analytic Set",
		Description: "A test analytic set",
		Analytics:   []string{"analytic-uuid-1"},
	}

	result, _, err := service.CreateAnalyticSet(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testUUID, result.UUID)
	assert.Equal(t, "Test Analytic Set", result.Name)
}

func TestAnalyticSetService_GetAnalyticSet(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterGetAnalyticSetMock()

	result, _, err := service.GetAnalyticSet(context.Background(), testUUID)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testUUID, result.UUID)
	assert.Equal(t, "Test Analytic Set", result.Name)
}

func TestAnalyticSetService_UpdateAnalyticSet(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterUpdateAnalyticSetMock()

	req := &analyticset.UpdateAnalyticSetRequest{
		Name:        "Updated Analytic Set",
		Description: "An updated analytic set",
		Analytics:   []string{"analytic-uuid-1"},
	}

	result, _, err := service.UpdateAnalyticSet(context.Background(), testUUID, req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testUUID, result.UUID)
	assert.Equal(t, "Updated Analytic Set", result.Name)
}

func TestAnalyticSetService_DeleteAnalyticSet(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterDeleteAnalyticSetMock()

	_, err := service.DeleteAnalyticSet(context.Background(), testUUID)

	require.NoError(t, err)
}

func TestAnalyticSetService_ListAnalyticSets(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListAnalyticSetsMock()

	result, _, err := service.ListAnalyticSets(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, testUUID, result[0].UUID)
}

func TestAnalyticSetService_ValidationErrors(t *testing.T) {
	service, _ := setupMockService(t)

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{
			name: "CreateAnalyticSet nil request",
			fn: func() error {
				_, _, err := service.CreateAnalyticSet(context.Background(), nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "CreateAnalyticSet missing name",
			fn: func() error {
				_, _, err := service.CreateAnalyticSet(context.Background(), &analyticset.CreateAnalyticSetRequest{
					Analytics: []string{"uuid-1"},
				})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "GetAnalyticSet empty uuid",
			fn: func() error {
				_, _, err := service.GetAnalyticSet(context.Background(), "")
				return err
			},
			wantErr: "uuid is required",
		},
		{
			name: "DeleteAnalyticSet empty uuid",
			fn: func() error {
				_, err := service.DeleteAnalyticSet(context.Background(), "")
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
