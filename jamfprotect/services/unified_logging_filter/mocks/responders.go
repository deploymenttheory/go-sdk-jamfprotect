package mocks

import (
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	coremocks "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/mocks"
)

// UnifiedLoggingFilterMock provides mock responses for the UnifiedLoggingFilter service GraphQL operations.
// All operations POST to the /graphql endpoint and are distinguished by operation name
// in the request body.
type UnifiedLoggingFilterMock struct {
	*coremocks.GenericGraphQLMock
}

// NewUnifiedLoggingFilterMock creates a new UnifiedLoggingFilterMock instance.
func NewUnifiedLoggingFilterMock() *UnifiedLoggingFilterMock {
	return &UnifiedLoggingFilterMock{
		GenericGraphQLMock: coremocks.NewGenericGraphQLMock(coremocks.GenericGraphQLMockConfig{
			Name: "UnifiedLoggingFilterMock",
		}),
	}
}

// RegisterMocks registers all successful response mocks for unified logging filter operations.
func (m *UnifiedLoggingFilterMock) RegisterMocks() {
	m.RegisterCreateUnifiedLoggingFilterMock()
	m.RegisterGetUnifiedLoggingFilterMock()
	m.RegisterUpdateUnifiedLoggingFilterMock()
	m.RegisterDeleteUnifiedLoggingFilterMock()
	m.RegisterListUnifiedLoggingFiltersMock()
	m.RegisterListUnifiedLoggingFilterNamesMock()
}

// RegisterErrorMocks registers error response mocks.
func (m *UnifiedLoggingFilterMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterUnauthorizedErrorMock()
}

// RegisterCreateUnifiedLoggingFilterMock registers a success mock for createUnifiedLoggingFilter.
func (m *UnifiedLoggingFilterMock) RegisterCreateUnifiedLoggingFilterMock() {
	m.Register(client.EndpointGraphQL, "createUnifiedLoggingFilter", 200, "create_unified_logging_filter_success.json")
}

// RegisterGetUnifiedLoggingFilterMock registers a success mock for getUnifiedLoggingFilter.
func (m *UnifiedLoggingFilterMock) RegisterGetUnifiedLoggingFilterMock() {
	m.Register(client.EndpointGraphQL, "getUnifiedLoggingFilter", 200, "get_unified_logging_filter_success.json")
}

// RegisterUpdateUnifiedLoggingFilterMock registers a success mock for updateUnifiedLoggingFilter.
func (m *UnifiedLoggingFilterMock) RegisterUpdateUnifiedLoggingFilterMock() {
	m.Register(client.EndpointGraphQL, "updateUnifiedLoggingFilter", 200, "update_unified_logging_filter_success.json")
}

// RegisterDeleteUnifiedLoggingFilterMock registers a success mock for deleteUnifiedLoggingFilter.
func (m *UnifiedLoggingFilterMock) RegisterDeleteUnifiedLoggingFilterMock() {
	m.Register(client.EndpointGraphQL, "deleteUnifiedLoggingFilter", 200, "delete_unified_logging_filter_success.json")
}

// RegisterListUnifiedLoggingFiltersMock registers a success mock for listUnifiedLoggingFilters.
func (m *UnifiedLoggingFilterMock) RegisterListUnifiedLoggingFiltersMock() {
	m.Register(client.EndpointGraphQL, "listUnifiedLoggingFilters", 200, "list_unified_logging_filters_success.json")
}

// RegisterListUnifiedLoggingFilterNamesMock registers a success mock for listUnifiedLoggingFilterNames.
func (m *UnifiedLoggingFilterMock) RegisterListUnifiedLoggingFilterNamesMock() {
	m.Register(client.EndpointGraphQL, "listUnifiedLoggingFilterNames", 200, "list_unified_logging_filter_names_success.json")
}

// RegisterNotFoundErrorMock registers a not-found error mock.
func (m *UnifiedLoggingFilterMock) RegisterNotFoundErrorMock() {
	m.RegisterError(client.EndpointGraphQL, "getUnifiedLoggingFilter", 200, "error_not_found.json", "graphql operation failed: UnifiedLoggingFilter not found")
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock.
func (m *UnifiedLoggingFilterMock) RegisterUnauthorizedErrorMock() {
	m.RegisterError(client.EndpointGraphQL, "getUnifiedLoggingFilter", 401, "error_unauthorized.json", "Jamf Protect API error: unauthorized")
}
