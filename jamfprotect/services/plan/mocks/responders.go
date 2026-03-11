package mocks

import (
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	coremocks "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/mocks"
)

// PlanMock provides mock responses for the Plan service GraphQL operations.
// All operations POST to the /app GraphQL endpoint and are distinguished by operation name
// in the request body.
type PlanMock struct {
	*coremocks.GenericGraphQLMock
}

// NewPlanMock creates a new PlanMock instance.
func NewPlanMock() *PlanMock {
	return &PlanMock{
		GenericGraphQLMock: coremocks.NewGenericGraphQLMock(coremocks.GenericGraphQLMockConfig{
			Name: "PlanMock",
		}),
	}
}

// RegisterMocks registers all successful response mocks for plan operations.
func (m *PlanMock) RegisterMocks() {
	m.RegisterCreatePlanMock()
	m.RegisterGetPlanMock()
	m.RegisterUpdatePlanMock()
	m.RegisterDeletePlanMock()
	m.RegisterListPlansMock()
	m.RegisterListPlanNamesMock()
	m.RegisterGetPlanConfigurationAndSetOptionsMock()
}

// RegisterErrorMocks registers error response mocks.
func (m *PlanMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterUnauthorizedErrorMock()
}

// RegisterCreatePlanMock registers a success mock for createPlan.
func (m *PlanMock) RegisterCreatePlanMock() {
	m.Register(client.EndpointApp, "createPlan", 200, "create_plan_success.json")
}

// RegisterGetPlanMock registers a success mock for getPlan.
func (m *PlanMock) RegisterGetPlanMock() {
	m.Register(client.EndpointApp, "getPlan", 200, "get_plan_success.json")
}

// RegisterUpdatePlanMock registers a success mock for updatePlan.
func (m *PlanMock) RegisterUpdatePlanMock() {
	m.Register(client.EndpointApp, "updatePlan", 200, "update_plan_success.json")
}

// RegisterDeletePlanMock registers a success mock for deletePlan.
func (m *PlanMock) RegisterDeletePlanMock() {
	m.Register(client.EndpointApp, "deletePlan", 200, "delete_plan_success.json")
}

// RegisterListPlansMock registers a success mock for listPlans.
func (m *PlanMock) RegisterListPlansMock() {
	m.Register(client.EndpointApp, "listPlans", 200, "list_plans_success.json")
}

// RegisterListPlanNamesMock registers a success mock for listPlanNames.
func (m *PlanMock) RegisterListPlanNamesMock() {
	m.Register(client.EndpointApp, "listPlanNames", 200, "list_plan_names_success.json")
}

// RegisterGetPlanConfigurationAndSetOptionsMock registers a success mock for getPlanConfigurationAndSetOptions.
func (m *PlanMock) RegisterGetPlanConfigurationAndSetOptionsMock() {
	m.Register(client.EndpointApp, "getPlanConfigurationAndSetOptions", 200, "get_plan_configuration_and_set_options_success.json")
}

// RegisterNotFoundErrorMock registers a not-found error mock.
func (m *PlanMock) RegisterNotFoundErrorMock() {
	m.RegisterError(client.EndpointApp, "getPlan", 200, "error_not_found.json", "graphql operation failed: Plan not found")
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock.
func (m *PlanMock) RegisterUnauthorizedErrorMock() {
	m.RegisterError(client.EndpointApp, "getPlan", 401, "error_unauthorized.json", "Jamf Protect API error: unauthorized")
}
