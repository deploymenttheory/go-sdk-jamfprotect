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

	// Get an API client by clientId
	clientID := "api-client-id-here" // Replace with actual client ID

	apiClient, _, err := client.ApiClient.GetApiClient(ctx, clientID)
	if err != nil {
		log.Fatalf("Failed to get API client: %v", err)
	}

	fmt.Printf("API client details:\n")
	fmt.Printf("  Client ID: %s\n", apiClient.ClientID)
	fmt.Printf("  Name: %s\n", apiClient.Name)
	fmt.Printf("  Assigned Roles: %d\n", len(apiClient.AssignedRoles))
	for _, role := range apiClient.AssignedRoles {
		fmt.Printf("    - %s (%s)\n", role.Name, role.ID)
	}

	os.Exit(0)
}
