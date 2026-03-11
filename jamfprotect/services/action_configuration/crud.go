package actionconfiguration

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"resty.dev/v3"
)

// Service provides operations for Jamf Protect Action Configurations.
type Service struct {
	client client.GraphQLClient
}

// NewService creates a new Action Configurations service.
func NewService(c client.GraphQLClient) *Service {
	return &Service{client: c}
}

// CreateActionConfig creates a new action configuration.
func (s *Service) CreateActionConfig(ctx context.Context, req *CreateActionConfigRequest) (*ActionConfig, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.AlertConfig == nil {
		return nil, nil, fmt.Errorf("%w: alertConfig is required", client.ErrInvalidInput)
	}

	vars := buildActionConfigVariables(req)
	var result struct {
		CreateActionConfigs *ActionConfig `json:"createActionConfigs"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(createActionConfigMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create action config: %w", err)
	}

	return result.CreateActionConfigs, resp, nil
}

// GetActionConfig retrieves an action configuration by ID.
func (s *Service) GetActionConfig(ctx context.Context, id string) (*ActionConfig, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}

	vars := map[string]any{"id": id}
	var result struct {
		GetActionConfigs *ActionConfig `json:"getActionConfigs"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(getActionConfigQuery).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get action config: %w", err)
	}

	return result.GetActionConfigs, resp, nil
}

// UpdateActionConfig updates an existing action configuration.
func (s *Service) UpdateActionConfig(ctx context.Context, id string, req *UpdateActionConfigRequest) (*ActionConfig, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.AlertConfig == nil {
		return nil, nil, fmt.Errorf("%w: alertConfig is required", client.ErrInvalidInput)
	}

	vars := buildActionConfigVariables(req)
	vars["id"] = id
	var result struct {
		UpdateActionConfigs *ActionConfig `json:"updateActionConfigs"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(updateActionConfigMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update action config: %w", err)
	}

	return result.UpdateActionConfigs, resp, nil
}

// DeleteActionConfig deletes an action configuration by ID.
func (s *Service) DeleteActionConfig(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}

	vars := map[string]any{"id": id}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(deleteActionConfigMutation).
		SetVariables(vars).
		Post(client.EndpointApp)
	if err != nil {
		return resp, fmt.Errorf("failed to delete action config: %w", err)
	}

	return resp, nil
}

// ListActionConfigs retrieves all action configurations with automatic pagination.
func (s *Service) ListActionConfigs(ctx context.Context) ([]ActionConfigListItem, *resty.Response, error) {
	allItems := make([]ActionConfigListItem, 0)
	var nextToken *string
	var lastResp *resty.Response

	for {
		vars := map[string]any{
			"direction": "ASC",
			"field":     "NAME",
		}
		if nextToken != nil {
			vars["nextToken"] = *nextToken
		}

		var result struct {
			ListActionConfigs *ListActionConfigsResponse `json:"listActionConfigs"`
		}

		resp, err := s.client.NewRequest(ctx).
			SetQuery(listActionConfigsQuery).
			SetVariables(vars).
			SetTarget(&result).
			Post(client.EndpointApp)
		lastResp = resp
		if err != nil {
			return nil, lastResp, fmt.Errorf("failed to list action configs: %w", err)
		}

		if result.ListActionConfigs != nil {
			allItems = append(allItems, result.ListActionConfigs.Items...)
			if result.ListActionConfigs.PageInfo.Next == nil {
				break
			}
			nextToken = result.ListActionConfigs.PageInfo.Next
		} else {
			break
		}
	}

	return allItems, lastResp, nil
}

// buildActionConfigVariables builds the GraphQL variables map from a request struct.
func buildActionConfigVariables(req any) map[string]any {
	var (
		name        string
		description string
		alertConfig map[string]any
		clients     []map[string]any
	)

	switch r := req.(type) {
	case *CreateActionConfigRequest:
		name = r.Name
		description = r.Description
		alertConfig = r.AlertConfig
		clients = r.Clients
	case *UpdateActionConfigRequest:
		name = r.Name
		description = r.Description
		alertConfig = r.AlertConfig
		clients = r.Clients
	}

	return map[string]any{
		"name":        name,
		"description": description,
		"alertConfig": alertConfig,
		"clients":     clients,
	}
}

// ListActionConfigNames retrieves only the names of all action configurations
func (s *Service) ListActionConfigNames(ctx context.Context) ([]string, *resty.Response, error) {
	var result struct {
		ListActionConfigNames *ListActionConfigNamesResponse `json:"listActionConfigNames"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(listActionConfigNamesQuery).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list action config names: %w", err)
	}

	names := []string{}
	if result.ListActionConfigNames != nil {
		for _, item := range result.ListActionConfigNames.Items {
			names = append(names, item.Name)
		}
	}

	return names, resp, nil
}
