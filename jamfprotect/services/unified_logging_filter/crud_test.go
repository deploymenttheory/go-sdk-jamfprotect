package unifiedloggingfilter_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/unified_logging_filter"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/unified_logging_filter/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testUUID = "aaaaaaaa-bbbb-4ccc-8ddd-eeeeeeeeeeee"

func setupMockService(t *testing.T) (*unifiedloggingfilter.Service, *mocks.UnifiedLoggingFilterMock) {
	t.Helper()
	mock := mocks.NewUnifiedLoggingFilterMock()
	return unifiedloggingfilter.NewService(mock), mock
}

func TestUnifiedLoggingFilterService_CreateUnifiedLoggingFilter(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterCreateUnifiedLoggingFilterMock()

	req := &unifiedloggingfilter.CreateUnifiedLoggingFilterRequest{
		Name:        "Test Unified Logging Filter",
		Description: "A test unified logging filter",
		Filter:      "process.name == \"test\"",
		Enabled:     true,
	}

	result, _, err := service.CreateUnifiedLoggingFilter(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testUUID, result.UUID)
	assert.Equal(t, "Test Unified Logging Filter", result.Name)
}

func TestUnifiedLoggingFilterService_GetUnifiedLoggingFilter(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterGetUnifiedLoggingFilterMock()

	result, _, err := service.GetUnifiedLoggingFilter(context.Background(), testUUID)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testUUID, result.UUID)
	assert.Equal(t, "Test Unified Logging Filter", result.Name)
}

func TestUnifiedLoggingFilterService_UpdateUnifiedLoggingFilter(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterUpdateUnifiedLoggingFilterMock()

	req := &unifiedloggingfilter.UpdateUnifiedLoggingFilterRequest{
		Name:        "Updated Unified Logging Filter",
		Description: "An updated unified logging filter",
		Filter:      "process.name == \"updated\"",
		Enabled:     true,
	}

	result, _, err := service.UpdateUnifiedLoggingFilter(context.Background(), testUUID, req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testUUID, result.UUID)
	assert.Equal(t, "Updated Unified Logging Filter", result.Name)
}

func TestUnifiedLoggingFilterService_DeleteUnifiedLoggingFilter(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterDeleteUnifiedLoggingFilterMock()

	_, err := service.DeleteUnifiedLoggingFilter(context.Background(), testUUID)

	require.NoError(t, err)
}

func TestUnifiedLoggingFilterService_ListUnifiedLoggingFilters(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListUnifiedLoggingFiltersMock()

	result, _, err := service.ListUnifiedLoggingFilters(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, testUUID, result[0].UUID)
	assert.Equal(t, "Test Unified Logging Filter", result[0].Name)
}

func TestUnifiedLoggingFilterService_ListUnifiedLoggingFilterNames(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListUnifiedLoggingFilterNamesMock()

	result, _, err := service.ListUnifiedLoggingFilterNames(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Test Unified Logging Filter", result[0])
}

func TestUnifiedLoggingFilterService_ListUnifiedLoggingFilters_EmptyResult(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/graphql", "listUnifiedLoggingFilters", 200, "list_unified_logging_filters_empty.json")

	result, _, err := service.ListUnifiedLoggingFilters(context.Background())

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

func TestUnifiedLoggingFilterService_ListUnifiedLoggingFilterNames_EmptyResult(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/graphql", "listUnifiedLoggingFilterNames", 200, "list_unified_logging_filter_names_empty.json")

	result, _, err := service.ListUnifiedLoggingFilterNames(context.Background())

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

func TestUnifiedLoggingFilterService_ValidationErrors(t *testing.T) {
	service, _ := setupMockService(t)

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{
			name: "CreateUnifiedLoggingFilter nil request",
			fn: func() error {
				_, _, err := service.CreateUnifiedLoggingFilter(context.Background(), nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "CreateUnifiedLoggingFilter missing name",
			fn: func() error {
				_, _, err := service.CreateUnifiedLoggingFilter(context.Background(), &unifiedloggingfilter.CreateUnifiedLoggingFilterRequest{
					Filter: "test",
				})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "CreateUnifiedLoggingFilter missing filter",
			fn: func() error {
				_, _, err := service.CreateUnifiedLoggingFilter(context.Background(), &unifiedloggingfilter.CreateUnifiedLoggingFilterRequest{
					Name: "test",
				})
				return err
			},
			wantErr: "filter is required",
		},
		{
			name: "GetUnifiedLoggingFilter empty uuid",
			fn: func() error {
				_, _, err := service.GetUnifiedLoggingFilter(context.Background(), "")
				return err
			},
			wantErr: "uuid is required",
		},
		{
			name: "UpdateUnifiedLoggingFilter empty uuid",
			fn: func() error {
				_, _, err := service.UpdateUnifiedLoggingFilter(context.Background(), "", &unifiedloggingfilter.UpdateUnifiedLoggingFilterRequest{
					Name:   "test",
					Filter: "test",
				})
				return err
			},
			wantErr: "uuid is required",
		},
		{
			name: "UpdateUnifiedLoggingFilter nil request",
			fn: func() error {
				_, _, err := service.UpdateUnifiedLoggingFilter(context.Background(), testUUID, nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "UpdateUnifiedLoggingFilter missing name",
			fn: func() error {
				_, _, err := service.UpdateUnifiedLoggingFilter(context.Background(), testUUID, &unifiedloggingfilter.UpdateUnifiedLoggingFilterRequest{
					Filter: "test",
				})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "UpdateUnifiedLoggingFilter missing filter",
			fn: func() error {
				_, _, err := service.UpdateUnifiedLoggingFilter(context.Background(), testUUID, &unifiedloggingfilter.UpdateUnifiedLoggingFilterRequest{
					Name: "test",
				})
				return err
			},
			wantErr: "filter is required",
		},
		{
			name: "DeleteUnifiedLoggingFilter empty uuid",
			fn: func() error {
				_, err := service.DeleteUnifiedLoggingFilter(context.Background(), "")
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

func TestUnifiedLoggingFilterService_Validators(t *testing.T) {
	assert.NoError(t, unifiedloggingfilter.ValidateCreateUnifiedLoggingFilterRequest(&unifiedloggingfilter.CreateUnifiedLoggingFilterRequest{}))
	assert.NoError(t, unifiedloggingfilter.ValidateCreateUnifiedLoggingFilterRequest(nil))

	assert.NoError(t, unifiedloggingfilter.ValidateUpdateUnifiedLoggingFilterRequest(&unifiedloggingfilter.UpdateUnifiedLoggingFilterRequest{}))
	assert.NoError(t, unifiedloggingfilter.ValidateUpdateUnifiedLoggingFilterRequest(nil))

	validUUID := "550e8400-e29b-41d4-a716-446655440000"
	assert.NoError(t, unifiedloggingfilter.ValidateUnifiedLoggingFilterUUID(validUUID))

	assert.Error(t, unifiedloggingfilter.ValidateUnifiedLoggingFilterUUID(""))
	assert.Error(t, unifiedloggingfilter.ValidateUnifiedLoggingFilterUUID("not-a-uuid"))
	assert.Error(t, unifiedloggingfilter.ValidateUnifiedLoggingFilterUUID("550e8400"))
}
