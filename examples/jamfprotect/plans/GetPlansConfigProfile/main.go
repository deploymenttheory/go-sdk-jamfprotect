package main

import (
	"context"
	"fmt"
	"log"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/plan"
)

func main() {
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	planID := "plan-id-here"

	options := &plan.PlanConfigProfileOptionsInput{
		Sign:              true,
		PPPC:              true,
		SystemExtension:   true,
		ServiceManagement: true,
		Websocket:         true,
		CA:                true,
		CSR:               true,
		Token:             true,
	}

	profile, _, err := client.Plan.GetPlansConfigProfile(ctx, planID, options)
	if err != nil {
		log.Fatalf("Failed to get plan configuration profile: %v", err)
	}

	fmt.Printf("Retrieved configuration profile for plan %s (%d bytes)\n", planID, len(profile))
}
