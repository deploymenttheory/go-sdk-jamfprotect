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

	// List all computers
	computers, _, err := client.Computer.ListComputers(ctx)
	if err != nil {
		log.Fatalf("Failed to list computers: %v", err)
	}

	fmt.Printf("Found %d computer(s):\n\n", len(computers))

	for i, computer := range computers {
		fmt.Printf("%d.", i+1)
		if computer.HostName != nil {
			fmt.Printf(" %s", *computer.HostName)
		}
		fmt.Println()
		if computer.UUID != nil {
			fmt.Printf("   UUID: %s\n", *computer.UUID)
		}
		if computer.OSString != nil {
			fmt.Printf("   OS: %s\n", *computer.OSString)
		}
		fmt.Println()
	}

	os.Exit(0)
}
