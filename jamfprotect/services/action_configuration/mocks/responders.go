package mocks

import (
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	coremocks "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/mocks"
)

// ActionConfigMock provides mock responses for the ActionConfiguration service GraphQL operations.
// All operations POST to the /app endpoint and are distinguished by operation name in the request body.
type ActionConfigMock struct {
	*coremocks.GenericGraphQLMock
}

// NewActionConfigMock creates a new ActionConfigMock instance.
func NewActionConfigMock() *ActionConfigMock {
	return &ActionConfigMock{
		GenericGraphQLMock: coremocks.NewGenericGraphQLMock(coremocks.GenericGraphQLMockConfig{
			Name: "ActionConfigMock",
		}),
	}
}

// RegisterMocks registers all successful response mocks for action config operations.
func (m *ActionConfigMock) RegisterMocks() {
	m.RegisterCreateActionConfigMock()
	m.RegisterGetActionConfigMock()
	m.RegisterUpdateActionConfigMock()
	m.RegisterDeleteActionConfigMock()
	m.RegisterListActionConfigsMock()
	m.RegisterListActionConfigNamesMock()
}

// RegisterErrorMocks registers error response mocks.
func (m *ActionConfigMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterUnauthorizedErrorMock()
}

// RegisterCreateActionConfigMock registers a success mock for createActionConfigs.
func (m *ActionConfigMock) RegisterCreateActionConfigMock() {
	m.Register(client.EndpointApp, "createActionConfigs", 200, "create_action_config_success.json")
}

// RegisterGetActionConfigMock registers a success mock for getActionConfigs.
func (m *ActionConfigMock) RegisterGetActionConfigMock() {
	m.Register(client.EndpointApp, "getActionConfigs", 200, "get_action_config_success.json")
}

// RegisterUpdateActionConfigMock registers a success mock for updateActionConfigs.
func (m *ActionConfigMock) RegisterUpdateActionConfigMock() {
	m.Register(client.EndpointApp, "updateActionConfigs", 200, "update_action_config_success.json")
}

// RegisterDeleteActionConfigMock registers a success mock for deleteActionConfigs.
func (m *ActionConfigMock) RegisterDeleteActionConfigMock() {
	m.Register(client.EndpointApp, "deleteActionConfigs", 200, "delete_action_config_success.json")
}

// RegisterListActionConfigsMock registers a success mock for listActionConfigs.
func (m *ActionConfigMock) RegisterListActionConfigsMock() {
	m.Register(client.EndpointApp, "listActionConfigs", 200, "list_action_configs_success.json")
}

// RegisterListActionConfigNamesMock registers a success mock for listActionConfigNames.
func (m *ActionConfigMock) RegisterListActionConfigNamesMock() {
	m.Register(client.EndpointApp, "listActionConfigNames", 200, "list_action_config_names_success.json")
}

// RegisterNotFoundErrorMock registers a not-found error mock.
func (m *ActionConfigMock) RegisterNotFoundErrorMock() {
	m.RegisterError(client.EndpointApp, "getActionConfigs", 200, "error_not_found.json", "graphql operation failed: ActionConfig not found")
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock.
func (m *ActionConfigMock) RegisterUnauthorizedErrorMock() {
	m.RegisterError(client.EndpointApp, "getActionConfigs", 401, "error_unauthorized.json", "Jamf Protect API error: unauthorized")
}
