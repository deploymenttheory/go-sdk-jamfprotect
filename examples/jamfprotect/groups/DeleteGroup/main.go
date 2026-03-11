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

	// Delete a group by ID
	groupID := "group-id-here" // Replace with actual group ID

	_, err = client.Group.DeleteGroup(ctx, groupID)
	if err != nil {
		log.Fatalf("Failed to delete group: %v", err)
	}

	fmt.Printf("Successfully deleted group with ID: %s\n", groupID)

	os.Exit(0)
}
