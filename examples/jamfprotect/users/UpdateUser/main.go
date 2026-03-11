package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
	user "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/user"
)

func main() {
	// Create client from environment variables
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Update a user by ID
	userID := "user-id-here" // Replace with actual user ID

	request := &user.UpdateUserRequest{
		ReceiveEmailAlert:     true,
		EmailAlertMinSeverity: "HIGH",
		RoleIDs:               []string{},
		GroupIDs:              []string{},
	}

	updated, _, err := client.User.UpdateUser(ctx, userID, request)
	if err != nil {
		log.Fatalf("Failed to update user: %v", err)
	}

	fmt.Printf("Successfully updated user:\n")
	fmt.Printf("  ID: %s\n", updated.ID)
	fmt.Printf("  Email: %s\n", updated.Email)
	fmt.Printf("  Receive Email Alert: %t\n", updated.ReceiveEmailAlert)
	fmt.Printf("  Email Alert Min Severity: %s\n", updated.EmailAlertMinSeverity)

	os.Exit(0)
}
