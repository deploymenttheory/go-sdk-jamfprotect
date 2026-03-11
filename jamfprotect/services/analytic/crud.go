package analytic

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"resty.dev/v3"
)

// Service provides operations for Jamf Protect Analytics
type Service struct {
	client client.GraphQLClient
}

// NewService creates a new Analytics service
func NewService(c client.GraphQLClient) *Service {
	return &Service{client: c}
}

// CreateAnalytic creates a new analytic
func (s *Service) CreateAnalytic(ctx context.Context, req *CreateAnalyticRequest) (*Analytic, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.InputType == "" {
		return nil, nil, fmt.Errorf("%w: inputType is required", client.ErrInvalidInput)
	}
	if req.Filter == "" {
		return nil, nil, fmt.Errorf("%w: filter is required", client.ErrInvalidInput)
	}
	if err := ValidateCreateAnalyticRequest(req); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}

	vars := analyticMutationVariables(req, false)
	var result struct {
		CreateAnalytic *Analytic `json:"createAnalytic"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(createAnalyticMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create analytic: %w", err)
	}

	return result.CreateAnalytic, resp, nil
}

// GetAnalytic retrieves an analytic by UUID
func (s *Service) GetAnalytic(ctx context.Context, uuid string) (*Analytic, *resty.Response, error) {
	if err := ValidateAnalyticID(uuid); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}

	vars := map[string]any{"uuid": uuid}
	var result struct {
		GetAnalytic *Analytic `json:"getAnalytic"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(getAnalyticQuery).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get analytic: %w", err)
	}

	return result.GetAnalytic, resp, nil
}

// UpdateAnalytic updates an existing analytic
func (s *Service) UpdateAnalytic(ctx context.Context, uuid string, req *UpdateAnalyticRequest) (*Analytic, *resty.Response, error) {
	if err := ValidateAnalyticID(uuid); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}
	if err := ValidateUpdateAnalyticRequest(req); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}

	vars := analyticMutationVariables(req, true)
	vars["uuid"] = uuid
	var result struct {
		UpdateAnalytic *Analytic `json:"updateAnalytic"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(updateAnalyticMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update analytic: %w", err)
	}

	return result.UpdateAnalytic, resp, nil
}

// DeleteAnalytic deletes an analytic by UUID
func (s *Service) DeleteAnalytic(ctx context.Context, uuid string) (*resty.Response, error) {
	if err := ValidateAnalyticID(uuid); err != nil {
		return nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}

	vars := map[string]any{"uuid": uuid}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(deleteAnalyticMutation).
		SetVariables(vars).
		Post(client.EndpointApp)
	if err != nil {
		return resp, fmt.Errorf("failed to delete analytic: %w", err)
	}

	return resp, nil
}

// ListAnalytics retrieves all analytics
func (s *Service) ListAnalytics(ctx context.Context) ([]Analytic, *resty.Response, error) {
	var result struct {
		ListAnalytics *ListAnalyticsResponse `json:"listAnalytics"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(listAnalyticsQuery).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list analytics: %w", err)
	}

	if result.ListAnalytics != nil {
		return result.ListAnalytics.Items, resp, nil
	}

	return []Analytic{}, resp, nil
}

// ListAnalyticsLite retrieves a lightweight summary of all analytics
func (s *Service) ListAnalyticsLite(ctx context.Context) ([]AnalyticLite, *resty.Response, error) {
	var result struct {
		ListAnalytics *ListAnalyticsLiteResponse `json:"listAnalytics"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(listAnalyticsLiteQuery).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list analytics lite: %w", err)
	}

	if result.ListAnalytics != nil {
		return result.ListAnalytics.Items, resp, nil
	}

	return []AnalyticLite{}, resp, nil
}

// ListAnalyticsNames retrieves only the names of all analytics
func (s *Service) ListAnalyticsNames(ctx context.Context) ([]string, *resty.Response, error) {
	var result struct {
		ListAnalyticsNames *struct {
			Items []struct {
				Name string `json:"name"`
			} `json:"items"`
		} `json:"listAnalyticsNames"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(listAnalyticsNamesQuery).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list analytics names: %w", err)
	}

	names := []string{}
	if result.ListAnalyticsNames != nil {
		for _, item := range result.ListAnalyticsNames.Items {
			names = append(names, item.Name)
		}
	}

	return names, resp, nil
}

// ListAnalyticsCategories retrieves all analytics categories with their counts
func (s *Service) ListAnalyticsCategories(ctx context.Context) ([]AnalyticCategory, *resty.Response, error) {
	var result struct {
		ListAnalyticsCategories []AnalyticCategory `json:"listAnalyticsCategories"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(listAnalyticsCategoriesQuery).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list analytics categories: %w", err)
	}

	if result.ListAnalyticsCategories != nil {
		return result.ListAnalyticsCategories, resp, nil
	}

	return []AnalyticCategory{}, resp, nil
}

// ListAnalyticsTags retrieves all analytics tags with their counts
func (s *Service) ListAnalyticsTags(ctx context.Context) ([]AnalyticTag, *resty.Response, error) {
	var result struct {
		ListAnalyticsTags []AnalyticTag `json:"listAnalyticsTags"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(listAnalyticsTagsQuery).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list analytics tags: %w", err)
	}

	if result.ListAnalyticsTags != nil {
		return result.ListAnalyticsTags, resp, nil
	}

	return []AnalyticTag{}, resp, nil
}

// ListAnalyticsFilterOptions retrieves both tags and categories for populating filter UIs
func (s *Service) ListAnalyticsFilterOptions(ctx context.Context) (*AnalyticsFilterOptions, *resty.Response, error) {
	var result struct {
		ListAnalyticsTags       []AnalyticTag      `json:"listAnalyticsTags"`
		ListAnalyticsCategories []AnalyticCategory `json:"listAnalyticsCategories"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(listAnalyticsFilterOptionsQuery).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list analytics filter options: %w", err)
	}

	opts := &AnalyticsFilterOptions{
		Tags:       result.ListAnalyticsTags,
		Categories: result.ListAnalyticsCategories,
	}

	if opts.Tags == nil {
		opts.Tags = []AnalyticTag{}
	}
	if opts.Categories == nil {
		opts.Categories = []AnalyticCategory{}
	}

	return opts, resp, nil
}
