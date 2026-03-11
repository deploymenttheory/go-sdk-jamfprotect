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

	// Delete an API client by clientId
	clientID := "api-client-id-here" // Replace with actual client ID

	_, err = client.ApiClient.DeleteApiClient(ctx, clientID)
	if err != nil {
		log.Fatalf("Failed to delete API client: %v", err)
	}

	fmt.Printf("Successfully deleted API client with client ID: %s\n", clientID)

	os.Exit(0)
}
