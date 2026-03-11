package role

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"resty.dev/v3"
)

// Service provides operations for Jamf Protect Roles.
type Service struct {
	client client.GraphQLClient
}

// NewService creates a new Role service.
func NewService(c client.GraphQLClient) *Service {
	return &Service{client: c}
}

// CreateRole creates a new role.
func (s *Service) CreateRole(ctx context.Context, req *CreateRoleRequest) (*Role, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.ReadResources == nil {
		return nil, nil, fmt.Errorf("%w: readResources is required", client.ErrInvalidInput)
	}
	if req.WriteResources == nil {
		return nil, nil, fmt.Errorf("%w: writeResources is required", client.ErrInvalidInput)
	}

	vars := map[string]any{
		"name":           req.Name,
		"readResources":  req.ReadResources,
		"writeResources": req.WriteResources,
	}

	var result struct {
		CreateRole *Role `json:"createRole"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(createRoleMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create role: %w", err)
	}

	return result.CreateRole, resp, nil
}

// GetRole retrieves a role by ID.
func (s *Service) GetRole(ctx context.Context, id string) (*Role, *resty.Response, error) {
	if err := validateRoleID(id); err != nil {
		return nil, nil, err
	}

	vars := map[string]any{"id": id}

	var result struct {
		GetRole *Role `json:"getRole"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(getRoleQuery).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get role: %w", err)
	}

	return result.GetRole, resp, nil
}

// UpdateRole updates an existing role.
func (s *Service) UpdateRole(ctx context.Context, id string, req *UpdateRoleRequest) (*Role, *resty.Response, error) {
	if err := validateRoleID(id); err != nil {
		return nil, nil, err
	}
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}

	vars := map[string]any{
		"id":             id,
		"name":           req.Name,
		"readResources":  req.ReadResources,
		"writeResources": req.WriteResources,
	}

	var result struct {
		UpdateRole *Role `json:"updateRole"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(updateRoleMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update role: %w", err)
	}

	return result.UpdateRole, resp, nil
}

// DeleteRole deletes a role by ID.
func (s *Service) DeleteRole(ctx context.Context, id string) (*resty.Response, error) {
	if err := validateRoleID(id); err != nil {
		return nil, err
	}

	vars := map[string]any{"id": id}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(deleteRoleMutation).
		SetVariables(vars).
		Post(client.EndpointApp)
	if err != nil {
		return resp, fmt.Errorf("failed to delete role: %w", err)
	}

	return resp, nil
}

// ListRoles retrieves all roles.
func (s *Service) ListRoles(ctx context.Context) ([]Role, *resty.Response, error) {
	vars := map[string]any{
		"direction": "DESC",
		"field":     "created",
	}

	var result struct {
		ListRoles *ListRolesResponse `json:"listRoles"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(listRolesQuery).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list roles: %w", err)
	}

	if result.ListRoles != nil {
		return result.ListRoles.Items, resp, nil
	}

	return []Role{}, resp, nil
}
