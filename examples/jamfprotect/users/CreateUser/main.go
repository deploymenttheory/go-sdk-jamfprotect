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

	// Create a new user
	request := &user.CreateUserRequest{
		Email:                 "example@company.com",
		ReceiveEmailAlert:     false,
		EmailAlertMinSeverity: "LOW",
		RoleIDs:               []string{},
		GroupIDs:              []string{},
	}

	created, _, err := client.User.CreateUser(ctx, request)
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	fmt.Printf("Successfully created user:\n")
	fmt.Printf("  ID: %s\n", created.ID)
	fmt.Printf("  Email: %s\n", created.Email)

	os.Exit(0)
}
