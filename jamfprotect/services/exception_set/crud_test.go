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
			name: "UpdateExceptionSet empty uuid",
			fn: func() error {
				_, _, err := service.UpdateExceptionSet(context.Background(), "", &exceptionset.UpdateExceptionSetRequest{Name: "test"})
				return err
			},
			wantErr: "uuid is required",
		},
		{
			name: "UpdateExceptionSet nil request",
			fn: func() error {
				_, _, err := service.UpdateExceptionSet(context.Background(), testUUID, nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "UpdateExceptionSet missing name",
			fn: func() error {
				_, _, err := service.UpdateExceptionSet(context.Background(), testUUID, &exceptionset.UpdateExceptionSetRequest{})
				return err
			},
			wantErr: "name is required",
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

func TestExceptionSetService_CreateWithExceptions(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterCreateExceptionSetMock()

	req := &exceptionset.CreateExceptionSetRequest{
		Name:        "Test with Exceptions",
		Description: "Exception set with various exception types",
		Exceptions: []exceptionset.ExceptionInput{
			{
				Type:           "Path",
				Value:          "/usr/bin/test",
				IgnoreActivity: "Analytics",
			},
			{
				Type:  "AppSigningInfo",
				Value: "com.example.app",
				AppSigningInfo: &exceptionset.AppSigningInfoInput{
					AppId:  "app123",
					TeamId: "team456",
				},
				IgnoreActivity: "ThreatPrevention",
				AnalyticTypes:  []string{"MALWARE", "SUSPICIOUS"},
			},
			{
				Type:           "TeamId",
				Value:          "signing-id",
				IgnoreActivity: "Telemetry",
				AnalyticUuid:   "analytic-uuid-123",
				AnalyticTypes:  []string{"MALWARE"},
			},
		},
	}

	result, _, err := service.CreateExceptionSet(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
}

func TestExceptionSetService_CreateWithEsExceptions(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterCreateExceptionSetMock()

	req := &exceptionset.CreateExceptionSetRequest{
		Name:        "Test with ES Exceptions",
		Description: "Exception set with ES exception types",
		EsExceptions: []exceptionset.EsExceptionInput{
			{
				Type:           "Path",
				Value:          "/usr/bin/test",
				IgnoreActivity: "TelemetryV2",
			},
			{
				Type:  "AppSigningInfo",
				Value: "com.example.app",
				AppSigningInfo: &exceptionset.AppSigningInfoInput{
					AppId:  "app123",
					TeamId: "team456",
				},
				IgnoreActivity:    "ThreatPrevention",
				IgnoreListType:    "ALLOW",
				IgnoreListSubType: "PROCESS",
				EventType:         "ES_EVENT_TYPE_AUTH_EXEC",
			},
		},
	}

	result, _, err := service.CreateExceptionSet(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
}

func TestExceptionSetService_UpdateWithExceptions(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterUpdateExceptionSetMock()

	req := &exceptionset.UpdateExceptionSetRequest{
		Name:        "Updated with Exceptions",
		Description: "Updated exception set",
		Exceptions: []exceptionset.ExceptionInput{
			{
				Type:           "Path",
				Value:          "/usr/bin/updated",
				IgnoreActivity: "Analytics",
				AppSigningInfo: &exceptionset.AppSigningInfoInput{
					AppId:  "updated-app",
					TeamId: "updated-team",
				},
				AnalyticTypes: []string{"SUSPICIOUS"},
				AnalyticUuid:  "updated-uuid",
			},
		},
		EsExceptions: []exceptionset.EsExceptionInput{
			{
				Type:              "Executable",
				Value:             "/usr/bin/es-updated",
				IgnoreActivity:    "Telemetry",
				IgnoreListType:    "DENY",
				IgnoreListSubType: "FILE",
				EventType:         "ES_EVENT_TYPE_AUTH_OPEN",
			},
		},
	}

	result, _, err := service.UpdateExceptionSet(context.Background(), testUUID, req)

	require.NoError(t, err)
	require.NotNil(t, result)
}

func TestExceptionSetService_Validators(t *testing.T) {
	assert.NoError(t, exceptionset.ValidateExceptionSetID("test-id"))
	assert.NoError(t, exceptionset.ValidateExceptionSetID(""))

	validUUID := "550e8400-e29b-41d4-a716-446655440000"
	assert.NoError(t, exceptionset.ValidateExceptionSetUUID(validUUID))
	assert.Error(t, exceptionset.ValidateExceptionSetUUID(""))
	assert.Error(t, exceptionset.ValidateExceptionSetUUID("not-a-uuid"))

	assert.NoError(t, exceptionset.ValidateExceptionType("User"))
	assert.NoError(t, exceptionset.ValidateExceptionType("Path"))
	assert.Error(t, exceptionset.ValidateExceptionType("INVALID"))

	assert.NoError(t, exceptionset.ValidateIgnoreActivity("Analytics"))
	assert.NoError(t, exceptionset.ValidateIgnoreActivity("ThreatPrevention"))
	assert.Error(t, exceptionset.ValidateIgnoreActivity("INVALID"))

	assert.NoError(t, exceptionset.ValidateExceptionInput(exceptionset.ExceptionInput{
		Type:           "Path",
		Value:          "/test",
		IgnoreActivity: "Analytics",
	}))
	assert.Error(t, exceptionset.ValidateExceptionInput(exceptionset.ExceptionInput{
		Type:           "INVALID",
		Value:          "/test",
		IgnoreActivity: "Analytics",
	}))

	assert.NoError(t, exceptionset.ValidateEsExceptionInput(exceptionset.EsExceptionInput{
		Type:           "Path",
		Value:          "/test",
		IgnoreActivity: "Telemetry",
	}))
	assert.Error(t, exceptionset.ValidateEsExceptionInput(exceptionset.EsExceptionInput{
		Type:           "Path",
		Value:          "/test",
		IgnoreActivity: "INVALID",
	}))

	assert.NoError(t, exceptionset.ValidateCreateExceptionSetRequest(&exceptionset.CreateExceptionSetRequest{
		Name: "test",
		Exceptions: []exceptionset.ExceptionInput{
			{Type: "Path", Value: "/test", IgnoreActivity: "Analytics"},
		},
		EsExceptions: []exceptionset.EsExceptionInput{
			{Type: "Path", Value: "/test", IgnoreActivity: "Telemetry"},
		},
	}))
	assert.NoError(t, exceptionset.ValidateCreateExceptionSetRequest(nil))
	assert.Error(t, exceptionset.ValidateCreateExceptionSetRequest(&exceptionset.CreateExceptionSetRequest{
		Exceptions: []exceptionset.ExceptionInput{
			{Type: "INVALID", Value: "/test", IgnoreActivity: "Analytics"},
		},
	}))

	assert.NoError(t, exceptionset.ValidateUpdateExceptionSetRequest(&exceptionset.UpdateExceptionSetRequest{
		Name: "test",
		Exceptions: []exceptionset.ExceptionInput{
			{Type: "Path", Value: "/test", IgnoreActivity: "Analytics"},
		},
	}))
	assert.NoError(t, exceptionset.ValidateUpdateExceptionSetRequest(nil))
	assert.Error(t, exceptionset.ValidateUpdateExceptionSetRequest(&exceptionset.UpdateExceptionSetRequest{
		EsExceptions: []exceptionset.EsExceptionInput{
			{Type: "Path", Value: "/test", IgnoreActivity: "INVALID"},
		},
	}))
}
