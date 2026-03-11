package mocks

import (
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	coremocks "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/mocks"
)

// ExceptionSetMock provides mock responses for the ExceptionSet service GraphQL operations.
// All operations POST to the /app endpoint and are distinguished by operation name in the request body.
type ExceptionSetMock struct {
	*coremocks.GenericGraphQLMock
}

// NewExceptionSetMock creates a new ExceptionSetMock instance.
func NewExceptionSetMock() *ExceptionSetMock {
	return &ExceptionSetMock{
		GenericGraphQLMock: coremocks.NewGenericGraphQLMock(coremocks.GenericGraphQLMockConfig{
			Name: "ExceptionSetMock",
		}),
	}
}

// RegisterMocks registers all successful response mocks for exception set operations.
func (m *ExceptionSetMock) RegisterMocks() {
	m.RegisterCreateExceptionSetMock()
	m.RegisterGetExceptionSetMock()
	m.RegisterUpdateExceptionSetMock()
	m.RegisterDeleteExceptionSetMock()
	m.RegisterListExceptionSetsMock()
	m.RegisterListExceptionSetNamesMock()
}

// RegisterErrorMocks registers error response mocks.
func (m *ExceptionSetMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterUnauthorizedErrorMock()
}

// RegisterCreateExceptionSetMock registers a success mock for createExceptionSet.
func (m *ExceptionSetMock) RegisterCreateExceptionSetMock() {
	m.Register(client.EndpointApp, "createExceptionSet", 200, "create_exception_set_success.json")
}

// RegisterGetExceptionSetMock registers a success mock for getExceptionSet.
func (m *ExceptionSetMock) RegisterGetExceptionSetMock() {
	m.Register(client.EndpointApp, "getExceptionSet", 200, "get_exception_set_success.json")
}

// RegisterUpdateExceptionSetMock registers a success mock for updateExceptionSet.
func (m *ExceptionSetMock) RegisterUpdateExceptionSetMock() {
	m.Register(client.EndpointApp, "updateExceptionSet", 200, "update_exception_set_success.json")
}

// RegisterDeleteExceptionSetMock registers a success mock for deleteExceptionSet.
func (m *ExceptionSetMock) RegisterDeleteExceptionSetMock() {
	m.Register(client.EndpointApp, "deleteExceptionSet", 200, "delete_exception_set_success.json")
}

// RegisterListExceptionSetsMock registers a success mock for listExceptionSets.
func (m *ExceptionSetMock) RegisterListExceptionSetsMock() {
	m.Register(client.EndpointApp, "listExceptionSets", 200, "list_exception_sets_success.json")
}

// RegisterListExceptionSetNamesMock registers a success mock for listExceptionSetNames.
func (m *ExceptionSetMock) RegisterListExceptionSetNamesMock() {
	m.Register(client.EndpointApp, "listExceptionSetNames", 200, "list_exception_set_names_success.json")
}

// RegisterNotFoundErrorMock registers a not-found error mock.
func (m *ExceptionSetMock) RegisterNotFoundErrorMock() {
	m.RegisterError(client.EndpointApp, "getExceptionSet", 200, "error_not_found.json", "graphql operation failed: ExceptionSet not found")
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock.
func (m *ExceptionSetMock) RegisterUnauthorizedErrorMock() {
	m.RegisterError(client.EndpointApp, "getExceptionSet", 401, "error_unauthorized.json", "Jamf Protect API error: unauthorized")
}
