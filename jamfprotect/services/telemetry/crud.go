package telemetry

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"resty.dev/v3"
)

// Service provides operations for Jamf Protect Telemetry V2
type Service struct {
	client client.GraphQLClient
}

// NewService creates a new Telemetry V2 service
func NewService(c client.GraphQLClient) *Service {
	return &Service{client: c}
}

// CreateTelemetryV2 creates a new telemetry v2 configuration
func (s *Service) CreateTelemetryV2(ctx context.Context, req *CreateTelemetryV2Request) (*TelemetryV2, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.LogFiles == nil {
		return nil, nil, fmt.Errorf("%w: logFiles is required", client.ErrInvalidInput)
	}

	vars := telemetryMutationVariables(req)
	vars["RBAC_Plan"] = true
	var result struct {
		CreateTelemetryV2 *TelemetryV2 `json:"createTelemetryV2"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(createTelemetryV2Mutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create telemetry v2: %w", err)
	}

	return result.CreateTelemetryV2, resp, nil
}

// GetTelemetryV2 retrieves telemetry v2 by ID
func (s *Service) GetTelemetryV2(ctx context.Context, id string) (*TelemetryV2, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}

	vars := map[string]any{
		"id":        id,
		"RBAC_Plan": true,
	}
	var result struct {
		GetTelemetryV2 *TelemetryV2 `json:"getTelemetryV2"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(getTelemetryV2Query).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get telemetry v2: %w", err)
	}

	return result.GetTelemetryV2, resp, nil
}

// UpdateTelemetryV2 updates telemetry v2 by ID
func (s *Service) UpdateTelemetryV2(ctx context.Context, id string, req *UpdateTelemetryV2Request) (*TelemetryV2, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.LogFiles == nil {
		return nil, nil, fmt.Errorf("%w: logFiles is required", client.ErrInvalidInput)
	}

	vars := telemetryMutationVariables(req)
	vars["id"] = id
	vars["RBAC_Plan"] = true
	var result struct {
		UpdateTelemetryV2 *TelemetryV2 `json:"updateTelemetryV2"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(updateTelemetryV2Mutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update telemetry v2: %w", err)
	}

	return result.UpdateTelemetryV2, resp, nil
}

// DeleteTelemetryV2 deletes telemetry v2 by ID
func (s *Service) DeleteTelemetryV2(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}

	vars := map[string]any{"id": id}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(deleteTelemetryV2Mutation).
		SetVariables(vars).
		Post(client.EndpointApp)
	if err != nil {
		return resp, fmt.Errorf("failed to delete telemetry v2: %w", err)
	}

	return resp, nil
}

// ListTelemetriesV2 retrieves all telemetry v2 configurations with automatic pagination
func (s *Service) ListTelemetriesV2(ctx context.Context) ([]TelemetryV2, *resty.Response, error) {
	allItems := make([]TelemetryV2, 0)
	var nextToken *string
	var lastResp *resty.Response

	for {
		vars := map[string]any{
			"direction": "DESC",
			"field":     "created",
			"RBAC_Plan": true,
		}
		if nextToken != nil {
			vars["nextToken"] = *nextToken
		}

		var result struct {
			ListTelemetriesV2 *ListTelemetriesV2Response `json:"listTelemetriesV2"`
		}

		resp, err := s.client.NewRequest(ctx).
			SetQuery(listTelemetriesV2Query).
			SetVariables(vars).
			SetTarget(&result).
			Post(client.EndpointApp)
		lastResp = resp
		if err != nil {
			return nil, lastResp, fmt.Errorf("failed to list telemetries v2: %w", err)
		}

		if result.ListTelemetriesV2 != nil {
			allItems = append(allItems, result.ListTelemetriesV2.Items...)
			if result.ListTelemetriesV2.PageInfo.Next == nil {
				break
			}
			nextToken = result.ListTelemetriesV2.PageInfo.Next
		} else {
			break
		}
	}

	return allItems, lastResp, nil
}

// ListTelemetriesCombined retrieves both v1 and v2 telemetries in a single query.
// The RBAC_Plan flag controls whether plan associations are included in the response.
func (s *Service) ListTelemetriesCombined(ctx context.Context, includePlans bool) (*TelemetriesCombinedResponse, *resty.Response, error) {
	vars := map[string]any{
		"field":     "created",
		"direction": "ASC",
		"RBAC_Plan": includePlans,
	}

	var result struct {
		ListTelemetries *struct {
			Items []TelemetryV1 `json:"items"`
		} `json:"listTelemetries"`
		ListTelemetriesV2 *struct {
			Items []TelemetryV2 `json:"items"`
		} `json:"listTelemetriesV2"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(listTelemetriesCombinedQuery).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list combined telemetries: %w", err)
	}

	combined := &TelemetriesCombinedResponse{}
	if result.ListTelemetries != nil {
		combined.Telemetries = result.ListTelemetries.Items
	}
	if result.ListTelemetriesV2 != nil {
		combined.TelemetriesV2 = result.ListTelemetriesV2.Items
	}

	return combined, resp, nil
}
