package plan

// planMutationVariables returns GraphQL variables for createPlan/updatePlan mutations.
func planMutationVariables(req any) map[string]any {
	var (
		name                 string
		description          string
		logLevel             *string
		actionConfigs        string
		exceptionSets        []string
		telemetry            *string
		telemetryV2          *string
		telemetryV2Null      bool
		analyticSets         []AnalyticSetInput
		usbControlSet        *string
		commsConfig          CommsConfigInput
		infoSync             InfoSyncInput
		autoUpdate           bool
		signaturesFeedConfig SignaturesFeedConfigInput
	)

	switch r := req.(type) {
	case *CreatePlanRequest:
		name = r.Name
		description = r.Description
		logLevel = r.LogLevel
		actionConfigs = r.ActionConfigs
		exceptionSets = r.ExceptionSets
		telemetry = r.Telemetry
		telemetryV2 = r.TelemetryV2
		telemetryV2Null = r.TelemetryV2Null
		analyticSets = r.AnalyticSets
		usbControlSet = r.USBControlSet
		commsConfig = r.CommsConfig
		infoSync = r.InfoSync
		autoUpdate = r.AutoUpdate
		signaturesFeedConfig = r.SignaturesFeedConfig
	case *UpdatePlanRequest:
		name = r.Name
		description = r.Description
		logLevel = r.LogLevel
		actionConfigs = r.ActionConfigs
		exceptionSets = r.ExceptionSets
		telemetry = r.Telemetry
		telemetryV2 = r.TelemetryV2
		telemetryV2Null = r.TelemetryV2Null
		analyticSets = r.AnalyticSets
		usbControlSet = r.USBControlSet
		commsConfig = r.CommsConfig
		infoSync = r.InfoSync
		autoUpdate = r.AutoUpdate
		signaturesFeedConfig = r.SignaturesFeedConfig
	}

	vars := map[string]any{
		"name":          name,
		"description":   description,
		"actionConfigs": actionConfigs,
		"autoUpdate":    autoUpdate,
		"commsConfig": map[string]any{
			"fqdn":     commsConfig.FQDN,
			"protocol": commsConfig.Protocol,
		},
		"infoSync": map[string]any{
			"attrs":                infoSync.Attrs,
			"insightsSyncInterval": infoSync.InsightsSyncInterval,
		},
		"signaturesFeedConfig": map[string]any{
			"mode": signaturesFeedConfig.Mode,
		},
	}

	if logLevel != nil {
		vars["logLevel"] = *logLevel
	}

	if exceptionSets != nil {
		vars["exceptionSets"] = exceptionSets
	}

	if telemetry != nil {
		vars["telemetry"] = *telemetry
	}

	if telemetryV2Null {
		vars["telemetryV2"] = nil
	} else if telemetryV2 != nil {
		vars["telemetryV2"] = *telemetryV2
	}

	if analyticSets != nil {
		analyticSetsVars := make([]map[string]any, 0, len(analyticSets))
		for _, set := range analyticSets {
			analyticSetsVars = append(analyticSetsVars, map[string]any{
				"type": set.Type,
				"uuid": set.UUID,
			})
		}
		vars["analyticSets"] = analyticSetsVars
	}

	if usbControlSet != nil {
		vars["usbControlSet"] = *usbControlSet
	}

	return vars
}

func planConfigProfileOptionsVariables(input *PlanConfigProfileOptionsInput) map[string]any {
	if input == nil {
		return nil
	}

	vars := map[string]any{
		"pppc":    input.PPPC,
		"token":   input.Token,
		"tokenOptions": map[string]any{
			"xpc":               input.TokenOptions.XPC,
			"keychain_client_id": input.TokenOptions.KeychainClientID,
		},
		"ca":                input.CA,
		"csr":               input.CSR,
		"websocket":         input.Websocket,
		"sign":              input.Sign,
		"systemExtension":   input.SystemExtension,
		"serviceManagement": input.ServiceManagement,
	}

	if input.ConfigVersion != nil {
		vars["configVersion"] = *input.ConfigVersion
	}

	return vars
}
