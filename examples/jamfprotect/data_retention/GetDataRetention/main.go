package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
)

func main() {
	// Create client from environment variables
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Get data retention settings
	settings, _, err := client.DataRetention.GetDataRetention(ctx)
	if err != nil {
		log.Fatalf("Failed to get data retention settings: %v", err)
	}

	fmt.Printf("Data retention settings:\n")
	fmt.Printf("  Database Log Days: %d\n", settings.Database.Log.NumberOfDays)
	fmt.Printf("  Database Alert Days: %d\n", settings.Database.Alert.NumberOfDays)
	fmt.Printf("  Cold Alert Days: %d\n", settings.Cold.Alert.NumberOfDays)
	fmt.Printf("  Updated: %s\n", settings.Updated)

	os.Exit(0)
}
