package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
	dataforwarding "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/data_forwarding"
)

func main() {
	// Create client from environment variables
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Update data forwarding settings with S3, Sentinel, and SentinelV2 disabled
	request := &dataforwarding.UpdateDataForwardingRequest{
		S3: dataforwarding.ForwardS3Input{
			Enabled: false,
			Bucket:  "",
			Prefix:  "",
			Role:    "",
		},
		Sentinel: dataforwarding.ForwardSentinelInput{
			Enabled:    false,
			CustomerID: "",
			SharedKey:  "",
			LogType:    "",
			Domain:     "",
		},
		SentinelV2: dataforwarding.ForwardSentinelV2Input{
			Enabled:       false,
			AzureTenantID: "",
			AzureClientID: "",
			Endpoint:      "",
			Alerts:        dataforwarding.DataStreamInput{Enabled: false},
			ULogs:         dataforwarding.DataStreamInput{Enabled: false},
			Telemetries:   dataforwarding.DataStreamInput{Enabled: false},
			TelemetriesV2: dataforwarding.DataStreamInput{Enabled: false},
		},
	}

	updated, _, err := client.DataForwarding.UpdateDataForwarding(ctx, request)
	if err != nil {
		log.Fatalf("Failed to update data forwarding settings: %v", err)
	}

	fmt.Printf("Successfully updated data forwarding settings:\n")
	fmt.Printf("  UUID: %s\n", updated.UUID)
	fmt.Printf("  S3 Enabled: %t\n", updated.Forward.S3.Enabled)
	fmt.Printf("  Sentinel Enabled: %t\n", updated.Forward.Sentinel.Enabled)
	fmt.Printf("  Sentinel V2 Enabled: %t\n", updated.Forward.SentinelV2.Enabled)

	os.Exit(0)
}
