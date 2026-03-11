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

	// Delete a user by ID
	userID := "user-id-here" // Replace with actual user ID

	_, err = client.User.DeleteUser(ctx, userID)
	if err != nil {
		log.Fatalf("Failed to delete user: %v", err)
	}

	fmt.Printf("Successfully deleted user with ID: %s\n", userID)

	os.Exit(0)
}
