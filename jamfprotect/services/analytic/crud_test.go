package analytic_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/analytic"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/analytic/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testUUID = "aaaaaaaa-bbbb-4ccc-8ddd-eeeeeeeeeeee"

func setupMockService(t *testing.T) (*analytic.Service, *mocks.AnalyticMock) {
	t.Helper()
	mock := mocks.NewAnalyticMock()
	return analytic.NewService(mock), mock
}

func TestAnalyticService_CreateAnalytic(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterCreateAnalyticMock()

	req := &analytic.CreateAnalyticRequest{
		Name:        "Test Analytic",
		InputType:   "GPFSEvent",
		Filter:      "process.name = 'test'",
		Description: "A test analytic",
		Level:       3,
		Severity:    "Medium",
	}

	result, _, err := service.CreateAnalytic(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testUUID, result.UUID)
	assert.Equal(t, "Test Analytic", result.Name)
}

func TestAnalyticService_GetAnalytic(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterGetAnalyticMock()

	result, _, err := service.GetAnalytic(context.Background(), testUUID)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testUUID, result.UUID)
	assert.Equal(t, "Test Analytic", result.Name)
}

func TestAnalyticService_UpdateAnalytic(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterUpdateAnalyticMock()

	req := &analytic.UpdateAnalyticRequest{
		Name:        "Updated Analytic",
		InputType:   "GPFSEvent",
		Filter:      "process.name = 'updated'",
		Description: "An updated test analytic",
		Level:       5,
	}

	result, _, err := service.UpdateAnalytic(context.Background(), testUUID, req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testUUID, result.UUID)
	assert.Equal(t, "Updated Analytic", result.Name)
}

func TestAnalyticService_DeleteAnalytic(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterDeleteAnalyticMock()

	_, err := service.DeleteAnalytic(context.Background(), testUUID)

	require.NoError(t, err)
}

func TestAnalyticService_ListAnalytics(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListAnalyticsMock()

	result, _, err := service.ListAnalytics(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, testUUID, result[0].UUID)
	assert.Equal(t, "Test Analytic", result[0].Name)
}

func TestAnalyticService_ListAnalyticsLite(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListAnalyticsLiteMock()

	result, _, err := service.ListAnalyticsLite(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, testUUID, result[0].UUID)
	assert.Equal(t, "Test Analytic", result[0].Name)
}

func TestAnalyticService_ListAnalyticsNames(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListAnalyticsNamesMock()

	result, _, err := service.ListAnalyticsNames(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Test Analytic", result[0])
}

func TestAnalyticService_ListAnalyticsCategories(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListAnalyticsCategoriesMock()

	result, _, err := service.ListAnalyticsCategories(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Security", result[0].Value)
	assert.Equal(t, 5, result[0].Count)
}

func TestAnalyticService_ListAnalyticsTags(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListAnalyticsTagsMock()

	result, _, err := service.ListAnalyticsTags(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "endpoint", result[0].Value)
	assert.Equal(t, 10, result[0].Count)
}

func TestAnalyticService_ListAnalyticsFilterOptions(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListAnalyticsFilterOptionsMock()

	result, _, err := service.ListAnalyticsFilterOptions(context.Background())

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Len(t, result.Tags, 1)
	assert.Len(t, result.Categories, 1)
	assert.Equal(t, "endpoint", result.Tags[0].Value)
	assert.Equal(t, "Security", result.Categories[0].Value)
}

func TestAnalyticService_ListAnalytics_EmptyResult(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/graphql", "listAnalytics", 200, "list_analytics_empty.json")

	result, _, err := service.ListAnalytics(context.Background())

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

func TestAnalyticService_ListAnalyticsLite_EmptyResult(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/graphql", "listAnalyticsLite", 200, "list_analytics_empty.json")

	result, _, err := service.ListAnalyticsLite(context.Background())

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

func TestAnalyticService_ListAnalyticsNames_EmptyResult(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/graphql", "listAnalyticsNames", 200, "list_analytics_empty.json")

	result, _, err := service.ListAnalyticsNames(context.Background())

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

func TestAnalyticService_ListAnalyticsCategories_EmptyResult(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/graphql", "listAnalyticsCategories", 200, "list_analytics_categories_empty.json")

	result, _, err := service.ListAnalyticsCategories(context.Background())

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

func TestAnalyticService_ListAnalyticsTags_EmptyResult(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/graphql", "listAnalyticsTags", 200, "list_analytics_tags_empty.json")

	result, _, err := service.ListAnalyticsTags(context.Background())

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

func TestAnalyticService_ListAnalyticsFilterOptions_EmptyResult(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/graphql", "listAnalyticsFilterOptions", 200, "list_analytics_filter_options_empty.json")

	result, _, err := service.ListAnalyticsFilterOptions(context.Background())

	require.NoError(t, err)
	require.NotNil(t, result)
}

func TestAnalyticService_ValidationErrors(t *testing.T) {
	service, _ := setupMockService(t)

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{
			name: "CreateAnalytic nil request",
			fn: func() error {
				_, _, err := service.CreateAnalytic(context.Background(), nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "CreateAnalytic missing name",
			fn: func() error {
				_, _, err := service.CreateAnalytic(context.Background(), &analytic.CreateAnalyticRequest{
					InputType: "GPFSEvent",
					Filter:    "test",
				})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "CreateAnalytic missing inputType",
			fn: func() error {
				_, _, err := service.CreateAnalytic(context.Background(), &analytic.CreateAnalyticRequest{
					Name:   "test",
					Filter: "test",
				})
				return err
			},
			wantErr: "inputType is required",
		},
		{
			name: "CreateAnalytic missing filter",
			fn: func() error {
				_, _, err := service.CreateAnalytic(context.Background(), &analytic.CreateAnalyticRequest{
					Name:      "test",
					InputType: "GPFSEvent",
				})
				return err
			},
			wantErr: "filter is required",
		},
		{
			name: "GetAnalytic empty uuid",
			fn: func() error {
				_, _, err := service.GetAnalytic(context.Background(), "")
				return err
			},
			wantErr: "uuid is required",
		},
		{
			name: "GetAnalytic invalid uuid",
			fn: func() error {
				_, _, err := service.GetAnalytic(context.Background(), "not-a-uuid")
				return err
			},
			wantErr: "uuid must be a valid UUID",
		},
		{
			name: "UpdateAnalytic empty uuid",
			fn: func() error {
				_, _, err := service.UpdateAnalytic(context.Background(), "", &analytic.UpdateAnalyticRequest{})
				return err
			},
			wantErr: "uuid is required",
		},
		{
			name: "UpdateAnalytic invalid uuid",
			fn: func() error {
				_, _, err := service.UpdateAnalytic(context.Background(), "not-a-uuid", &analytic.UpdateAnalyticRequest{})
				return err
			},
			wantErr: "uuid must be a valid UUID",
		},
		{
			name: "DeleteAnalytic empty uuid",
			fn: func() error {
				_, err := service.DeleteAnalytic(context.Background(), "")
				return err
			},
			wantErr: "uuid is required",
		},
		{
			name: "DeleteAnalytic invalid uuid",
			fn: func() error {
				_, err := service.DeleteAnalytic(context.Background(), "not-a-uuid")
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

func TestAnalyticService_Validators(t *testing.T) {
	validUUID := "550e8400-e29b-41d4-a716-446655440000"
	assert.NoError(t, analytic.ValidateAnalyticID(validUUID))
	assert.Error(t, analytic.ValidateAnalyticID(""))
	assert.Error(t, analytic.ValidateAnalyticID("not-a-uuid"))

	assert.NoError(t, analytic.ValidateInputType("GPFSEvent"))
	assert.NoError(t, analytic.ValidateInputType("GPProcessEvent"))
	assert.NoError(t, analytic.ValidateInputType("GPGatekeeperEvent"))
	assert.Error(t, analytic.ValidateInputType("INVALID"))

	assert.NoError(t, analytic.ValidateLevel(0))
	assert.NoError(t, analytic.ValidateLevel(5))
	assert.NoError(t, analytic.ValidateLevel(10))
	assert.Error(t, analytic.ValidateLevel(-1))
	assert.Error(t, analytic.ValidateLevel(11))

	assert.NoError(t, analytic.ValidateSeverity("High"))
	assert.NoError(t, analytic.ValidateSeverity("Medium"))
	assert.NoError(t, analytic.ValidateSeverity("Low"))
	assert.NoError(t, analytic.ValidateSeverity("Informational"))
	assert.Error(t, analytic.ValidateSeverity("INVALID"))

	severity := "High"
	invalidSeverity := "INVALID"
	assert.NoError(t, analytic.ValidateCreateAnalyticRequest(&analytic.CreateAnalyticRequest{
		InputType: "GPFSEvent",
		Level:     5,
		Severity:  "High",
	}))
	assert.NoError(t, analytic.ValidateCreateAnalyticRequest(nil))
	assert.Error(t, analytic.ValidateCreateAnalyticRequest(&analytic.CreateAnalyticRequest{
		InputType: "INVALID",
		Level:     5,
		Severity:  "High",
	}))
	assert.Error(t, analytic.ValidateCreateAnalyticRequest(&analytic.CreateAnalyticRequest{
		InputType: "GPFSEvent",
		Level:     11,
		Severity:  "High",
	}))
	assert.Error(t, analytic.ValidateCreateAnalyticRequest(&analytic.CreateAnalyticRequest{
		InputType: "GPFSEvent",
		Level:     5,
		Severity:  "INVALID",
	}))

	assert.NoError(t, analytic.ValidateUpdateAnalyticRequest(&analytic.UpdateAnalyticRequest{
		InputType: "GPFSEvent",
		Level:     5,
		Severity:  &severity,
	}))
	assert.NoError(t, analytic.ValidateUpdateAnalyticRequest(nil))
	assert.Error(t, analytic.ValidateUpdateAnalyticRequest(&analytic.UpdateAnalyticRequest{
		InputType: "INVALID",
	}))
	assert.Error(t, analytic.ValidateUpdateAnalyticRequest(&analytic.UpdateAnalyticRequest{
		Level: 11,
	}))
	assert.Error(t, analytic.ValidateUpdateAnalyticRequest(&analytic.UpdateAnalyticRequest{
		Severity: &invalidSeverity,
	}))
}
