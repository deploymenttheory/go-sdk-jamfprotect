package unifiedloggingfilter

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"resty.dev/v3"
)

// Service provides operations for Jamf Protect Unified Logging Filters
type Service struct {
	client client.GraphQLClient
}

// NewService creates a new Unified Logging Filters service
func NewService(c client.GraphQLClient) *Service {
	return &Service{client: c}
}

// CreateUnifiedLoggingFilter creates a new unified logging filter
func (s *Service) CreateUnifiedLoggingFilter(ctx context.Context, req *CreateUnifiedLoggingFilterRequest) (*UnifiedLoggingFilter, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.Filter == "" {
		return nil, nil, fmt.Errorf("%w: filter is required", client.ErrInvalidInput)
	}

	vars := unifiedLoggingFilterMutationVariables(req)
	var result struct {
		CreateUnifiedLoggingFilter *UnifiedLoggingFilter `json:"createUnifiedLoggingFilter"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(createUnifiedLoggingFilterMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create unified logging filter: %w", err)
	}

	return result.CreateUnifiedLoggingFilter, resp, nil
}

// GetUnifiedLoggingFilter retrieves a unified logging filter by UUID
func (s *Service) GetUnifiedLoggingFilter(ctx context.Context, uuid string) (*UnifiedLoggingFilter, *resty.Response, error) {
	if err := ValidateUnifiedLoggingFilterUUID(uuid); err != nil {
		return nil, nil, err
	}

	vars := map[string]any{"uuid": uuid}
	var result struct {
		GetUnifiedLoggingFilter *UnifiedLoggingFilter `json:"getUnifiedLoggingFilter"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(getUnifiedLoggingFilterQuery).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get unified logging filter: %w", err)
	}

	return result.GetUnifiedLoggingFilter, resp, nil
}

// UpdateUnifiedLoggingFilter updates an existing unified logging filter
func (s *Service) UpdateUnifiedLoggingFilter(ctx context.Context, uuid string, req *UpdateUnifiedLoggingFilterRequest) (*UnifiedLoggingFilter, *resty.Response, error) {
	if err := ValidateUnifiedLoggingFilterUUID(uuid); err != nil {
		return nil, nil, err
	}
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.Filter == "" {
		return nil, nil, fmt.Errorf("%w: filter is required", client.ErrInvalidInput)
	}

	vars := unifiedLoggingFilterMutationVariables(req)
	vars["uuid"] = uuid
	var result struct {
		UpdateUnifiedLoggingFilter *UnifiedLoggingFilter `json:"updateUnifiedLoggingFilter"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(updateUnifiedLoggingFilterMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update unified logging filter: %w", err)
	}

	return result.UpdateUnifiedLoggingFilter, resp, nil
}

// DeleteUnifiedLoggingFilter deletes a unified logging filter by UUID
func (s *Service) DeleteUnifiedLoggingFilter(ctx context.Context, uuid string) (*resty.Response, error) {
	if err := ValidateUnifiedLoggingFilterUUID(uuid); err != nil {
		return nil, err
	}

	vars := map[string]any{"uuid": uuid}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(deleteUnifiedLoggingFilterMutation).
		SetVariables(vars).
		Post(client.EndpointGraphQL)
	if err != nil {
		return resp, fmt.Errorf("failed to delete unified logging filter: %w", err)
	}

	return resp, nil
}

// ListUnifiedLoggingFilters retrieves all unified logging filters with automatic pagination
func (s *Service) ListUnifiedLoggingFilters(ctx context.Context) ([]UnifiedLoggingFilter, *resty.Response, error) {
	allItems := make([]UnifiedLoggingFilter, 0)
	var nextToken *string
	var lastResp *resty.Response

	for {
		vars := map[string]any{
			"direction": "ASC",
			"field":     "NAME",
			"filter":    map[string]any{},
		}
		if nextToken != nil {
			vars["nextToken"] = *nextToken
		}

		var result struct {
			ListUnifiedLoggingFilters *ListUnifiedLoggingFiltersResponse `json:"listUnifiedLoggingFilters"`
		}

		resp, err := s.client.NewRequest(ctx).
			SetQuery(listUnifiedLoggingFiltersQuery).
			SetVariables(vars).
			SetTarget(&result).
			Post(client.EndpointGraphQL)
		lastResp = resp
		if err != nil {
			return nil, lastResp, fmt.Errorf("failed to list unified logging filters: %w", err)
		}

		if result.ListUnifiedLoggingFilters != nil {
			allItems = append(allItems, result.ListUnifiedLoggingFilters.Items...)
			if result.ListUnifiedLoggingFilters.PageInfo.Next == nil {
				break
			}
			nextToken = result.ListUnifiedLoggingFilters.PageInfo.Next
		} else {
			break
		}
	}

	return allItems, lastResp, nil
}

// ListUnifiedLoggingFilterNames retrieves only the names of all unified logging filters
func (s *Service) ListUnifiedLoggingFilterNames(ctx context.Context) ([]string, *resty.Response, error) {
	var result struct {
		ListUnifiedLoggingFilterNames *ListUnifiedLoggingFilterNamesResponse `json:"listUnifiedLoggingFilterNames"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(listUnifiedLoggingFilterNamesQuery).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list unified logging filter names: %w", err)
	}

	names := []string{}
	if result.ListUnifiedLoggingFilterNames != nil {
		for _, item := range result.ListUnifiedLoggingFilterNames.Items {
			names = append(names, item.Name)
		}
	}

	return names, resp, nil
}
