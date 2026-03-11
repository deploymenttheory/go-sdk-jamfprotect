package group

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"resty.dev/v3"
)

// Service provides operations for Jamf Protect Groups.
type Service struct {
	client client.GraphQLClient
}

// NewService creates a new Group service.
func NewService(c client.GraphQLClient) *Service {
	return &Service{client: c}
}

// CreateGroup creates a new group.
func (s *Service) CreateGroup(ctx context.Context, req *CreateGroupRequest) (*Group, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}

	vars := map[string]any{
		"name":             req.Name,
		"connectionId":     req.ConnectionID,
		"accessGroup":      req.AccessGroup,
		"roleIds":          req.RoleIDs,
		"RBAC_Connection":  true,
		"RBAC_Role":        true,
	}

	var result struct {
		CreateGroup *Group `json:"createGroup"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(createGroupMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create group: %w", err)
	}

	return result.CreateGroup, resp, nil
}

// GetGroup retrieves a group by ID.
func (s *Service) GetGroup(ctx context.Context, id string) (*Group, *resty.Response, error) {
	if err := validateGroupID(id); err != nil {
		return nil, nil, err
	}

	vars := map[string]any{
		"id":              id,
		"RBAC_Connection": true,
		"RBAC_Role":       true,
	}

	var result struct {
		GetGroup *Group `json:"getGroup"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(getGroupQuery).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get group: %w", err)
	}

	return result.GetGroup, resp, nil
}

// UpdateGroup updates an existing group.
func (s *Service) UpdateGroup(ctx context.Context, id string, req *UpdateGroupRequest) (*Group, *resty.Response, error) {
	if err := validateGroupID(id); err != nil {
		return nil, nil, err
	}
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}

	vars := map[string]any{
		"id":              id,
		"name":            req.Name,
		"accessGroup":     req.AccessGroup,
		"roleIds":         req.RoleIDs,
		"RBAC_Connection": true,
		"RBAC_Role":       true,
	}

	var result struct {
		UpdateGroup *Group `json:"updateGroup"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(updateGroupMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update group: %w", err)
	}

	return result.UpdateGroup, resp, nil
}

// DeleteGroup deletes a group by ID.
func (s *Service) DeleteGroup(ctx context.Context, id string) (*resty.Response, error) {
	if err := validateGroupID(id); err != nil {
		return nil, err
	}

	vars := map[string]any{"id": id}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(deleteGroupMutation).
		SetVariables(vars).
		Post(client.EndpointApp)
	if err != nil {
		return resp, fmt.Errorf("failed to delete group: %w", err)
	}

	return resp, nil
}

// ListGroups retrieves all groups.
func (s *Service) ListGroups(ctx context.Context) ([]Group, *resty.Response, error) {
	vars := map[string]any{
		"direction":       "DESC",
		"field":           "created",
		"RBAC_Connection": true,
		"RBAC_Role":       true,
	}

	var result struct {
		ListGroups *ListGroupsResponse `json:"listGroups"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(listGroupsQuery).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list groups: %w", err)
	}

	if result.ListGroups != nil {
		return result.ListGroups.Items, resp, nil
	}

	return []Group{}, resp, nil
}
