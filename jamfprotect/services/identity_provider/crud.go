package identityprovider

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"resty.dev/v3"
)

// Service provides operations for Jamf Protect Identity Provider Connections.
type Service struct {
	client client.GraphQLClient
}

// NewService creates a new IdentityProvider service.
func NewService(c client.GraphQLClient) *Service {
	return &Service{client: c}
}

// ListConnections retrieves all identity provider connections.
func (s *Service) ListConnections(ctx context.Context) ([]Connection, *resty.Response, error) {
	vars := map[string]any{
		"direction": "ASC",
		"field":     "name",
	}

	var result struct {
		ListConnections *ListConnectionsResponse `json:"listConnections"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(listConnectionsQuery).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list connections: %w", err)
	}

	if result.ListConnections != nil {
		return result.ListConnections.Items, resp, nil
	}

	return []Connection{}, resp, nil
}
