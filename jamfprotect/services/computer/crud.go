package computer

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"resty.dev/v3"
)

// Service provides read-only operations for Jamf Protect Computers.
type Service struct {
	client client.GraphQLClient
}

// NewService creates a new Computer service.
func NewService(c client.GraphQLClient) *Service {
	return &Service{client: c}
}

// GetComputer retrieves a computer by UUID.
func (s *Service) GetComputer(ctx context.Context, uuid string) (*Computer, *resty.Response, error) {
	if err := ValidateComputerUUID(uuid); err != nil {
		return nil, nil, err
	}

	vars := map[string]any{
		"uuid":                         uuid,
		"isList":                       false,
		"RBAC_ThreatPreventionVersion": true,
		"RBAC_Plan":                    true,
		"RBAC_Insight":                 true,
	}

	var result struct {
		GetComputer *Computer `json:"getComputer"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(getComputerQuery).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get computer: %w", err)
	}

	return result.GetComputer, resp, nil
}

// ListComputers retrieves all computers.
func (s *Service) ListComputers(ctx context.Context) ([]Computer, *resty.Response, error) {
	vars := map[string]any{
		"pageSize":                     100,
		"direction":                    "ASC",
		"field":                        []any{"hostName"},
		"isList":                       true,
		"RBAC_ThreatPreventionVersion": true,
		"RBAC_Plan":                    true,
		"RBAC_Insight":                 true,
	}

	var result struct {
		ListComputers *ListComputersResponse `json:"listComputers"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(listComputersQuery).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list computers: %w", err)
	}

	if result.ListComputers != nil {
		return result.ListComputers.Items, resp, nil
	}

	return []Computer{}, resp, nil
}
