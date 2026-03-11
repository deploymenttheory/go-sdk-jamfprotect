package exceptionset

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"resty.dev/v3"
)

// Service provides operations for Jamf Protect Exception Sets
type Service struct {
	client client.GraphQLClient
}

// NewService creates a new Exception Sets service
func NewService(c client.GraphQLClient) *Service {
	return &Service{client: c}
}

// CreateExceptionSet creates a new exception set
func (s *Service) CreateExceptionSet(ctx context.Context, req *CreateExceptionSetRequest) (*ExceptionSet, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if err := ValidateCreateExceptionSetRequest(req); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}

	vars := exceptionSetMutationVariables(req, "")
	var result struct {
		CreateExceptionSet *ExceptionSet `json:"createExceptionSet"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(createExceptionSetMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create exception set: %w", err)
	}

	return result.CreateExceptionSet, resp, nil
}

// GetExceptionSet retrieves an exception set by UUID
func (s *Service) GetExceptionSet(ctx context.Context, uuid string) (*ExceptionSet, *resty.Response, error) {
	if err := ValidateExceptionSetUUID(uuid); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}

	vars := map[string]any{
		"uuid":          uuid,
		"minimal":       false,
		"RBAC_Analytic": true,
	}
	var result struct {
		GetExceptionSet *ExceptionSet `json:"getExceptionSet"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(getExceptionSetQuery).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get exception set: %w", err)
	}

	return result.GetExceptionSet, resp, nil
}

// UpdateExceptionSet updates an existing exception set
func (s *Service) UpdateExceptionSet(ctx context.Context, uuid string, req *UpdateExceptionSetRequest) (*ExceptionSet, *resty.Response, error) {
	if err := ValidateExceptionSetUUID(uuid); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if err := ValidateUpdateExceptionSetRequest(req); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}

	vars := exceptionSetMutationVariables(req, uuid)
	var result struct {
		UpdateExceptionSet *ExceptionSet `json:"updateExceptionSet"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(updateExceptionSetMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update exception set: %w", err)
	}

	return result.UpdateExceptionSet, resp, nil
}

// DeleteExceptionSet deletes an exception set by UUID
func (s *Service) DeleteExceptionSet(ctx context.Context, uuid string) (*resty.Response, error) {
	if err := ValidateExceptionSetUUID(uuid); err != nil {
		return nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}

	vars := map[string]any{"uuid": uuid}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(deleteExceptionSetMutation).
		SetVariables(vars).
		Post(client.EndpointApp)
	if err != nil {
		return resp, fmt.Errorf("failed to delete exception set: %w", err)
	}

	return resp, nil
}

// ListExceptionSets retrieves all exception sets with automatic pagination
func (s *Service) ListExceptionSets(ctx context.Context) ([]ExceptionSetListItem, *resty.Response, error) {
	allItems := make([]ExceptionSetListItem, 0)
	var nextToken *string
	var lastResp *resty.Response

	for {
		vars := map[string]any{
			"direction": "DESC",
			"field":     "created",
		}
		if nextToken != nil {
			vars["nextToken"] = *nextToken
		}

		var result struct {
			ListExceptionSets *ListExceptionSetsResponse `json:"listExceptionSets"`
		}

		resp, err := s.client.NewRequest(ctx).
			SetQuery(listExceptionSetsQuery).
			SetVariables(vars).
			SetTarget(&result).
			Post(client.EndpointApp)
		lastResp = resp
		if err != nil {
			return nil, lastResp, fmt.Errorf("failed to list exception sets: %w", err)
		}

		if result.ListExceptionSets != nil {
			allItems = append(allItems, result.ListExceptionSets.Items...)
			if result.ListExceptionSets.PageInfo.Next == nil {
				break
			}
			nextToken = result.ListExceptionSets.PageInfo.Next
		} else {
			break
		}
	}

	return allItems, lastResp, nil
}

// exceptionSetMutationVariables returns GraphQL variables for createExceptionSet/updateExceptionSet mutations.
func exceptionSetMutationVariables(req any, uuid string) map[string]any {
	var (
		name         string
		description  string
		exceptions   []ExceptionInput
		esExceptions []EsExceptionInput
	)

	switch r := req.(type) {
	case *CreateExceptionSetRequest:
		name = r.Name
		description = r.Description
		exceptions = r.Exceptions
		esExceptions = r.EsExceptions
	case *UpdateExceptionSetRequest:
		name = r.Name
		description = r.Description
		exceptions = r.Exceptions
		esExceptions = r.EsExceptions
	}

	vars := map[string]any{
		"name":          name,
		"description":   description,
		"exceptions":    buildExceptionInputVars(exceptions),
		"esExceptions":  buildEsExceptionInputVars(esExceptions),
		"minimal":       false,
		"RBAC_Analytic": true,
	}

	if uuid != "" {
		vars["uuid"] = uuid
	}

	return vars
}

func buildExceptionInputVars(inputs []ExceptionInput) []map[string]any {
	out := make([]map[string]any, 0, len(inputs))
	for _, e := range inputs {
		m := map[string]any{
			"type":           e.Type,
			"value":          e.Value,
			"ignoreActivity": e.IgnoreActivity,
		}
		if e.AppSigningInfo != nil {
			m["appSigningInfo"] = map[string]any{
				"appId":  e.AppSigningInfo.AppId,
				"teamId": e.AppSigningInfo.TeamId,
			}
		}
		if len(e.AnalyticTypes) > 0 {
			m["analyticTypes"] = e.AnalyticTypes
		}
		if e.AnalyticUuid != "" {
			m["analyticUuid"] = e.AnalyticUuid
		}
		out = append(out, m)
	}
	return out
}

func buildEsExceptionInputVars(inputs []EsExceptionInput) []map[string]any {
	out := make([]map[string]any, 0, len(inputs))
	for _, e := range inputs {
		m := map[string]any{
			"type":           e.Type,
			"value":          e.Value,
			"ignoreActivity": e.IgnoreActivity,
		}
		if e.AppSigningInfo != nil {
			m["appSigningInfo"] = map[string]any{
				"appId":  e.AppSigningInfo.AppId,
				"teamId": e.AppSigningInfo.TeamId,
			}
		}
		if e.IgnoreListType != "" {
			m["ignoreListType"] = e.IgnoreListType
		}
		if e.IgnoreListSubType != "" {
			m["ignoreListSubType"] = e.IgnoreListSubType
		}
		if e.EventType != "" {
			m["eventType"] = e.EventType
		}
		out = append(out, m)
	}
	return out
}

// ListExceptionSetNames retrieves only the names of all exception sets
func (s *Service) ListExceptionSetNames(ctx context.Context) ([]string, *resty.Response, error) {
	var result struct {
		ListExceptionSetNames *ListExceptionSetNamesResponse `json:"listExceptionSetNames"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(listExceptionSetNamesQuery).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list exception set names: %w", err)
	}

	names := []string{}
	if result.ListExceptionSetNames != nil {
		for _, item := range result.ListExceptionSetNames.Items {
			names = append(names, item.Name)
		}
	}

	return names, resp, nil
}
