package plan_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/plan"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/plan/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*plan.Service, *mocks.PlanMock) {
	t.Helper()
	mock := mocks.NewPlanMock()
	return plan.NewService(mock), mock
}

func TestPlanService_CreatePlan(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterCreatePlanMock()

	req := &plan.CreatePlanRequest{
		Name:          "Test Plan",
		Description:   "A test plan",
		ActionConfigs: "action-config-123",
		CommsConfig: plan.CommsConfigInput{
			Protocol: "mqtt",
		},
	}

	result, _, err := service.CreatePlan(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Test Plan", result.Name)
}

func TestPlanService_GetPlan(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterGetPlanMock()

	result, _, err := service.GetPlan(context.Background(), "test-id-1234")

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Test Plan", result.Name)
}

func TestPlanService_UpdatePlan(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterUpdatePlanMock()

	req := &plan.UpdatePlanRequest{
		Name:          "Updated Plan",
		Description:   "An updated plan",
		ActionConfigs: "action-config-123",
		CommsConfig: plan.CommsConfigInput{
			Protocol: "mqtt",
		},
	}

	result, _, err := service.UpdatePlan(context.Background(), "test-id-1234", req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Updated Plan", result.Name)
}

func TestPlanService_DeletePlan(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterDeletePlanMock()

	_, err := service.DeletePlan(context.Background(), "test-id-1234")

	require.NoError(t, err)
}

func TestPlanService_ListPlans(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListPlansMock()

	result, _, err := service.ListPlans(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "test-id-1234", result[0].ID)
	assert.Equal(t, "Test Plan", result[0].Name)
}

func TestPlanService_ListPlanNames(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListPlanNamesMock()

	result, _, err := service.ListPlanNames(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Test Plan", result[0])
}

func TestPlanService_GetPlanConfigurationAndSetOptions(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterGetPlanConfigurationAndSetOptionsMock()

	req := &plan.GetPlanConfigurationAndSetOptionsRequest{
		RBACActionConfigs: true,
		RBACTelemetry:     true,
		RBACUSBControlSet: true,
		RBACExceptionSet:  true,
		RBACAnalyticSet:   true,
	}

	result, _, err := service.GetPlanConfigurationAndSetOptions(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Len(t, result.ActionConfigs, 1)
	assert.Equal(t, "ac-id-1", result.ActionConfigs[0].ID)
	assert.Equal(t, "Action Config 1", result.ActionConfigs[0].Name)
	assert.Len(t, result.ExceptionSets, 1)
	assert.Equal(t, "exc-uuid-1", result.ExceptionSets[0].UUID)
	assert.Len(t, result.ManagedAnalyticSets, 1)
	assert.Equal(t, "mas-uuid-1", result.ManagedAnalyticSets[0].UUID)
}

func TestPlanService_ValidationErrors(t *testing.T) {
	service, _ := setupMockService(t)

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{
			name: "CreatePlan nil request",
			fn: func() error {
				_, _, err := service.CreatePlan(context.Background(), nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "CreatePlan missing name",
			fn: func() error {
				_, _, err := service.CreatePlan(context.Background(), &plan.CreatePlanRequest{
					ActionConfigs: "action-123",
					CommsConfig:   plan.CommsConfigInput{Protocol: "mqtt"},
				})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "CreatePlan missing actionConfigs",
			fn: func() error {
				_, _, err := service.CreatePlan(context.Background(), &plan.CreatePlanRequest{
					Name:        "test",
					CommsConfig: plan.CommsConfigInput{Protocol: "mqtt"},
				})
				return err
			},
			wantErr: "actionConfigs is required",
		},
		{
			name: "GetPlan empty id",
			fn: func() error {
				_, _, err := service.GetPlan(context.Background(), "")
				return err
			},
			wantErr: "id is required",
		},
		{
			name: "DeletePlan empty id",
			fn: func() error {
				_, err := service.DeletePlan(context.Background(), "")
				return err
			},
			wantErr: "id is required",
		},
		{
			name: "GetPlanConfigurationAndSetOptions nil request",
			fn: func() error {
				_, _, err := service.GetPlanConfigurationAndSetOptions(context.Background(), nil)
				return err
			},
			wantErr: "request is required",
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
