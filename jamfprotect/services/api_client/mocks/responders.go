package mocks

import (
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	coremocks "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/mocks"
)

// ApiClientMock provides mock responses for the ApiClient service GraphQL operations.
// Mutations (create/update/delete) POST to the /app endpoint; queries POST to the /graphql endpoint.
// Operations are distinguished by operation name in the request body.
type ApiClientMock struct {
	*coremocks.GenericGraphQLMock
}

// NewApiClientMock creates a new ApiClientMock instance.
func NewApiClientMock() *ApiClientMock {
	return &ApiClientMock{
		GenericGraphQLMock: coremocks.NewGenericGraphQLMock(coremocks.GenericGraphQLMockConfig{
			Name: "ApiClientMock",
		}),
	}
}

// RegisterMocks registers all successful response mocks for api client operations.
func (m *ApiClientMock) RegisterMocks() {
	m.RegisterCreateApiClientMock()
	m.RegisterGetApiClientMock()
	m.RegisterUpdateApiClientMock()
	m.RegisterDeleteApiClientMock()
	m.RegisterListApiClientsMock()
}

// RegisterErrorMocks registers error response mocks.
func (m *ApiClientMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterUnauthorizedErrorMock()
}

// RegisterCreateApiClientMock registers a success mock for createApiClient.
func (m *ApiClientMock) RegisterCreateApiClientMock() {
	m.Register(client.EndpointApp, "createApiClient", 200, "create_api_client_success.json")
}

// RegisterGetApiClientMock registers a success mock for getApiClient.
func (m *ApiClientMock) RegisterGetApiClientMock() {
	m.Register(client.EndpointGraphQL, "getApiClient", 200, "get_api_client_success.json")
}

// RegisterUpdateApiClientMock registers a success mock for updateApiClient.
func (m *ApiClientMock) RegisterUpdateApiClientMock() {
	m.Register(client.EndpointApp, "updateApiClient", 200, "update_api_client_success.json")
}

// RegisterDeleteApiClientMock registers a success mock for deleteApiClient.
func (m *ApiClientMock) RegisterDeleteApiClientMock() {
	m.Register(client.EndpointApp, "deleteApiClient", 200, "delete_api_client_success.json")
}

// RegisterListApiClientsMock registers a success mock for listApiClients.
func (m *ApiClientMock) RegisterListApiClientsMock() {
	m.Register(client.EndpointGraphQL, "listApiClients", 200, "list_api_clients_success.json")
}

// RegisterNotFoundErrorMock registers a not-found error mock.
func (m *ApiClientMock) RegisterNotFoundErrorMock() {
	m.RegisterError(client.EndpointGraphQL, "getApiClient", 200, "error_not_found.json", "graphql operation failed: ApiClient not found")
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock.
func (m *ApiClientMock) RegisterUnauthorizedErrorMock() {
	m.RegisterError(client.EndpointGraphQL, "getApiClient", 401, "error_unauthorized.json", "Jamf Protect API error: unauthorized")
}
