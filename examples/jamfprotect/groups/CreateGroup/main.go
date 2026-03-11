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

	// Create a new group
	request := &group.CreateGroupRequest{
		Name:        "Example Group",
		AccessGroup: false,
		RoleIDs:     []string{},
	}

	created, _, err := client.Group.CreateGroup(ctx, request)
	if err != nil {
		log.Fatalf("Failed to create group: %v", err)
	}

	fmt.Printf("Successfully created group:\n")
	fmt.Printf("  ID: %s\n", created.ID)
	fmt.Printf("  Name: %s\n", created.Name)

	os.Exit(0)
}
