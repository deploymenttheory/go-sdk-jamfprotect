package mocks

import (
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	coremocks "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/mocks"
)

// DataForwardingMock provides mock responses for the DataForwarding service GraphQL operations.
type DataForwardingMock struct {
	*coremocks.GenericGraphQLMock
}

// NewDataForwardingMock creates a new DataForwardingMock instance.
func NewDataForwardingMock() *DataForwardingMock {
	return &DataForwardingMock{
		GenericGraphQLMock: coremocks.NewGenericGraphQLMock(coremocks.GenericGraphQLMockConfig{
			Name: "DataForwardingMock",
		}),
	}
}

// RegisterMocks registers all successful response mocks for data forwarding operations.
func (m *DataForwardingMock) RegisterMocks() {
	m.RegisterGetDataForwardingMock()
	m.RegisterUpdateDataForwardingMock()
}

// RegisterGetDataForwardingMock registers a success mock for getDataForwarding.
func (m *DataForwardingMock) RegisterGetDataForwardingMock() {
	m.Register(client.EndpointGraphQL, "getDataForwarding", 200, "get_data_forwarding_success.json")
}

// RegisterUpdateDataForwardingMock registers a success mock for updateOrganizationForward.
func (m *DataForwardingMock) RegisterUpdateDataForwardingMock() {
	m.Register(client.EndpointApp, "updateOrganizationForward", 200, "update_data_forwarding_success.json")
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock.
func (m *DataForwardingMock) RegisterUnauthorizedErrorMock() {
	m.RegisterError(client.EndpointGraphQL, "getDataForwarding", 401, "error_unauthorized.json", "Jamf Protect API error: unauthorized")
}
