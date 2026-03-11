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

	// Get a computer by UUID
	uuid := "computer-uuid-here" // Replace with actual computer UUID

	computer, _, err := client.Computer.GetComputer(ctx, uuid)
	if err != nil {
		log.Fatalf("Failed to get computer: %v", err)
	}

	fmt.Printf("Computer details:\n")
	if computer.UUID != nil {
		fmt.Printf("  UUID: %s\n", *computer.UUID)
	}
	if computer.Serial != nil {
		fmt.Printf("  Serial: %s\n", *computer.Serial)
	}
	if computer.HostName != nil {
		fmt.Printf("  Host Name: %s\n", *computer.HostName)
	}
	if computer.OSString != nil {
		fmt.Printf("  OS: %s\n", *computer.OSString)
	}
	if computer.Version != nil {
		fmt.Printf("  Version: %s\n", *computer.Version)
	}
	if computer.Plan != nil && computer.Plan.Name != nil {
		fmt.Printf("  Plan: %s\n", *computer.Plan.Name)
	}

	os.Exit(0)
}
