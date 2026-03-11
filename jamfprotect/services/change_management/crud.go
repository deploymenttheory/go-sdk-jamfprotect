package changemanagement

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"resty.dev/v3"
)

// Service provides operations for Jamf Protect Change Management (config freeze).
type Service struct {
	client client.GraphQLClient
}

// NewService creates a new ChangeManagement service.
func NewService(c client.GraphQLClient) *Service {
	return &Service{client: c}
}

// GetConfigFreeze retrieves the current config freeze setting.
func (s *Service) GetConfigFreeze(ctx context.Context) (*ChangeManagementConfig, *resty.Response, error) {
	var result struct {
		GetAppInitializationData *ChangeManagementConfig `json:"getAppInitializationData"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(getConfigFreezeQuery).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get config freeze: %w", err)
	}

	return result.GetAppInitializationData, resp, nil
}

// UpdateConfigFreeze updates the config freeze setting.
func (s *Service) UpdateConfigFreeze(ctx context.Context, freeze bool) (*ChangeManagementConfig, *resty.Response, error) {
	vars := map[string]any{"configFreeze": freeze}

	var result struct {
		UpdateOrganizationConfigFreeze *ChangeManagementConfig `json:"updateOrganizationConfigFreeze"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(updateConfigFreezeMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update config freeze: %w", err)
	}

	return result.UpdateOrganizationConfigFreeze, resp, nil
}
