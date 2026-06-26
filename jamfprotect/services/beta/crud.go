package beta

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"resty.dev/v3"
)

// Service provides operations for Jamf Protect beta enrollment.
type Service struct {
	client client.GraphQLClient
}

// NewService creates a new Beta service.
func NewService(c client.GraphQLClient) *Service {
	return &Service{client: c}
}

// GetBetaAcceptanceStatus retrieves the tenant's beta acceptance statuses.
func (s *Service) GetBetaAcceptanceStatus(ctx context.Context) ([]BetaAcceptanceStatus, *resty.Response, error) {
	var result struct {
		GetAppInitializationData *AppInitializationData `json:"getAppInitializationData"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(getBetaAcceptanceStatusQuery).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get beta acceptance status: %w", err)
	}

	if result.GetAppInitializationData == nil || result.GetAppInitializationData.BetaAcceptanceStatus == nil {
		return []BetaAcceptanceStatus{}, resp, nil
	}

	return result.GetAppInitializationData.BetaAcceptanceStatus, resp, nil
}

// UpdateBetaAcceptanceStatus opts the tenant into a beta program.
func (s *Service) UpdateBetaAcceptanceStatus(ctx context.Context, betaName BetaName) ([]BetaAcceptanceStatus, *resty.Response, error) {
	if betaName == "" {
		return nil, nil, fmt.Errorf("%w: betaName is required", client.ErrInvalidInput)
	}

	vars := map[string]any{"betaName": betaName}
	var result struct {
		UpdateBetaAcceptanceStatus *AppInitializationData `json:"updateBetaAcceptanceStatus"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(updateBetaAcceptanceStatusMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update beta acceptance status: %w", err)
	}

	if result.UpdateBetaAcceptanceStatus == nil || result.UpdateBetaAcceptanceStatus.BetaAcceptanceStatus == nil {
		return []BetaAcceptanceStatus{}, resp, nil
	}

	return result.UpdateBetaAcceptanceStatus.BetaAcceptanceStatus, resp, nil
}
