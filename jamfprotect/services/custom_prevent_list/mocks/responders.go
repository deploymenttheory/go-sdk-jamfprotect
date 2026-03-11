package mocks

import (
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	coremocks "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/mocks"
)

// PreventListMock provides mock responses for the CustomPreventList service GraphQL operations.
// All operations POST to the /graphql endpoint and are distinguished by operation name
// in the request body.
type PreventListMock struct {
	*coremocks.GenericGraphQLMock
}

// NewPreventListMock creates a new PreventListMock instance.
func NewPreventListMock() *PreventListMock {
	return &PreventListMock{
		GenericGraphQLMock: coremocks.NewGenericGraphQLMock(coremocks.GenericGraphQLMockConfig{
			Name: "PreventListMock",
		}),
	}
}

// RegisterMocks registers all successful response mocks for prevent list operations.
func (m *PreventListMock) RegisterMocks() {
	m.RegisterCreatePreventListMock()
	m.RegisterGetPreventListMock()
	m.RegisterUpdatePreventListMock()
	m.RegisterDeletePreventListMock()
	m.RegisterListPreventListsMock()
	m.RegisterListPreventListNamesMock()
}

// RegisterErrorMocks registers error response mocks.
func (m *PreventListMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterUnauthorizedErrorMock()
}

// RegisterCreatePreventListMock registers a success mock for createPreventList.
func (m *PreventListMock) RegisterCreatePreventListMock() {
	m.Register(client.EndpointGraphQL, "createPreventList", 200, "create_prevent_list_success.json")
}

// RegisterGetPreventListMock registers a success mock for getPreventList.
func (m *PreventListMock) RegisterGetPreventListMock() {
	m.Register(client.EndpointGraphQL, "getPreventList", 200, "get_prevent_list_success.json")
}

// RegisterUpdatePreventListMock registers a success mock for updatePreventList.
func (m *PreventListMock) RegisterUpdatePreventListMock() {
	m.Register(client.EndpointGraphQL, "updatePreventList", 200, "update_prevent_list_success.json")
}

// RegisterDeletePreventListMock registers a success mock for deletePreventList.
func (m *PreventListMock) RegisterDeletePreventListMock() {
	m.Register(client.EndpointGraphQL, "deletePreventList", 200, "delete_prevent_list_success.json")
}

// RegisterListPreventListsMock registers a success mock for listPreventLists.
func (m *PreventListMock) RegisterListPreventListsMock() {
	m.Register(client.EndpointGraphQL, "listPreventLists", 200, "list_prevent_lists_success.json")
}

// RegisterListPreventListNamesMock registers a success mock for listPreventListNames.
func (m *PreventListMock) RegisterListPreventListNamesMock() {
	m.Register(client.EndpointGraphQL, "listPreventListNames", 200, "list_prevent_list_names_success.json")
}

// RegisterNotFoundErrorMock registers a not-found error mock.
func (m *PreventListMock) RegisterNotFoundErrorMock() {
	m.RegisterError(client.EndpointGraphQL, "getPreventList", 200, "error_not_found.json", "graphql operation failed: PreventList not found")
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock.
func (m *PreventListMock) RegisterUnauthorizedErrorMock() {
	m.RegisterError(client.EndpointGraphQL, "getPreventList", 401, "error_unauthorized.json", "Jamf Protect API error: unauthorized")
}
