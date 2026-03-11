package mocks

import (
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	coremocks "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/mocks"
)

// ComputerMock provides mock responses for the Computer service GraphQL operations.
type ComputerMock struct {
	*coremocks.GenericGraphQLMock
}

// NewComputerMock creates a new ComputerMock instance.
func NewComputerMock() *ComputerMock {
	return &ComputerMock{
		GenericGraphQLMock: coremocks.NewGenericGraphQLMock(coremocks.GenericGraphQLMockConfig{
			Name: "ComputerMock",
		}),
	}
}

// RegisterMocks registers all successful response mocks for computer operations.
func (m *ComputerMock) RegisterMocks() {
	m.RegisterGetComputerMock()
	m.RegisterListComputersMock()
}

// RegisterErrorMocks registers error response mocks.
func (m *ComputerMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterUnauthorizedErrorMock()
}

// RegisterGetComputerMock registers a success mock for getComputer.
func (m *ComputerMock) RegisterGetComputerMock() {
	m.Register(client.EndpointGraphQL, "getComputer", 200, "get_computer_success.json")
}

// RegisterListComputersMock registers a success mock for listComputers.
func (m *ComputerMock) RegisterListComputersMock() {
	m.Register(client.EndpointGraphQL, "listComputers", 200, "list_computers_success.json")
}

// RegisterNotFoundErrorMock registers a not-found error mock.
func (m *ComputerMock) RegisterNotFoundErrorMock() {
	m.RegisterError(client.EndpointGraphQL, "getComputer", 200, "error_not_found.json", "graphql operation failed: Computer not found")
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock.
func (m *ComputerMock) RegisterUnauthorizedErrorMock() {
	m.RegisterError(client.EndpointGraphQL, "getComputer", 401, "error_unauthorized.json", "Jamf Protect API error: unauthorized")
}
