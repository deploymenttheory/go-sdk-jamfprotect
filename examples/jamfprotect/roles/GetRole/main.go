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

	// Get a role by ID
	roleID := "role-id-here" // Replace with actual role ID

	r, _, err := client.Role.GetRole(ctx, roleID)
	if err != nil {
		log.Fatalf("Failed to get role: %v", err)
	}

	fmt.Printf("Role details:\n")
	fmt.Printf("  ID: %s\n", r.ID)
	fmt.Printf("  Name: %s\n", r.Name)
	fmt.Printf("  Read Permissions: %v\n", r.Permissions.Read)
	fmt.Printf("  Write Permissions: %v\n", r.Permissions.Write)

	os.Exit(0)
}
