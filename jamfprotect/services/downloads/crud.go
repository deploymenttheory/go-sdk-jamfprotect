package downloads

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"resty.dev/v3"
)

// Service provides operations for Jamf Protect Downloads.
type Service struct {
	client client.GraphQLClient
}

// NewService creates a new Downloads service.
func NewService(c client.GraphQLClient) *Service {
	return &Service{client: c}
}

// GetOrganizationDownloads retrieves the download resources for the organization.
func (s *Service) GetOrganizationDownloads(ctx context.Context) (*OrganizationDownloads, *resty.Response, error) {
	var result struct {
		Downloads *OrganizationDownloads `json:"downloads"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(getOrganizationDownloadsQuery).
		SetTarget(&result).
		Post(client.EndpointGraphQL)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get organization downloads: %w", err)
	}

	return result.Downloads, resp, nil
}
