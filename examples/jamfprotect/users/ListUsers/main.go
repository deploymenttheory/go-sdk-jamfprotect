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

	// List all users
	users, _, err := client.User.ListUsers(ctx)
	if err != nil {
		log.Fatalf("Failed to list users: %v", err)
	}

	fmt.Printf("Found %d user(s):\n\n", len(users))

	for i, u := range users {
		fmt.Printf("%d. %s\n", i+1, u.Email)
		fmt.Printf("   ID: %s\n", u.ID)
		fmt.Println()
	}

	os.Exit(0)
}
