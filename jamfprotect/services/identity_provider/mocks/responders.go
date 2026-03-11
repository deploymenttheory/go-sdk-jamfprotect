package mocks

import (
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	coremocks "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/mocks"
)

// IdentityProviderMock provides mock responses for the IdentityProvider service GraphQL operations.
// Queries POST to the /graphql endpoint.
// Operations are distinguished by operation name in the request body.
type IdentityProviderMock struct {
	*coremocks.GenericGraphQLMock
}

// NewIdentityProviderMock creates a new IdentityProviderMock instance.
func NewIdentityProviderMock() *IdentityProviderMock {
	return &IdentityProviderMock{
		GenericGraphQLMock: coremocks.NewGenericGraphQLMock(coremocks.GenericGraphQLMockConfig{
			Name: "IdentityProviderMock",
		}),
	}
}

// RegisterMocks registers all successful response mocks for identity provider operations.
func (m *IdentityProviderMock) RegisterMocks() {
	m.RegisterListConnectionsMock()
}

// RegisterListConnectionsMock registers a success mock for listConnections.
func (m *IdentityProviderMock) RegisterListConnectionsMock() {
	m.Register(client.EndpointGraphQL, "listConnections", 200, "list_connections_success.json")
}
