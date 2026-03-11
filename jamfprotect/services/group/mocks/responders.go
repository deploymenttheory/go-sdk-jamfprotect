package mocks

import (
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	coremocks "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/mocks"
)

// GroupMock provides mock responses for the Group service GraphQL operations.
// Mutations (create/update/delete) POST to the /app endpoint; queries POST to the /graphql endpoint.
// Operations are distinguished by operation name in the request body.
type GroupMock struct {
	*coremocks.GenericGraphQLMock
}

// NewGroupMock creates a new GroupMock instance.
func NewGroupMock() *GroupMock {
	return &GroupMock{
		GenericGraphQLMock: coremocks.NewGenericGraphQLMock(coremocks.GenericGraphQLMockConfig{
			Name: "GroupMock",
		}),
	}
}

// RegisterMocks registers all successful response mocks for group operations.
func (m *GroupMock) RegisterMocks() {
	m.RegisterCreateGroupMock()
	m.RegisterGetGroupMock()
	m.RegisterUpdateGroupMock()
	m.RegisterDeleteGroupMock()
	m.RegisterListGroupsMock()
}

// RegisterErrorMocks registers error response mocks.
func (m *GroupMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterUnauthorizedErrorMock()
}

// RegisterCreateGroupMock registers a success mock for createGroup.
func (m *GroupMock) RegisterCreateGroupMock() {
	m.Register(client.EndpointApp, "createGroup", 200, "create_group_success.json")
}

// RegisterGetGroupMock registers a success mock for getGroup.
func (m *GroupMock) RegisterGetGroupMock() {
	m.Register(client.EndpointGraphQL, "getGroup", 200, "get_group_success.json")
}

// RegisterUpdateGroupMock registers a success mock for updateGroup.
func (m *GroupMock) RegisterUpdateGroupMock() {
	m.Register(client.EndpointApp, "updateGroup", 200, "update_group_success.json")
}

// RegisterDeleteGroupMock registers a success mock for deleteGroup.
func (m *GroupMock) RegisterDeleteGroupMock() {
	m.Register(client.EndpointApp, "deleteGroup", 200, "delete_group_success.json")
}

// RegisterListGroupsMock registers a success mock for listGroups.
func (m *GroupMock) RegisterListGroupsMock() {
	m.Register(client.EndpointGraphQL, "listGroups", 200, "list_groups_success.json")
}

// RegisterNotFoundErrorMock registers a not-found error mock.
func (m *GroupMock) RegisterNotFoundErrorMock() {
	m.RegisterError(client.EndpointGraphQL, "getGroup", 200, "error_not_found.json", "graphql operation failed: Group not found")
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock.
func (m *GroupMock) RegisterUnauthorizedErrorMock() {
	m.RegisterError(client.EndpointGraphQL, "getGroup", 401, "error_unauthorized.json", "Jamf Protect API error: unauthorized")
}
