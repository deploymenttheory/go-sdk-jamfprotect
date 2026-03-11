package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
	group "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/group"
)

func main() {
	// Create client from environment variables
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Update a group by ID
	groupID := "group-id-here" // Replace with actual group ID

	request := &group.UpdateGroupRequest{
		Name:        "Updated Group",
		AccessGroup: false,
		RoleIDs:     []string{},
	}

	updated, _, err := client.Group.UpdateGroup(ctx, groupID, request)
	if err != nil {
		log.Fatalf("Failed to update group: %v", err)
	}

	fmt.Printf("Successfully updated group:\n")
	fmt.Printf("  ID: %s\n", updated.ID)
	fmt.Printf("  Name: %s\n", updated.Name)
	fmt.Printf("  Access Group: %t\n", updated.AccessGroup)

	os.Exit(0)
}
