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

	// List all groups
	groups, _, err := client.Group.ListGroups(ctx)
	if err != nil {
		log.Fatalf("Failed to list groups: %v", err)
	}

	fmt.Printf("Found %d group(s):\n\n", len(groups))

	for i, g := range groups {
		fmt.Printf("%d. %s\n", i+1, g.Name)
		fmt.Printf("   ID: %s\n", g.ID)
		fmt.Println()
	}

	os.Exit(0)
}
