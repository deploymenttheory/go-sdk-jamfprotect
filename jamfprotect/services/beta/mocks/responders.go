package mocks

import (
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	coremocks "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/mocks"
)

// BetaMock provides mock responses for the Beta service GraphQL operations.
type BetaMock struct {
	*coremocks.GenericGraphQLMock
}

// NewBetaMock creates a new BetaMock instance.
func NewBetaMock() *BetaMock {
	return &BetaMock{
		GenericGraphQLMock: coremocks.NewGenericGraphQLMock(coremocks.GenericGraphQLMockConfig{
			Name: "BetaMock",
		}),
	}
}

// RegisterMocks registers all successful response mocks for beta operations.
func (m *BetaMock) RegisterMocks() {
	m.RegisterGetBetaAcceptanceStatusMock()
	m.RegisterUpdateBetaAcceptanceStatusMock()
}

// RegisterGetBetaAcceptanceStatusMock registers a success mock for getBetaAcceptanceStatus.
func (m *BetaMock) RegisterGetBetaAcceptanceStatusMock() {
	m.Register(client.EndpointApp, "getBetaAcceptanceStatus", 200, "get_beta_acceptance_status_success.json")
}

// RegisterUpdateBetaAcceptanceStatusMock registers a success mock for updateBetaAcceptanceStatus.
func (m *BetaMock) RegisterUpdateBetaAcceptanceStatusMock() {
	m.Register(client.EndpointApp, "updateBetaAcceptanceStatus", 200, "update_beta_acceptance_status_success.json")
}
