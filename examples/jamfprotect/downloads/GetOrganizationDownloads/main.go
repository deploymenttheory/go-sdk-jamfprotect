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

	// Get organization downloads
	downloads, _, err := client.Downloads.GetOrganizationDownloads(ctx)
	if err != nil {
		log.Fatalf("Failed to get organization downloads: %v", err)
	}

	fmt.Printf("Organization downloads:\n")
	fmt.Printf("  Installer UUID: %s\n", downloads.InstallerUUID)
	if downloads.VanillaPackage != nil {
		fmt.Printf("  Vanilla Package Version: %s\n", downloads.VanillaPackage.Version)
	}
	fmt.Printf("  PPPC non-empty: %t\n", downloads.PPPC != "")
	fmt.Printf("  Root CA non-empty: %t\n", downloads.RootCA != "")
	fmt.Printf("  CSR non-empty: %t\n", downloads.CSR != "")

	os.Exit(0)
}
