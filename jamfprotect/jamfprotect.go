package jamfprotect

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/zap"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	actionconfigs "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/action_configuration"
	analytics "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/analytic"
	analyticsets "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/analytic_set"
	apiclients "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/api_client"
	changemanagement "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/change_management"
	computers "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/computer"
	preventlists "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/custom_prevent_list"
	dataforwarding "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/data_forwarding"
	dataretention "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/data_retention"
	downloads "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/downloads"
	exceptionsets "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/exception_set"
	groups "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/group"
	identityproviders "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/identity_provider"
	plans "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/plan"
	usbcontrolsets "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/removable_storage_control_set"
	roles "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/role"
	telemetryv2 "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/telemetry"
	unifiedloggingfilters "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/unified_logging_filter"
	users "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/user"
)

// Client is the main entry point for the Jamf Protect API SDK.
// It aggregates all service clients and provides a unified interface.
// Users should interact with the API exclusively through the provided service methods.
type Client struct {
	// transport is the internal HTTP transport layer (not exposed to users)
	transport *client.Transport

	// Services
	ActionConfig         *actionconfigs.Service
	Analytic             *analytics.Service
	AnalyticSet          *analyticsets.Service
	ApiClient            *apiclients.Service
	ChangeManagement     *changemanagement.Service
	Computer             *computers.Service
	DataForwarding       *dataforwarding.Service
	DataRetention        *dataretention.Service
	Downloads            *downloads.Service
	ExceptionSet         *exceptionsets.Service
	Group                *groups.Service
	IdentityProvider     *identityproviders.Service
	Plan                 *plans.Service
	PreventList          *preventlists.Service
	Role                 *roles.Service
	TelemetryV2          *telemetryv2.Service
	UnifiedLoggingFilter *unifiedloggingfilters.Service
	USBControlSet        *usbcontrolsets.Service
	User                 *users.Service
}

// NewClient creates a new Jamf Protect API client
//
// Parameters:
//   - clientID: The Jamf Protect OAuth2 client ID
//   - clientSecret: The Jamf Protect OAuth2 client secret
//   - options: Optional client configuration options
//
// Example:
//
//	client, err := jamfprotect.NewClient(
//	    "your-client-id",
//	    "your-client-secret",
//	    jamfprotect.WithDebug(),
//	)
func NewClient(clientID, clientSecret string, options ...client.ClientOption) (*Client, error) {
	transport, err := client.NewTransport(clientID, clientSecret, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP transport: %w", err)
	}

	// Initialize service clients
	c := &Client{
		transport:            transport,
		ActionConfig:         actionconfigs.NewService(transport),
		Analytic:             analytics.NewService(transport),
		AnalyticSet:          analyticsets.NewService(transport),
		ApiClient:            apiclients.NewService(transport),
		ChangeManagement:     changemanagement.NewService(transport),
		Computer:             computers.NewService(transport),
		DataForwarding:       dataforwarding.NewService(transport),
		DataRetention:        dataretention.NewService(transport),
		Downloads:            downloads.NewService(transport),
		ExceptionSet:         exceptionsets.NewService(transport),
		Group:                groups.NewService(transport),
		IdentityProvider:     identityproviders.NewService(transport),
		Plan:                 plans.NewService(transport),
		PreventList:          preventlists.NewService(transport),
		Role:                 roles.NewService(transport),
		TelemetryV2:          telemetryv2.NewService(transport),
		UnifiedLoggingFilter: unifiedloggingfilters.NewService(transport),
		USBControlSet:        usbcontrolsets.NewService(transport),
		User:                 users.NewService(transport),
	}

	return c, nil
}

// NewClientFromEnv creates a new client using environment variables
//
// Required environment variables:
//   - JAMFPROTECT_CLIENT_ID: The OAuth2 client ID
//   - JAMFPROTECT_CLIENT_SECRET: The OAuth2 client secret
//
// Optional environment variables:
//   - JAMFPROTECT_BASE_URL: Custom base URL (defaults to https://apis.jamfprotect.cloud)
//
// Example:
//
//	client, err := jamfprotect.NewClientFromEnv()
func NewClientFromEnv(options ...client.ClientOption) (*Client, error) {
	clientID := os.Getenv("JAMFPROTECT_CLIENT_ID")
	if clientID == "" {
		return nil, fmt.Errorf("JAMFPROTECT_CLIENT_ID environment variable is required")
	}

	clientSecret := os.Getenv("JAMFPROTECT_CLIENT_SECRET")
	if clientSecret == "" {
		return nil, fmt.Errorf("JAMFPROTECT_CLIENT_SECRET environment variable is required")
	}

	// Check for optional environment variables and append to options
	if baseURL := os.Getenv("JAMFPROTECT_BASE_URL"); baseURL != "" {
		options = append(options, client.WithBaseURL(baseURL))
	}

	return NewClient(clientID, clientSecret, options...)
}

// GetLogger returns the configured zap logger instance.
// Use this to add custom logging within your application using the same logger.
//
// Returns:
//   - *zap.Logger: The configured logger instance
func (c *Client) GetLogger() *zap.Logger {
	return c.transport.GetLogger()
}

// GetTransport returns the underlying transport layer.
// This is useful for advanced configuration like setting custom loggers at runtime.
//
// Returns:
//   - *client.Transport: The transport instance
func (c *Client) GetTransport() *client.Transport {
	return c.transport
}

// GetTokenManager returns the token manager for advanced token lifecycle operations.
//
// Returns:
//   - *client.TokenManager: The token manager instance
func (c *Client) GetTokenManager() *client.TokenManager {
	return c.transport.GetTokenManager()
}

// RefreshToken manually refreshes the OAuth2 access token.
// Useful when you need to explicitly force a token refresh ahead of expiry.
//
// Parameters:
//   - ctx: Request context
//
// Returns:
//   - error: Any error encountered during token refresh
func (c *Client) RefreshToken(ctx context.Context) error {
	return c.transport.RefreshToken(ctx)
}

// InvalidateToken invalidates the current cached token, forcing a refresh on next use.
func (c *Client) InvalidateToken() {
	c.transport.InvalidateToken()
}
