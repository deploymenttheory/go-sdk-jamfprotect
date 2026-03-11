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

	// Create a new API client
	request := &apiclient.CreateApiClientRequest{
		Name:    "Example API Client",
		RoleIDs: []string{},
	}

	created, _, err := client.ApiClient.CreateApiClient(ctx, request)
	if err != nil {
		log.Fatalf("Failed to create API client: %v", err)
	}

	fmt.Printf("Successfully created API client:\n")
	fmt.Printf("  Client ID: %s\n", created.ClientID)
	fmt.Printf("  Name: %s\n", created.Name)
	fmt.Printf("  Password: %s\n", created.Password)

	os.Exit(0)
}
