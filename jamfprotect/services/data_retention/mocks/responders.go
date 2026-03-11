package mocks

import (
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	coremocks "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/mocks"
)

// DataRetentionMock provides mock responses for the DataRetention service GraphQL operations.
type DataRetentionMock struct {
	*coremocks.GenericGraphQLMock
}

// NewDataRetentionMock creates a new DataRetentionMock instance.
func NewDataRetentionMock() *DataRetentionMock {
	return &DataRetentionMock{
		GenericGraphQLMock: coremocks.NewGenericGraphQLMock(coremocks.GenericGraphQLMockConfig{
			Name: "DataRetentionMock",
		}),
	}
}

// RegisterMocks registers all successful response mocks for data retention operations.
func (m *DataRetentionMock) RegisterMocks() {
	m.RegisterGetDataRetentionMock()
	m.RegisterUpdateDataRetentionMock()
}

// RegisterGetDataRetentionMock registers a success mock for getDataRetention.
func (m *DataRetentionMock) RegisterGetDataRetentionMock() {
	m.Register(client.EndpointGraphQL, "getDataRetention", 200, "get_data_retention_success.json")
}

// RegisterUpdateDataRetentionMock registers a success mock for updateOrganizationRetention.
func (m *DataRetentionMock) RegisterUpdateDataRetentionMock() {
	m.Register(client.EndpointApp, "updateOrganizationRetention", 200, "update_data_retention_success.json")
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock.
func (m *DataRetentionMock) RegisterUnauthorizedErrorMock() {
	m.RegisterError(client.EndpointGraphQL, "getDataRetention", 401, "error_unauthorized.json", "Jamf Protect API error: unauthorized")
}
