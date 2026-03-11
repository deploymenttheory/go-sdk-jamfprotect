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

	// Get a group by ID
	groupID := "group-id-here" // Replace with actual group ID

	g, _, err := client.Group.GetGroup(ctx, groupID)
	if err != nil {
		log.Fatalf("Failed to get group: %v", err)
	}

	fmt.Printf("Group details:\n")
	fmt.Printf("  ID: %s\n", g.ID)
	fmt.Printf("  Name: %s\n", g.Name)
	fmt.Printf("  Assigned Roles: %d\n", len(g.AssignedRoles))
	fmt.Printf("  Access Group: %t\n", g.AccessGroup)

	os.Exit(0)
}
