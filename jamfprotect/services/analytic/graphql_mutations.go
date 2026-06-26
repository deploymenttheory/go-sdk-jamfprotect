package analytic

// analyticMutationVariables returns GraphQL variables for createAnalytic/updateAnalytic mutations.
func analyticMutationVariables(req any, isUpdate bool) map[string]any {
	var (
		name            string
		inputType       string
		description     string
		actions         []string
		analyticActions []AnalyticActionInput
		tags            []string
		categories      []string
		filter          string
		context         []AnalyticContextInput
		level           int
		severity        *string
		snapshotFiles   []string
	)

	switch r := req.(type) {
	case *CreateAnalyticRequest:
		name = r.Name
		inputType = r.InputType
		description = r.Description
		actions = r.Actions
		analyticActions = r.AnalyticActions
		tags = r.Tags
		categories = r.Categories
		filter = r.Filter
		context = r.Context
		level = r.Level
		sev := r.Severity
		severity = &sev
		snapshotFiles = r.SnapshotFiles
	case *UpdateAnalyticRequest:
		name = r.Name
		inputType = r.InputType
		description = r.Description
		actions = r.Actions
		analyticActions = r.AnalyticActions
		tags = r.Tags
		categories = r.Categories
		filter = r.Filter
		context = r.Context
		level = r.Level
		severity = r.Severity
		snapshotFiles = r.SnapshotFiles
	}

	vars := map[string]any{
		"name":          name,
		"inputType":     inputType,
		"description":   description,
		"actions":       actions,
		"tags":          tags,
		"categories":    categories,
		"filter":        filter,
		"level":         level,
		"snapshotFiles": snapshotFiles,
	}

	// Build analytic actions
	analyticActionsVars := make([]map[string]any, 0, len(analyticActions))
	for _, action := range analyticActions {
		analyticActionsVars = append(analyticActionsVars, map[string]any{
			"name":       action.Name,
			"parameters": action.Parameters,
		})
	}
	vars["analyticActions"] = analyticActionsVars

	// Build context
	contextVars := make([]map[string]any, 0, len(context))
	for _, ctx := range context {
		contextVars = append(contextVars, map[string]any{
			"name":  ctx.Name,
			"type":  ctx.Type,
			"exprs": ctx.Exprs,
		})
	}
	vars["context"] = contextVars

	// Severity is required for create, optional for update
	if !isUpdate || severity != nil {
		if severity != nil {
			vars["severity"] = *severity
		}
	}

	return vars
}

func internalAnalyticMutationVariables(req *UpdateInternalAnalyticRequest) map[string]any {
	vars := map[string]any{}
	if req == nil {
		return vars
	}

	if req.TenantActions != nil {
		tenantActionsVars := make([]map[string]any, 0, len(req.TenantActions))
		for _, action := range req.TenantActions {
			tenantActionsVars = append(tenantActionsVars, map[string]any{
				"name":       action.Name,
				"parameters": action.Parameters,
			})
		}
		vars["tenantActions"] = tenantActionsVars
	}

	if req.TenantSeverity != "" {
		vars["tenantSeverity"] = req.TenantSeverity
	}

	return vars
}
