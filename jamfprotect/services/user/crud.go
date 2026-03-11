package user

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"resty.dev/v3"
)

// Service provides operations for Jamf Protect Users.
type Service struct {
	client client.GraphQLClient
}

// NewService creates a new User service.
func NewService(c client.GraphQLClient) *Service {
	return &Service{client: c}
}

// CreateUser creates a new user.
func (s *Service) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Email == "" {
		return nil, nil, fmt.Errorf("%w: email is required", client.ErrInvalidInput)
	}
	if req.EmailAlertMinSeverity == "" {
		return nil, nil, fmt.Errorf("%w: emailAlertMinSeverity is required", client.ErrInvalidInput)
	}

	vars := map[string]any{
		"email":                 req.Email,
		"roleIds":               req.RoleIDs,
		"groupIds":              req.GroupIDs,
		"connectionId":          req.ConnectionID,
		"receiveEmailAlert":     req.ReceiveEmailAlert,
		"emailAlertMinSeverity": req.EmailAlertMinSeverity,
		"RBAC_Connection":       true,
		"RBAC_Role":             true,
		"RBAC_Group":            true,
		"hasLimitedAppAccess":   false,
	}

	var result struct {
		CreateUser *User `json:"createUser"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(createUserMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create user: %w", err)
	}

	return result.CreateUser, resp, nil
}

// GetUser retrieves a user by ID.
func (s *Service) GetUser(ctx context.Context, id string) (*User, *resty.Response, error) {
	if err := validateUserID(id); err != nil {
		return nil, nil, err
	}

	vars := map[string]any{
		"id":                  id,
		"hasLimitedAppAccess": false,
		"RBAC_Connection":     true,
		"RBAC_Role":           true,
		"RBAC_Group":          true,
	}

	var result struct {
		GetUser *User `json:"getUser"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(getUserQuery).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get user: %w", err)
	}

	return result.GetUser, resp, nil
}

// UpdateUser updates an existing user.
func (s *Service) UpdateUser(ctx context.Context, id string, req *UpdateUserRequest) (*User, *resty.Response, error) {
	if err := validateUserID(id); err != nil {
		return nil, nil, err
	}
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}

	vars := map[string]any{
		"id":                    id,
		"roleIds":               req.RoleIDs,
		"groupIds":              req.GroupIDs,
		"receiveEmailAlert":     req.ReceiveEmailAlert,
		"emailAlertMinSeverity": req.EmailAlertMinSeverity,
		"RBAC_Connection":       true,
		"RBAC_Role":             true,
		"RBAC_Group":            true,
		"hasLimitedAppAccess":   false,
	}

	var result struct {
		UpdateUser *User `json:"updateUser"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(updateUserMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update user: %w", err)
	}

	return result.UpdateUser, resp, nil
}

// DeleteUser deletes a user by ID.
func (s *Service) DeleteUser(ctx context.Context, id string) (*resty.Response, error) {
	if err := validateUserID(id); err != nil {
		return nil, err
	}

	vars := map[string]any{"id": id}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(deleteUserMutation).
		SetVariables(vars).
		Post(client.EndpointApp)
	if err != nil {
		return resp, fmt.Errorf("failed to delete user: %w", err)
	}

	return resp, nil
}

// ListUsers retrieves all users.
func (s *Service) ListUsers(ctx context.Context) ([]User, *resty.Response, error) {
	vars := map[string]any{
		"direction":           "DESC",
		"field":               "created",
		"hasLimitedAppAccess": false,
		"RBAC_Connection":     true,
		"RBAC_Role":           true,
		"RBAC_Group":          true,
	}

	var result struct {
		ListUsers *ListUsersResponse `json:"listUsers"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(listUsersQuery).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list users: %w", err)
	}

	if result.ListUsers != nil {
		return result.ListUsers.Items, resp, nil
	}

	return []User{}, resp, nil
}
