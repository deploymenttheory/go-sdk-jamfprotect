package analyticset

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"resty.dev/v3"
)

// Service provides operations for Jamf Protect Analytic Sets
type Service struct {
	client client.GraphQLClient
}

// NewService creates a new Analytic Sets service
func NewService(c client.GraphQLClient) *Service {
	return &Service{client: c}
}

// CreateAnalyticSet creates a new analytic set
func (s *Service) CreateAnalyticSet(ctx context.Context, req *CreateAnalyticSetRequest) (*AnalyticSet, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.Analytics == nil {
		return nil, nil, fmt.Errorf("%w: analytics is required", client.ErrInvalidInput)
	}

	vars := analyticSetMutationVariables(req, "")
	var result struct {
		CreateAnalyticSet *AnalyticSet `json:"createAnalyticSet"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(createAnalyticSetMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create analytic set: %w", err)
	}

	return result.CreateAnalyticSet, resp, nil
}

// GetAnalyticSet retrieves an analytic set by UUID
func (s *Service) GetAnalyticSet(ctx context.Context, uuid string) (*AnalyticSet, *resty.Response, error) {
	if err := ValidateAnalyticSetUUID(uuid); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}

	vars := map[string]any{
		"uuid":             uuid,
		"RBAC_Plan":        true,
		"excludeAnalytics": false,
	}
	var result struct {
		GetAnalyticSet *AnalyticSet `json:"getAnalyticSet"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(getAnalyticSetQuery).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get analytic set: %w", err)
	}

	return result.GetAnalyticSet, resp, nil
}

// UpdateAnalyticSet updates an existing analytic set
func (s *Service) UpdateAnalyticSet(ctx context.Context, uuid string, req *UpdateAnalyticSetRequest) (*AnalyticSet, *resty.Response, error) {
	if err := ValidateAnalyticSetUUID(uuid); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.Analytics == nil {
		return nil, nil, fmt.Errorf("%w: analytics is required", client.ErrInvalidInput)
	}

	vars := analyticSetMutationVariables(req, uuid)
	var result struct {
		UpdateAnalyticSet *AnalyticSet `json:"updateAnalyticSet"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(updateAnalyticSetMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update analytic set: %w", err)
	}

	return result.UpdateAnalyticSet, resp, nil
}

// DeleteAnalyticSet deletes an analytic set by UUID
func (s *Service) DeleteAnalyticSet(ctx context.Context, uuid string) (*resty.Response, error) {
	if err := ValidateAnalyticSetUUID(uuid); err != nil {
		return nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}

	vars := map[string]any{"uuid": uuid}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(deleteAnalyticSetMutation).
		SetVariables(vars).
		Post(client.EndpointApp)
	if err != nil {
		return resp, fmt.Errorf("failed to delete analytic set: %w", err)
	}

	return resp, nil
}

// ListAnalyticSets retrieves all analytic sets with automatic pagination
func (s *Service) ListAnalyticSets(ctx context.Context) ([]AnalyticSet, *resty.Response, error) {
	allItems := make([]AnalyticSet, 0)
	var nextToken *string
	var lastResp *resty.Response

	for {
		vars := map[string]any{
			"RBAC_Plan":        true,
			"excludeAnalytics": false,
		}
		if nextToken != nil {
			vars["nextToken"] = *nextToken
		}

		var result struct {
			ListAnalyticSets *ListAnalyticSetsResponse `json:"listAnalyticSets"`
		}

		resp, err := s.client.NewRequest(ctx).
			SetQuery(listAnalyticSetsQuery).
			SetVariables(vars).
			SetTarget(&result).
			Post(client.EndpointApp)
		lastResp = resp
		if err != nil {
			return nil, lastResp, fmt.Errorf("failed to list analytic sets: %w", err)
		}

		if result.ListAnalyticSets != nil {
			allItems = append(allItems, result.ListAnalyticSets.Items...)
			if result.ListAnalyticSets.PageInfo.Next == nil {
				break
			}
			nextToken = result.ListAnalyticSets.PageInfo.Next
		} else {
			break
		}
	}

	return allItems, lastResp, nil
}
