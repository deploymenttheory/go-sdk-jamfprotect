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

	// Get a user by ID
	userID := "user-id-here" // Replace with actual user ID

	u, _, err := client.User.GetUser(ctx, userID)
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}

	fmt.Printf("User details:\n")
	fmt.Printf("  ID: %s\n", u.ID)
	fmt.Printf("  Email: %s\n", u.Email)
	fmt.Printf("  Source: %s\n", u.Source)
	fmt.Printf("  Receive Email Alert: %t\n", u.ReceiveEmailAlert)
	fmt.Printf("  Assigned Roles: %d\n", len(u.AssignedRoles))
	fmt.Printf("  Assigned Groups: %d\n", len(u.AssignedGroups))

	os.Exit(0)
}
