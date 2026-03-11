package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
	role "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/role"
)

func main() {
	// Create client from environment variables
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Update a role by ID
	roleID := "role-id-here" // Replace with actual role ID

	request := &role.UpdateRoleRequest{
		Name:           "Updated Role",
		ReadResources:  []string{"ANALYTIC"},
		WriteResources: []string{},
	}

	updated, _, err := client.Role.UpdateRole(ctx, roleID, request)
	if err != nil {
		log.Fatalf("Failed to update role: %v", err)
	}

	fmt.Printf("Successfully updated role:\n")
	fmt.Printf("  ID: %s\n", updated.ID)
	fmt.Printf("  Name: %s\n", updated.Name)
	fmt.Printf("  Read Permissions: %v\n", updated.Permissions.Read)
	fmt.Printf("  Write Permissions: %v\n", updated.Permissions.Write)

	os.Exit(0)
}
