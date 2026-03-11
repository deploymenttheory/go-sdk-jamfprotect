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

	// List all identity provider connections
	connections, _, err := client.IdentityProvider.ListConnections(ctx)
	if err != nil {
		log.Fatalf("Failed to list connections: %v", err)
	}

	fmt.Printf("Found %d connection(s):\n\n", len(connections))

	for i, conn := range connections {
		fmt.Printf("%d. %s\n", i+1, conn.Name)
		fmt.Printf("   ID: %s\n", conn.ID)
		fmt.Printf("   Strategy: %s\n", conn.Strategy)
		fmt.Printf("   Source: %s\n", conn.Source)
		fmt.Println()
	}

	os.Exit(0)
}
