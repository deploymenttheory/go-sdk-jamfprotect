package apiclient

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"resty.dev/v3"
)

// Service provides operations for Jamf Protect API Clients.
type Service struct {
	client client.GraphQLClient
}

// NewService creates a new ApiClient service.
func NewService(c client.GraphQLClient) *Service {
	return &Service{client: c}
}

// CreateApiClient creates a new API client.
func (s *Service) CreateApiClient(ctx context.Context, req *CreateApiClientRequest) (*ApiClient, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}

	vars := map[string]any{
		"name":    req.Name,
		"roleIds": req.RoleIDs,
	}

	var result struct {
		CreateApiClient *ApiClient `json:"createApiClient"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(createApiClientMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create api client: %w", err)
	}

	return result.CreateApiClient, resp, nil
}

// GetApiClient retrieves an API client by clientId.
func (s *Service) GetApiClient(ctx context.Context, clientID string) (*ApiClient, *resty.Response, error) {
	if err := validateClientID(clientID); err != nil {
		return nil, nil, err
	}

	vars := map[string]any{"clientId": clientID}
	var result struct {
		GetApiClient *ApiClient `json:"getApiClient"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(getApiClientQuery).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get api client: %w", err)
	}

	return result.GetApiClient, resp, nil
}

// UpdateApiClient updates an existing API client.
func (s *Service) UpdateApiClient(ctx context.Context, clientID string, req *UpdateApiClientRequest) (*ApiClient, *resty.Response, error) {
	if err := validateClientID(clientID); err != nil {
		return nil, nil, err
	}
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}

	vars := map[string]any{
		"clientId": clientID,
		"name":     req.Name,
		"roleIds":  req.RoleIDs,
	}

	var result struct {
		UpdateApiClient *ApiClient `json:"updateApiClient"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(updateApiClientMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update api client: %w", err)
	}

	return result.UpdateApiClient, resp, nil
}

// DeleteApiClient deletes an API client by clientId.
func (s *Service) DeleteApiClient(ctx context.Context, clientID string) (*resty.Response, error) {
	if err := validateClientID(clientID); err != nil {
		return nil, err
	}

	vars := map[string]any{"clientId": clientID}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(deleteApiClientMutation).
		SetVariables(vars).
		Post(client.EndpointApp)
	if err != nil {
		return resp, fmt.Errorf("failed to delete api client: %w", err)
	}

	return resp, nil
}

// ListApiClients retrieves all API clients.
func (s *Service) ListApiClients(ctx context.Context) ([]ApiClient, *resty.Response, error) {
	vars := map[string]any{
		"direction": "DESC",
		"field":     "created",
	}

	var result struct {
		ListApiClients *ListApiClientsResponse `json:"listApiClients"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(listApiClientsQuery).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list api clients: %w", err)
	}

	if result.ListApiClients != nil {
		return result.ListApiClients.Items, resp, nil
	}

	return []ApiClient{}, resp, nil
}
