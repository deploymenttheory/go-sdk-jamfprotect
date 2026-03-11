package dataforwarding

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"resty.dev/v3"
)

// Service provides operations for Jamf Protect Data Forwarding settings.
type Service struct {
	client client.GraphQLClient
}

// NewService creates a new DataForwarding service.
func NewService(c client.GraphQLClient) *Service {
	return &Service{client: c}
}

// GetDataForwarding retrieves organization data forwarding settings.
func (s *Service) GetDataForwarding(ctx context.Context) (*DataForwarding, *resty.Response, error) {
	var result struct {
		GetOrganization *DataForwarding `json:"getOrganization"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(getDataForwardingQuery).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get data forwarding: %w", err)
	}

	return result.GetOrganization, resp, nil
}

// UpdateDataForwarding updates organization data forwarding settings.
func (s *Service) UpdateDataForwarding(ctx context.Context, req *UpdateDataForwardingRequest) (*DataForwarding, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}

	vars := map[string]any{
		"s3":         req.S3,
		"sentinel":   req.Sentinel,
		"sentinelV2": req.SentinelV2,
	}

	var result struct {
		UpdateOrganizationForward *DataForwarding `json:"updateOrganizationForward"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(updateDataForwardingMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update data forwarding: %w", err)
	}

	return result.UpdateOrganizationForward, resp, nil
}
