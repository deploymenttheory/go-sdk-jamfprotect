package actionconfiguration_test

import (
	"context"
	"testing"

	actionconfiguration "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/action_configuration"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/action_configuration/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*actionconfiguration.Service, *mocks.ActionConfigMock) {
	t.Helper()
	mock := mocks.NewActionConfigMock()
	return actionconfiguration.NewService(mock), mock
}

func TestActionConfigService_CreateActionConfig(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterCreateActionConfigMock()

	req := &actionconfiguration.CreateActionConfigRequest{
		Name:        "Test Action Config",
		Description: "A test action configuration",
		AlertConfig: map[string]any{
			"type": "alert",
		},
	}

	result, _, err := service.CreateActionConfig(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Test Action Config", result.Name)
}

func TestActionConfigService_GetActionConfig(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterGetActionConfigMock()

	result, _, err := service.GetActionConfig(context.Background(), "test-id-1234")

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Test Action Config", result.Name)
}

func TestActionConfigService_UpdateActionConfig(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterUpdateActionConfigMock()

	req := &actionconfiguration.UpdateActionConfigRequest{
		Name:        "Updated Action Config",
		Description: "An updated action configuration",
		AlertConfig: map[string]any{
			"type": "alert",
		},
	}

	result, _, err := service.UpdateActionConfig(context.Background(), "test-id-1234", req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Updated Action Config", result.Name)
}

func TestActionConfigService_DeleteActionConfig(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterDeleteActionConfigMock()

	_, err := service.DeleteActionConfig(context.Background(), "test-id-1234")

	require.NoError(t, err)
}

func TestActionConfigService_ListActionConfigs(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListActionConfigsMock()

	result, _, err := service.ListActionConfigs(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "test-id-1234", result[0].ID)
	assert.Equal(t, "Test Action Config", result[0].Name)
}

func TestActionConfigService_ListActionConfigNames(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListActionConfigNamesMock()

	result, _, err := service.ListActionConfigNames(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Test Action Config", result[0])
}

func TestActionConfigService_ListActionConfigs_EmptyResult(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/app", "listActionConfigs", 200, "list_action_configs_empty.json")

	result, _, err := service.ListActionConfigs(context.Background())

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

func TestActionConfigService_ListActionConfigNames_EmptyResult(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/app", "listActionConfigNames", 200, "list_action_config_names_empty.json")

	result, _, err := service.ListActionConfigNames(context.Background())

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

func TestActionConfigService_ValidationErrors(t *testing.T) {
	service, _ := setupMockService(t)

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{
			name: "CreateActionConfig nil request",
			fn: func() error {
				_, _, err := service.CreateActionConfig(context.Background(), nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "CreateActionConfig missing name",
			fn: func() error {
				_, _, err := service.CreateActionConfig(context.Background(), &actionconfiguration.CreateActionConfigRequest{
					AlertConfig: map[string]any{"type": "alert"},
				})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "CreateActionConfig missing alertConfig",
			fn: func() error {
				_, _, err := service.CreateActionConfig(context.Background(), &actionconfiguration.CreateActionConfigRequest{
					Name: "test",
				})
				return err
			},
			wantErr: "alertConfig is required",
		},
		{
			name: "GetActionConfig empty id",
			fn: func() error {
				_, _, err := service.GetActionConfig(context.Background(), "")
				return err
			},
			wantErr: "id is required",
		},
		{
			name: "UpdateActionConfig empty id",
			fn: func() error {
				_, _, err := service.UpdateActionConfig(context.Background(), "", &actionconfiguration.UpdateActionConfigRequest{
					Name:        "test",
					AlertConfig: map[string]any{"type": "alert"},
				})
				return err
			},
			wantErr: "id is required",
		},
		{
			name: "UpdateActionConfig nil request",
			fn: func() error {
				_, _, err := service.UpdateActionConfig(context.Background(), "test-id", nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "UpdateActionConfig missing name",
			fn: func() error {
				_, _, err := service.UpdateActionConfig(context.Background(), "test-id", &actionconfiguration.UpdateActionConfigRequest{
					AlertConfig: map[string]any{"type": "alert"},
				})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "UpdateActionConfig missing alertConfig",
			fn: func() error {
				_, _, err := service.UpdateActionConfig(context.Background(), "test-id", &actionconfiguration.UpdateActionConfigRequest{
					Name: "test",
				})
				return err
			},
			wantErr: "alertConfig is required",
		},
		{
			name: "DeleteActionConfig empty id",
			fn: func() error {
				_, err := service.DeleteActionConfig(context.Background(), "")
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

func TestActionConfigService_Validators(t *testing.T) {
	assert.NoError(t, actionconfiguration.ValidateActionConfigID("test-id"))
	assert.NoError(t, actionconfiguration.ValidateActionConfigID(""))

	assert.NoError(t, actionconfiguration.ValidateCreateActionConfigRequest(&actionconfiguration.CreateActionConfigRequest{}))
	assert.NoError(t, actionconfiguration.ValidateCreateActionConfigRequest(nil))

	assert.NoError(t, actionconfiguration.ValidateUpdateActionConfigRequest(&actionconfiguration.UpdateActionConfigRequest{}))
	assert.NoError(t, actionconfiguration.ValidateUpdateActionConfigRequest(nil))
}
