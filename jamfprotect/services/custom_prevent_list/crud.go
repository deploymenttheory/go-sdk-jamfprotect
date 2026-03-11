package custompreventlist

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"resty.dev/v3"
)

// Service provides operations for Jamf Protect Prevent Lists
type Service struct {
	client client.GraphQLClient
}

// NewService creates a new Prevent Lists service
func NewService(c client.GraphQLClient) *Service {
	return &Service{client: c}
}

// CreatePreventList creates a new prevent list
func (s *Service) CreatePreventList(ctx context.Context, req *CreatePreventListRequest) (*PreventList, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.Type == "" {
		return nil, nil, fmt.Errorf("%w: type is required", client.ErrInvalidInput)
	}
	if err := ValidateCreatePreventListRequest(req); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}

	vars := preventListMutationVariables(req)
	var result struct {
		CreatePreventList *PreventList `json:"createPreventList"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(createPreventListMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create prevent list: %w", err)
	}

	return result.CreatePreventList, resp, nil
}

// GetPreventList retrieves a prevent list by ID
func (s *Service) GetPreventList(ctx context.Context, id string) (*PreventList, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}

	vars := map[string]any{"id": id}
	var result struct {
		GetPreventList *PreventList `json:"getPreventList"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(getPreventListQuery).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get prevent list: %w", err)
	}

	return result.GetPreventList, resp, nil
}

// UpdatePreventList updates an existing prevent list
func (s *Service) UpdatePreventList(ctx context.Context, id string, req *UpdatePreventListRequest) (*PreventList, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.Type == "" {
		return nil, nil, fmt.Errorf("%w: type is required", client.ErrInvalidInput)
	}
	if err := ValidateUpdatePreventListRequest(req); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}

	vars := preventListMutationVariables(req)
	vars["id"] = id
	var result struct {
		UpdatePreventList *PreventList `json:"updatePreventList"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(updatePreventListMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update prevent list: %w", err)
	}

	return result.UpdatePreventList, resp, nil
}

// DeletePreventList deletes a prevent list by ID
func (s *Service) DeletePreventList(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}

	vars := map[string]any{"id": id}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(deletePreventListMutation).
		SetVariables(vars).
		Post(client.EndpointGraphQL)
	if err != nil {
		return resp, fmt.Errorf("failed to delete prevent list: %w", err)
	}

	return resp, nil
}

// ListPreventLists retrieves all prevent lists with automatic pagination
func (s *Service) ListPreventLists(ctx context.Context) ([]PreventList, *resty.Response, error) {
	allItems := make([]PreventList, 0)
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
			ListPreventLists *ListPreventListsResponse `json:"listPreventLists"`
		}

		resp, err := s.client.NewRequest(ctx).
			SetQuery(listPreventListsQuery).
			SetVariables(vars).
			SetTarget(&result).
			Post(client.EndpointGraphQL)
		lastResp = resp
		if err != nil {
			return nil, lastResp, fmt.Errorf("failed to list prevent lists: %w", err)
		}

		if result.ListPreventLists != nil {
			allItems = append(allItems, result.ListPreventLists.Items...)
			if result.ListPreventLists.PageInfo.Next == nil {
				break
			}
			nextToken = result.ListPreventLists.PageInfo.Next
		} else {
			break
		}
	}

	return allItems, lastResp, nil
}

// preventListMutationVariables returns GraphQL variables for createPreventList/updatePreventList mutations.
func preventListMutationVariables(req any) map[string]any {
	var (
		name        string
		description string
		typ         string
		tags        []string
		list        []string
	)

	switch r := req.(type) {
	case *CreatePreventListRequest:
		name = r.Name
		description = r.Description
		typ = r.Type
		tags = r.Tags
		list = r.List
	case *UpdatePreventListRequest:
		name = r.Name
		description = r.Description
		typ = r.Type
		tags = r.Tags
		list = r.List
	}

	return map[string]any{
		"name":        name,
		"description": description,
		"type":        typ,
		"tags":        tags,
		"list":        list,
	}
}

// ListPreventListNames retrieves only the names of all custom prevent lists
func (s *Service) ListPreventListNames(ctx context.Context) ([]string, *resty.Response, error) {
	var result struct {
		ListPreventListNames *ListPreventListNamesResponse `json:"listPreventListNames"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(listPreventListNamesQuery).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list prevent list names: %w", err)
	}

	names := []string{}
	if result.ListPreventListNames != nil {
		for _, item := range result.ListPreventListNames.Items {
			names = append(names, item.Name)
		}
	}

	return names, resp, nil
}
