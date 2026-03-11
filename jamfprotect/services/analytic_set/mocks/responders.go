package mocks

import (
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	coremocks "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/mocks"
)

// AnalyticSetMock provides mock responses for the AnalyticSet service GraphQL operations.
// Most operations POST to /app; updateAnalyticSet posts to /graphql.
type AnalyticSetMock struct {
	*coremocks.GenericGraphQLMock
}

// NewAnalyticSetMock creates a new AnalyticSetMock instance.
func NewAnalyticSetMock() *AnalyticSetMock {
	return &AnalyticSetMock{
		GenericGraphQLMock: coremocks.NewGenericGraphQLMock(coremocks.GenericGraphQLMockConfig{
			Name: "AnalyticSetMock",
		}),
	}
}

// RegisterMocks registers all successful response mocks for analytic set operations.
func (m *AnalyticSetMock) RegisterMocks() {
	m.RegisterCreateAnalyticSetMock()
	m.RegisterGetAnalyticSetMock()
	m.RegisterUpdateAnalyticSetMock()
	m.RegisterDeleteAnalyticSetMock()
	m.RegisterListAnalyticSetsMock()
}

// RegisterErrorMocks registers error response mocks.
func (m *AnalyticSetMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterUnauthorizedErrorMock()
}

// RegisterCreateAnalyticSetMock registers a success mock for createAnalyticSet.
func (m *AnalyticSetMock) RegisterCreateAnalyticSetMock() {
	m.Register(client.EndpointApp, "createAnalyticSet", 200, "create_analytic_set_success.json")
}

// RegisterGetAnalyticSetMock registers a success mock for getAnalyticSet.
func (m *AnalyticSetMock) RegisterGetAnalyticSetMock() {
	m.Register(client.EndpointApp, "getAnalyticSet", 200, "get_analytic_set_success.json")
}

// RegisterUpdateAnalyticSetMock registers a success mock for updateAnalyticSet.
// Note: updateAnalyticSet uses the /graphql endpoint.
func (m *AnalyticSetMock) RegisterUpdateAnalyticSetMock() {
	m.Register(client.EndpointGraphQL, "updateAnalyticSet", 200, "update_analytic_set_success.json")
}

// RegisterDeleteAnalyticSetMock registers a success mock for deleteAnalyticSet.
func (m *AnalyticSetMock) RegisterDeleteAnalyticSetMock() {
	m.Register(client.EndpointApp, "deleteAnalyticSet", 200, "delete_analytic_set_success.json")
}

// RegisterListAnalyticSetsMock registers a success mock for listAnalyticSets.
func (m *AnalyticSetMock) RegisterListAnalyticSetsMock() {
	m.Register(client.EndpointApp, "listAnalyticSets", 200, "list_analytic_sets_success.json")
}

// RegisterNotFoundErrorMock registers a not-found error mock.
func (m *AnalyticSetMock) RegisterNotFoundErrorMock() {
	m.RegisterError(client.EndpointApp, "getAnalyticSet", 200, "error_not_found.json", "graphql operation failed: AnalyticSet not found")
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock.
func (m *AnalyticSetMock) RegisterUnauthorizedErrorMock() {
	m.RegisterError(client.EndpointApp, "getAnalyticSet", 401, "error_unauthorized.json", "Jamf Protect API error: unauthorized")
}
