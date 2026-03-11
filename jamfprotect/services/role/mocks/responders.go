package mocks

import (
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	coremocks "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/mocks"
)

// RoleMock provides mock responses for the Role service GraphQL operations.
// Mutations (create/update/delete) POST to the /app endpoint; queries POST to the /graphql endpoint.
// Operations are distinguished by operation name in the request body.
type RoleMock struct {
	*coremocks.GenericGraphQLMock
}

// NewRoleMock creates a new RoleMock instance.
func NewRoleMock() *RoleMock {
	return &RoleMock{
		GenericGraphQLMock: coremocks.NewGenericGraphQLMock(coremocks.GenericGraphQLMockConfig{
			Name: "RoleMock",
		}),
	}
}

// RegisterMocks registers all successful response mocks for role operations.
func (m *RoleMock) RegisterMocks() {
	m.RegisterCreateRoleMock()
	m.RegisterGetRoleMock()
	m.RegisterUpdateRoleMock()
	m.RegisterDeleteRoleMock()
	m.RegisterListRolesMock()
}

// RegisterErrorMocks registers error response mocks.
func (m *RoleMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterUnauthorizedErrorMock()
}

// RegisterCreateRoleMock registers a success mock for createRole.
func (m *RoleMock) RegisterCreateRoleMock() {
	m.Register(client.EndpointApp, "createRole", 200, "create_role_success.json")
}

// RegisterGetRoleMock registers a success mock for getRole.
func (m *RoleMock) RegisterGetRoleMock() {
	m.Register(client.EndpointGraphQL, "getRole", 200, "get_role_success.json")
}

// RegisterUpdateRoleMock registers a success mock for updateRole.
func (m *RoleMock) RegisterUpdateRoleMock() {
	m.Register(client.EndpointApp, "updateRole", 200, "update_role_success.json")
}

// RegisterDeleteRoleMock registers a success mock for deleteRole.
func (m *RoleMock) RegisterDeleteRoleMock() {
	m.Register(client.EndpointApp, "deleteRole", 200, "delete_role_success.json")
}

// RegisterListRolesMock registers a success mock for listRoles.
func (m *RoleMock) RegisterListRolesMock() {
	m.Register(client.EndpointGraphQL, "listRoles", 200, "list_roles_success.json")
}

// RegisterNotFoundErrorMock registers a not-found error mock.
func (m *RoleMock) RegisterNotFoundErrorMock() {
	m.RegisterError(client.EndpointGraphQL, "getRole", 200, "error_not_found.json", "graphql operation failed: Role not found")
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock.
func (m *RoleMock) RegisterUnauthorizedErrorMock() {
	m.RegisterError(client.EndpointGraphQL, "getRole", 401, "error_unauthorized.json", "Jamf Protect API error: unauthorized")
}
