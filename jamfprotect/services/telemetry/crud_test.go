package telemetry_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/telemetry"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/telemetry/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*telemetry.Service, *mocks.TelemetryMock) {
	t.Helper()
	mock := mocks.NewTelemetryMock()
	return telemetry.NewService(mock), mock
}

func TestTelemetryService_CreateTelemetryV2(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterCreateTelemetryV2Mock()

	req := &telemetry.CreateTelemetryV2Request{
		Name:        "Test Telemetry V2",
		Description: "A test telemetry v2",
		LogFiles:    []string{},
	}

	result, _, err := service.CreateTelemetryV2(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Test Telemetry V2", result.Name)
}

func TestTelemetryService_GetTelemetryV2(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterGetTelemetryV2Mock()

	result, _, err := service.GetTelemetryV2(context.Background(), "test-id-1234")

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Test Telemetry V2", result.Name)
}

func TestTelemetryService_UpdateTelemetryV2(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterUpdateTelemetryV2Mock()

	req := &telemetry.UpdateTelemetryV2Request{
		Name:        "Updated Telemetry V2",
		Description: "An updated telemetry v2",
		LogFiles:    []string{},
	}

	result, _, err := service.UpdateTelemetryV2(context.Background(), "test-id-1234", req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Updated Telemetry V2", result.Name)
}

func TestTelemetryService_DeleteTelemetryV2(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterDeleteTelemetryV2Mock()

	_, err := service.DeleteTelemetryV2(context.Background(), "test-id-1234")

	require.NoError(t, err)
}

func TestTelemetryService_ListTelemetriesV2(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListTelemetriesV2Mock()

	result, _, err := service.ListTelemetriesV2(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "test-id-1234", result[0].ID)
	assert.Equal(t, "Test Telemetry V2", result[0].Name)
}

func TestTelemetryService_ListTelemetriesCombined(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListTelemetriesCombinedMock()

	result, _, err := service.ListTelemetriesCombined(context.Background(), false)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Len(t, result.Telemetries, 1)
	assert.Equal(t, "tel-id-1", result.Telemetries[0].ID)
	assert.Equal(t, "Test Telemetry V1", result.Telemetries[0].Name)
	assert.Len(t, result.TelemetriesV2, 1)
	assert.Equal(t, "telv2-id-1", result.TelemetriesV2[0].ID)
	assert.Equal(t, "Test Telemetry V2", result.TelemetriesV2[0].Name)
}

func TestTelemetryService_ListTelemetriesV2_EmptyResult(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/app", "listTelemetriesV2", 200, "list_telemetries_v2_empty.json")

	result, _, err := service.ListTelemetriesV2(context.Background())

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

func TestTelemetryService_ListTelemetriesV2_Error(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterError("/app", "listTelemetriesV2", 500, "", "Internal Server Error")

	result, _, err := service.ListTelemetriesV2(context.Background())

	require.Error(t, err)
	assert.Nil(t, result)
}

func TestTelemetryService_ValidationErrors(t *testing.T) {
	service, _ := setupMockService(t)

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{
			name: "CreateTelemetryV2 nil request",
			fn: func() error {
				_, _, err := service.CreateTelemetryV2(context.Background(), nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "CreateTelemetryV2 missing name",
			fn: func() error {
				_, _, err := service.CreateTelemetryV2(context.Background(), &telemetry.CreateTelemetryV2Request{
					LogFiles: []string{},
				})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "CreateTelemetryV2 nil logFiles",
			fn: func() error {
				_, _, err := service.CreateTelemetryV2(context.Background(), &telemetry.CreateTelemetryV2Request{
					Name: "test",
				})
				return err
			},
			wantErr: "logFiles is required",
		},
		{
			name: "GetTelemetryV2 empty id",
			fn: func() error {
				_, _, err := service.GetTelemetryV2(context.Background(), "")
				return err
			},
			wantErr: "id is required",
		},
		{
			name: "UpdateTelemetryV2 nil request",
			fn: func() error {
				_, _, err := service.UpdateTelemetryV2(context.Background(), "test-id", nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "UpdateTelemetryV2 missing name",
			fn: func() error {
				_, _, err := service.UpdateTelemetryV2(context.Background(), "test-id", &telemetry.UpdateTelemetryV2Request{
					LogFiles: []string{},
				})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "UpdateTelemetryV2 nil logFiles",
			fn: func() error {
				_, _, err := service.UpdateTelemetryV2(context.Background(), "test-id", &telemetry.UpdateTelemetryV2Request{
					Name: "test",
				})
				return err
			},
			wantErr: "logFiles is required",
		},
		{
			name: "UpdateTelemetryV2 empty id",
			fn: func() error {
				_, _, err := service.UpdateTelemetryV2(context.Background(), "", &telemetry.UpdateTelemetryV2Request{
					Name:     "test",
					LogFiles: []string{},
				})
				return err
			},
			wantErr: "id is required",
		},
		{
			name: "UpdateTelemetryV2 nil request",
			fn: func() error {
				_, _, err := service.UpdateTelemetryV2(context.Background(), "test-id", nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "DeleteTelemetryV2 empty id",
			fn: func() error {
				_, err := service.DeleteTelemetryV2(context.Background(), "")
				return err
			},
			wantErr: "id is required",
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

func TestTelemetryService_Validators(t *testing.T) {
	assert.NoError(t, telemetry.ValidateTelemetryV2ID("test-id"))
	assert.NoError(t, telemetry.ValidateTelemetryV2ID(""))

	assert.NoError(t, telemetry.ValidateCreateTelemetryV2Request(&telemetry.CreateTelemetryV2Request{}))
	assert.NoError(t, telemetry.ValidateCreateTelemetryV2Request(nil))

	assert.NoError(t, telemetry.ValidateUpdateTelemetryV2Request(&telemetry.UpdateTelemetryV2Request{}))
	assert.NoError(t, telemetry.ValidateUpdateTelemetryV2Request(nil))
}
