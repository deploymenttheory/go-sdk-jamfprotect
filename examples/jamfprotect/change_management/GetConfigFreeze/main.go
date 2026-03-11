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

	// Get the current config freeze setting
	config, _, err := client.ChangeManagement.GetConfigFreeze(ctx)
	if err != nil {
		log.Fatalf("Failed to get config freeze: %v", err)
	}

	fmt.Printf("Config freeze settings:\n")
	fmt.Printf("  Config Freeze: %t\n", config.ConfigFreeze)

	os.Exit(0)
}
