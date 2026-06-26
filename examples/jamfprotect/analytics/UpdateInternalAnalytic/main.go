package main

import (
	"context"
	"fmt"
	"log"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/analytic"
)

func main() {
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	uuid := "analytic-uuid-here"

	req := &analytic.UpdateInternalAnalyticRequest{
		TenantSeverity: "High",
		TenantActions: []analytic.AnalyticActionInput{
			{Name: "notify", Parameters: []string{}},
		},
	}

	result, _, err := client.Analytic.UpdateInternalAnalytic(ctx, uuid, req)
	if err != nil {
		log.Fatalf("Failed to update internal analytic: %v", err)
	}

	fmt.Printf("Updated managed analytic %s (tenant severity: %s)\n", result.UUID, result.TenantSeverity)
}
