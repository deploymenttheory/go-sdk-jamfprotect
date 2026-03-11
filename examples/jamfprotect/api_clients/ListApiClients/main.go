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

	// List all API clients
	apiClients, _, err := client.ApiClient.ListApiClients(ctx)
	if err != nil {
		log.Fatalf("Failed to list API clients: %v", err)
	}

	fmt.Printf("Found %d API client(s):\n\n", len(apiClients))

	for i, ac := range apiClients {
		fmt.Printf("%d. %s\n", i+1, ac.Name)
		fmt.Printf("   Client ID: %s\n", ac.ClientID)
		fmt.Println()
	}

	os.Exit(0)
}
