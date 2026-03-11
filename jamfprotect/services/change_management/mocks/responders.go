package mocks

import (
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	coremocks "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/mocks"
)

// ChangeManagementMock provides mock responses for the ChangeManagement service GraphQL operations.
type ChangeManagementMock struct {
	*coremocks.GenericGraphQLMock
}

// NewChangeManagementMock creates a new ChangeManagementMock instance.
func NewChangeManagementMock() *ChangeManagementMock {
	return &ChangeManagementMock{
		GenericGraphQLMock: coremocks.NewGenericGraphQLMock(coremocks.GenericGraphQLMockConfig{
			Name: "ChangeManagementMock",
		}),
	}
}

// RegisterMocks registers all successful response mocks for change management operations.
func (m *ChangeManagementMock) RegisterMocks() {
	m.RegisterGetConfigFreezeMock()
	m.RegisterUpdateConfigFreezeMock()
}

// RegisterGetConfigFreezeMock registers a success mock for getConfigFreeze.
func (m *ChangeManagementMock) RegisterGetConfigFreezeMock() {
	m.Register(client.EndpointGraphQL, "getConfigFreeze", 200, "get_config_freeze_success.json")
}

// RegisterUpdateConfigFreezeMock registers a success mock for updateOrganizationConfigFreeze.
func (m *ChangeManagementMock) RegisterUpdateConfigFreezeMock() {
	m.Register(client.EndpointApp, "updateOrganizationConfigFreeze", 200, "update_config_freeze_success.json")
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock.
func (m *ChangeManagementMock) RegisterUnauthorizedErrorMock() {
	m.RegisterError(client.EndpointGraphQL, "getConfigFreeze", 401, "error_unauthorized.json", "Jamf Protect API error: unauthorized")
}
