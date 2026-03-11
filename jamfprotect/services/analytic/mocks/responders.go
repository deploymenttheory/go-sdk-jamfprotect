package mocks

import (
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	coremocks "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/mocks"
)

// AnalyticMock provides mock responses for the Analytic service GraphQL operations.
// Mutations (create/update/delete) POST to the /app endpoint; queries POST to the /graphql endpoint.
// Operations are distinguished by operation name in the request body.
type AnalyticMock struct {
	*coremocks.GenericGraphQLMock
}

// NewAnalyticMock creates a new AnalyticMock instance.
func NewAnalyticMock() *AnalyticMock {
	return &AnalyticMock{
		GenericGraphQLMock: coremocks.NewGenericGraphQLMock(coremocks.GenericGraphQLMockConfig{
			Name: "AnalyticMock",
		}),
	}
}

// RegisterMocks registers all successful response mocks for analytic operations.
func (m *AnalyticMock) RegisterMocks() {
	m.RegisterCreateAnalyticMock()
	m.RegisterGetAnalyticMock()
	m.RegisterUpdateAnalyticMock()
	m.RegisterDeleteAnalyticMock()
	m.RegisterListAnalyticsMock()
	m.RegisterListAnalyticsLiteMock()
	m.RegisterListAnalyticsNamesMock()
	m.RegisterListAnalyticsCategoriesMock()
	m.RegisterListAnalyticsTagsMock()
	m.RegisterListAnalyticsFilterOptionsMock()
}

// RegisterErrorMocks registers error response mocks.
func (m *AnalyticMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterUnauthorizedErrorMock()
}

// RegisterCreateAnalyticMock registers a success mock for createAnalytic.
func (m *AnalyticMock) RegisterCreateAnalyticMock() {
	m.Register(client.EndpointApp, "createAnalytic", 200, "create_analytic_success.json")
}

// RegisterGetAnalyticMock registers a success mock for getAnalytic.
func (m *AnalyticMock) RegisterGetAnalyticMock() {
	m.Register(client.EndpointGraphQL, "getAnalytic", 200, "get_analytic_success.json")
}

// RegisterUpdateAnalyticMock registers a success mock for updateAnalytic.
func (m *AnalyticMock) RegisterUpdateAnalyticMock() {
	m.Register(client.EndpointApp, "updateAnalytic", 200, "update_analytic_success.json")
}

// RegisterDeleteAnalyticMock registers a success mock for deleteAnalytic.
func (m *AnalyticMock) RegisterDeleteAnalyticMock() {
	m.Register(client.EndpointApp, "deleteAnalytic", 200, "delete_analytic_success.json")
}

// RegisterListAnalyticsMock registers a success mock for listAnalytics.
func (m *AnalyticMock) RegisterListAnalyticsMock() {
	m.Register(client.EndpointGraphQL, "listAnalytics", 200, "list_analytics_success.json")
}

// RegisterListAnalyticsLiteMock registers a success mock for listAnalyticsLite.
func (m *AnalyticMock) RegisterListAnalyticsLiteMock() {
	m.Register(client.EndpointGraphQL, "listAnalyticsLite", 200, "list_analytics_lite_success.json")
}

// RegisterListAnalyticsNamesMock registers a success mock for listAnalyticsNames.
func (m *AnalyticMock) RegisterListAnalyticsNamesMock() {
	m.Register(client.EndpointGraphQL, "listAnalyticsNames", 200, "list_analytics_names_success.json")
}

// RegisterListAnalyticsCategoriesMock registers a success mock for listAnalyticsCategories.
func (m *AnalyticMock) RegisterListAnalyticsCategoriesMock() {
	m.Register(client.EndpointGraphQL, "listAnalyticsCategories", 200, "list_analytics_categories_success.json")
}

// RegisterListAnalyticsTagsMock registers a success mock for listAnalyticsTags.
func (m *AnalyticMock) RegisterListAnalyticsTagsMock() {
	m.Register(client.EndpointGraphQL, "listAnalyticsTags", 200, "list_analytics_tags_success.json")
}

// RegisterListAnalyticsFilterOptionsMock registers a success mock for listAnalyticsFilterOptions.
func (m *AnalyticMock) RegisterListAnalyticsFilterOptionsMock() {
	m.Register(client.EndpointGraphQL, "listAnalyticsFilterOptions", 200, "list_analytics_filter_options_success.json")
}

// RegisterNotFoundErrorMock registers a not-found error mock.
func (m *AnalyticMock) RegisterNotFoundErrorMock() {
	m.RegisterError(client.EndpointGraphQL, "getAnalytic", 200, "error_not_found.json", "graphql operation failed: Analytic not found")
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock.
func (m *AnalyticMock) RegisterUnauthorizedErrorMock() {
	m.RegisterError(client.EndpointGraphQL, "getAnalytic", 401, "error_unauthorized.json", "Jamf Protect API error: unauthorized")
}
