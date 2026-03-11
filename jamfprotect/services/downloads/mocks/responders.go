package mocks

import (
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	coremocks "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/mocks"
)

// DownloadsMock provides mock responses for the Downloads service GraphQL operations.
// Queries POST to the /graphql endpoint.
type DownloadsMock struct {
	*coremocks.GenericGraphQLMock
}

// NewDownloadsMock creates a new DownloadsMock instance.
func NewDownloadsMock() *DownloadsMock {
	return &DownloadsMock{
		GenericGraphQLMock: coremocks.NewGenericGraphQLMock(coremocks.GenericGraphQLMockConfig{
			Name: "DownloadsMock",
		}),
	}
}

// RegisterMocks registers all successful response mocks for downloads operations.
func (m *DownloadsMock) RegisterMocks() {
	m.RegisterGetOrganizationDownloadsMock()
}

// RegisterGetOrganizationDownloadsMock registers a success mock for getOrganizationDownloads.
func (m *DownloadsMock) RegisterGetOrganizationDownloadsMock() {
	m.Register(client.EndpointGraphQL, "getOrganizationDownloads", 200, "get_organization_downloads_success.json")
}
