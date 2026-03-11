package mocks

import (
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	coremocks "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/mocks"
)

// UserMock provides mock responses for the User service GraphQL operations.
// Mutations (create/update/delete) POST to the /app endpoint; queries POST to the /graphql endpoint.
// Operations are distinguished by operation name in the request body.
type UserMock struct {
	*coremocks.GenericGraphQLMock
}

// NewUserMock creates a new UserMock instance.
func NewUserMock() *UserMock {
	return &UserMock{
		GenericGraphQLMock: coremocks.NewGenericGraphQLMock(coremocks.GenericGraphQLMockConfig{
			Name: "UserMock",
		}),
	}
}

// RegisterMocks registers all successful response mocks for user operations.
func (m *UserMock) RegisterMocks() {
	m.RegisterCreateUserMock()
	m.RegisterGetUserMock()
	m.RegisterUpdateUserMock()
	m.RegisterDeleteUserMock()
	m.RegisterListUsersMock()
}

// RegisterErrorMocks registers error response mocks.
func (m *UserMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterUnauthorizedErrorMock()
}

// RegisterCreateUserMock registers a success mock for createUser.
func (m *UserMock) RegisterCreateUserMock() {
	m.Register(client.EndpointApp, "createUser", 200, "create_user_success.json")
}

// RegisterGetUserMock registers a success mock for getUser.
func (m *UserMock) RegisterGetUserMock() {
	m.Register(client.EndpointGraphQL, "getUser", 200, "get_user_success.json")
}

// RegisterUpdateUserMock registers a success mock for updateUser.
func (m *UserMock) RegisterUpdateUserMock() {
	m.Register(client.EndpointApp, "updateUser", 200, "update_user_success.json")
}

// RegisterDeleteUserMock registers a success mock for deleteUser.
func (m *UserMock) RegisterDeleteUserMock() {
	m.Register(client.EndpointApp, "deleteUser", 200, "delete_user_success.json")
}

// RegisterListUsersMock registers a success mock for listUsers.
func (m *UserMock) RegisterListUsersMock() {
	m.Register(client.EndpointGraphQL, "listUsers", 200, "list_users_success.json")
}

// RegisterNotFoundErrorMock registers a not-found error mock.
func (m *UserMock) RegisterNotFoundErrorMock() {
	m.RegisterError(client.EndpointGraphQL, "getUser", 200, "error_not_found.json", "graphql operation failed: User not found")
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock.
func (m *UserMock) RegisterUnauthorizedErrorMock() {
	m.RegisterError(client.EndpointGraphQL, "getUser", 401, "error_unauthorized.json", "Jamf Protect API error: unauthorized")
}
