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

	// Delete a role by ID
	roleID := "role-id-here" // Replace with actual role ID

	_, err = client.Role.DeleteRole(ctx, roleID)
	if err != nil {
		log.Fatalf("Failed to delete role: %v", err)
	}

	fmt.Printf("Successfully deleted role with ID: %s\n", roleID)

	os.Exit(0)
}
