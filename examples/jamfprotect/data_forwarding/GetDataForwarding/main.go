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

	// Get data forwarding settings
	settings, _, err := client.DataForwarding.GetDataForwarding(ctx)
	if err != nil {
		log.Fatalf("Failed to get data forwarding settings: %v", err)
	}

	fmt.Printf("Data forwarding settings:\n")
	fmt.Printf("  UUID: %s\n", settings.UUID)
	fmt.Printf("  S3:\n")
	fmt.Printf("    Enabled: %t\n", settings.Forward.S3.Enabled)
	fmt.Printf("    Bucket: %s\n", settings.Forward.S3.Bucket)
	fmt.Printf("  Sentinel:\n")
	fmt.Printf("    Enabled: %t\n", settings.Forward.Sentinel.Enabled)
	fmt.Printf("  Sentinel V2:\n")
	fmt.Printf("    Enabled: %t\n", settings.Forward.SentinelV2.Enabled)

	os.Exit(0)
}
