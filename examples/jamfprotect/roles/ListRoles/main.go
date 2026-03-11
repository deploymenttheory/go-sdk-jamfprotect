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

	// List all roles
	roles, _, err := client.Role.ListRoles(ctx)
	if err != nil {
		log.Fatalf("Failed to list roles: %v", err)
	}

	fmt.Printf("Found %d role(s):\n\n", len(roles))

	for i, r := range roles {
		fmt.Printf("%d. %s\n", i+1, r.Name)
		fmt.Printf("   ID: %s\n", r.ID)
		fmt.Println()
	}

	os.Exit(0)
}
