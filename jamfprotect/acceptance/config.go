package acceptance

import (
	"fmt"
	"log"
	"time"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
)

// TestConfig holds configuration for acceptance tests driven by environment variables.
type TestConfig struct {
	// Credentials
	ClientID     string
	ClientSecret string

	// Optional
	BaseURL string

	// Test behaviour
	RequestTimeout time.Duration
	SkipCleanup    bool
	Verbose        bool
}

var (
	// Config is the global acceptance test configuration, initialised from env.
	Config *TestConfig
	// Client is the shared Jamf Protect SDK client for acceptance tests.
	Client *jamfprotect.Client
)

func init() {
	Config = &TestConfig{
		ClientID:       getEnv("JAMFPROTECT_CLIENT_ID", ""),
		ClientSecret:   getEnv("JAMFPROTECT_CLIENT_SECRET", ""),
		BaseURL:        getEnv("JAMFPROTECT_BASE_URL", ""),
		RequestTimeout: getDurationEnv("JAMFPROTECT_REQUEST_TIMEOUT", 30*time.Second),
		SkipCleanup:    getBoolEnv("JAMFPROTECT_SKIP_CLEANUP", false),
		Verbose:        getBoolEnv("JAMFPROTECT_VERBOSE", false),
	}
}

// InitClient creates the shared Jamf Protect client from environment variables.
// Returns an error if required credentials are absent.
func InitClient() error {
	if Config.ClientID == "" {
		return fmt.Errorf("JAMFPROTECT_CLIENT_ID environment variable is required")
	}
	if Config.ClientSecret == "" {
		return fmt.Errorf("JAMFPROTECT_CLIENT_SECRET environment variable is required")
	}

	opts := []client.ClientOption{}
	if Config.BaseURL != "" {
		opts = append(opts, client.WithBaseURL(Config.BaseURL))
	}

	var err error
	Client, err = jamfprotect.NewClient(Config.ClientID, Config.ClientSecret, opts...)
	if err != nil {
		return fmt.Errorf("failed to create Jamf Protect client: %w", err)
	}

	if Config.Verbose {
		log.Printf("Acceptance test client initialised (client_id=%s)", Config.ClientID)
	}
	return nil
}

// IsConfigured returns true if the minimum required credentials are set.
func IsConfigured() bool {
	return Config.ClientID != "" && Config.ClientSecret != ""
}
