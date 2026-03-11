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

	// Create a new role
	request := &role.CreateRoleRequest{
		Name:           "Example Role",
		ReadResources:  []string{"ANALYTIC"},
		WriteResources: []string{},
	}

	created, _, err := client.Role.CreateRole(ctx, request)
	if err != nil {
		log.Fatalf("Failed to create role: %v", err)
	}

	fmt.Printf("Successfully created role:\n")
	fmt.Printf("  ID: %s\n", created.ID)
	fmt.Printf("  Name: %s\n", created.Name)
	fmt.Printf("  Read Permissions: %v\n", created.Permissions.Read)
	fmt.Printf("  Write Permissions: %v\n", created.Permissions.Write)

	os.Exit(0)
}
