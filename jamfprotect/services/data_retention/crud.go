package dataretention

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"resty.dev/v3"
)

// Service provides operations for Jamf Protect Data Retention settings.
type Service struct {
	client client.GraphQLClient
}

// NewService creates a new DataRetention service.
func NewService(c client.GraphQLClient) *Service {
	return &Service{client: c}
}

// GetDataRetention retrieves organization data retention settings.
func (s *Service) GetDataRetention(ctx context.Context) (*DataRetentionSettings, *resty.Response, error) {
	var result struct {
		GetOrganization *DataRetention `json:"getOrganization"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(getDataRetentionQuery).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get data retention: %w", err)
	}

	if result.GetOrganization != nil {
		return &result.GetOrganization.Retention, resp, nil
	}

	return nil, resp, nil
}

// UpdateDataRetention updates organization data retention settings.
func (s *Service) UpdateDataRetention(ctx context.Context, req *UpdateDataRetentionRequest) (*DataRetentionSettings, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if err := ValidateUpdateDataRetentionRequest(req); err != nil {
		return nil, nil, err
	}

	vars := map[string]any{
		"databaseLogDays":   req.DatabaseLogDays,
		"databaseAlertDays": req.DatabaseAlertDays,
		"coldAlertDays":     req.ColdAlertDays,
	}

	var result struct {
		UpdateOrganizationRetention *DataRetention `json:"updateOrganizationRetention"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(updateDataRetentionMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update data retention: %w", err)
	}

	if result.UpdateOrganizationRetention != nil {
		return &result.UpdateOrganizationRetention.Retention, resp, nil
	}

	return nil, resp, nil
}
