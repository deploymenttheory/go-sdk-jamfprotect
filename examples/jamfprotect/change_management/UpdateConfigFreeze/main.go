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

	// Update the config freeze setting to false
	updated, _, err := client.ChangeManagement.UpdateConfigFreeze(ctx, false)
	if err != nil {
		log.Fatalf("Failed to update config freeze: %v", err)
	}

	fmt.Printf("Successfully updated config freeze:\n")
	fmt.Printf("  Config Freeze: %t\n", updated.ConfigFreeze)

	os.Exit(0)
}
