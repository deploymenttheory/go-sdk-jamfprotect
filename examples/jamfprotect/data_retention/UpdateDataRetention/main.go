package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
	dataretention "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/data_retention"
)

func main() {
	// Create client from environment variables
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Update data retention settings
	request := &dataretention.UpdateDataRetentionRequest{
		DatabaseLogDays:   90,
		DatabaseAlertDays: 90,
		ColdAlertDays:     365,
	}

	updated, _, err := client.DataRetention.UpdateDataRetention(ctx, request)
	if err != nil {
		log.Fatalf("Failed to update data retention settings: %v", err)
	}

	fmt.Printf("Successfully updated data retention settings:\n")
	fmt.Printf("  Database Log Days: %d\n", updated.Database.Log.NumberOfDays)
	fmt.Printf("  Database Alert Days: %d\n", updated.Database.Alert.NumberOfDays)
	fmt.Printf("  Cold Alert Days: %d\n", updated.Cold.Alert.NumberOfDays)
	fmt.Printf("  Updated: %s\n", updated.Updated)

	os.Exit(0)
}
