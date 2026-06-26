package main

import (
	"context"
	"fmt"
	"log"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
)

func main() {
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	statuses, _, err := client.Beta.GetBetaAcceptanceStatus(ctx)
	if err != nil {
		log.Fatalf("Failed to get beta acceptance status: %v", err)
	}

	if len(statuses) == 0 {
		fmt.Println("No beta programs enrolled")
		return
	}

	for _, status := range statuses {
		fmt.Printf("Beta: %s\n", status.BetaName)
		fmt.Printf("  Accepted: %s by %s\n", status.AcceptedTimestamp, status.AcceptedUser)
	}
}
