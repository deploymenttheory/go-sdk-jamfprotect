package apiclient

import (
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
)

// validateClientID checks that clientId is non-empty.
func validateClientID(clientID string) error {
	if clientID == "" {
		return fmt.Errorf("%w: clientId is required", client.ErrInvalidInput)
	}
	return nil
}
