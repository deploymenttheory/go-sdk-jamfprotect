package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
	apiclient "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/api_client"
)

func main() {
	// Create client from environment variables
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Update an API client by clientId
	clientID := "api-client-id-here" // Replace with actual client ID

	request := &apiclient.UpdateApiClientRequest{
		Name:    "Updated API Client",
		RoleIDs: []string{},
	}

	updated, _, err := client.ApiClient.UpdateApiClient(ctx, clientID, request)
	if err != nil {
		log.Fatalf("Failed to update API client: %v", err)
	}

	fmt.Printf("Successfully updated API client:\n")
	fmt.Printf("  Client ID: %s\n", updated.ClientID)
	fmt.Printf("  Name: %s\n", updated.Name)
	fmt.Printf("  Assigned Roles: %d\n", len(updated.AssignedRoles))

	os.Exit(0)
}
