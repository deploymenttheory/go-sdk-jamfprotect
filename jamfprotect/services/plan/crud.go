package plan

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"resty.dev/v3"
)

// Service provides operations for Jamf Protect Plans
type Service struct {
	client client.GraphQLClient
}

// NewService creates a new Plans service
func NewService(c client.GraphQLClient) *Service {
	return &Service{client: c}
}

// CreatePlan creates a new plan
func (s *Service) CreatePlan(ctx context.Context, req *CreatePlanRequest) (*Plan, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.ActionConfigs == "" {
		return nil, nil, fmt.Errorf("%w: actionConfigs is required", client.ErrInvalidInput)
	}
	if err := ValidateCreatePlanRequest(req); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}

	vars := planMutationVariables(req)
	var result struct {
		CreatePlan *Plan `json:"createPlan"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(createPlanMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create plan: %w", err)
	}

	return result.CreatePlan, resp, nil
}

// GetPlan retrieves a plan by ID
func (s *Service) GetPlan(ctx context.Context, id string) (*Plan, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}

	vars := map[string]any{"id": id}
	var result struct {
		GetPlan *Plan `json:"getPlan"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(getPlanQuery).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get plan: %w", err)
	}

	return result.GetPlan, resp, nil
}

// UpdatePlan updates an existing plan
func (s *Service) UpdatePlan(ctx context.Context, id string, req *UpdatePlanRequest) (*Plan, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}
	if err := ValidateUpdatePlanRequest(req); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}

	vars := planMutationVariables(req)
	vars["id"] = id
	var result struct {
		UpdatePlan *Plan `json:"updatePlan"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(updatePlanMutation).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update plan: %w", err)
	}

	return result.UpdatePlan, resp, nil
}

// DeletePlan deletes a plan by ID
func (s *Service) DeletePlan(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}

	vars := map[string]any{"id": id}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(deletePlanMutation).
		SetVariables(vars).
		Post(client.EndpointApp)
	if err != nil {
		return resp, fmt.Errorf("failed to delete plan: %w", err)
	}

	return resp, nil
}

// ListPlans retrieves all plans with automatic pagination
func (s *Service) ListPlans(ctx context.Context) ([]Plan, *resty.Response, error) {
	allItems := make([]Plan, 0)
	var nextToken *string
	var lastResp *resty.Response

	for {
		vars := map[string]any{
			"direction": "ASC",
			"field":     "CREATED",
		}
		if nextToken != nil {
			vars["nextToken"] = *nextToken
		}

		var result struct {
			ListPlans *ListPlansResponse `json:"listPlans"`
		}

		resp, err := s.client.NewRequest(ctx).
			SetQuery(listPlansQuery).
			SetVariables(vars).
			SetTarget(&result).
			Post(client.EndpointApp)
		lastResp = resp
		if err != nil {
			return nil, lastResp, fmt.Errorf("failed to list plans: %w", err)
		}

		if result.ListPlans != nil {
			allItems = append(allItems, result.ListPlans.Items...)
			if result.ListPlans.PageInfo.Next == nil {
				break
			}
			nextToken = result.ListPlans.PageInfo.Next
		} else {
			break
		}
	}

	return allItems, lastResp, nil
}

// ListPlanNames retrieves only the names of all plans with automatic pagination
func (s *Service) ListPlanNames(ctx context.Context) ([]string, *resty.Response, error) {
	allNames := make([]string, 0)
	var nextToken *string
	var lastResp *resty.Response

	for {
		vars := map[string]any{}
		if nextToken != nil {
			vars["nextToken"] = *nextToken
		}

		var result struct {
			ListPlanNames *ListPlanNamesResponse `json:"listPlanNames"`
		}

		resp, err := s.client.NewRequest(ctx).
			SetQuery(listPlanNamesQuery).
			SetVariables(vars).
			SetTarget(&result).
			Post(client.EndpointApp)
		lastResp = resp
		if err != nil {
			return nil, lastResp, fmt.Errorf("failed to list plan names: %w", err)
		}

		if result.ListPlanNames != nil {
			for _, item := range result.ListPlanNames.Items {
				allNames = append(allNames, item.Name)
			}
			if result.ListPlanNames.PageInfo.Next == nil {
				break
			}
			nextToken = result.ListPlanNames.PageInfo.Next
		} else {
			break
		}
	}

	return allNames, lastResp, nil
}

// GetPlanConfigurationAndSetOptions retrieves all resources available for plan configuration,
// gated by RBAC flags. Returns action configs, telemetries (v1 and v2), USB control sets,
// exception sets, and both managed and unmanaged analytic sets.
func (s *Service) GetPlanConfigurationAndSetOptions(ctx context.Context, req *GetPlanConfigurationAndSetOptionsRequest) (*PlanConfigurationAndSetOptions, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request is required", client.ErrInvalidInput)
	}

	vars := map[string]any{
		"RBAC_ActionConfigs": req.RBACActionConfigs,
		"RBAC_Telemetry":     req.RBACTelemetry,
		"RBAC_USBControlSet": req.RBACUSBControlSet,
		"RBAC_ExceptionSet":  req.RBACExceptionSet,
		"RBAC_AnalyticSet":   req.RBACAnalyticSet,
	}

	var result struct {
		ActionConfigs *struct {
			Items []PlanConfigRefItem `json:"items"`
		} `json:"actionConfigs"`
		Telemetries *struct {
			Items []PlanConfigRefItem `json:"items"`
		} `json:"telemetries"`
		TelemetriesV2 *struct {
			Items []PlanConfigRefItem `json:"items"`
		} `json:"telemetriesV2"`
		USBControlSets *struct {
			Items []PlanConfigRefItem `json:"items"`
		} `json:"usbControlSets"`
		ExceptionSets *struct {
			Items []PlanConfigExceptionSetItem `json:"items"`
		} `json:"exceptionSets"`
		AnalyticSets *struct {
			Items []PlanConfigAnalyticSetItem `json:"items"`
		} `json:"analyticSets"`
		ManagedAnalyticSets *struct {
			Items []PlanConfigAnalyticSetItem `json:"items"`
		} `json:"managedAnalyticSets"`
	}

	resp, err := s.client.NewRequest(ctx).
		SetQuery(getPlanConfigurationAndSetOptionsQuery).
		SetVariables(vars).
		SetTarget(&result).
		Post(client.EndpointApp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get plan configuration and set options: %w", err)
	}

	opts := &PlanConfigurationAndSetOptions{}

	if result.ActionConfigs != nil {
		opts.ActionConfigs = result.ActionConfigs.Items
	}
	if result.Telemetries != nil {
		opts.Telemetries = result.Telemetries.Items
	}
	if result.TelemetriesV2 != nil {
		opts.TelemetriesV2 = result.TelemetriesV2.Items
	}
	if result.USBControlSets != nil {
		opts.USBControlSets = result.USBControlSets.Items
	}
	if result.ExceptionSets != nil {
		opts.ExceptionSets = result.ExceptionSets.Items
	}
	if result.AnalyticSets != nil {
		opts.AnalyticSets = result.AnalyticSets.Items
	}
	if result.ManagedAnalyticSets != nil {
		opts.ManagedAnalyticSets = result.ManagedAnalyticSets.Items
	}

	return opts, resp, nil
}
