package user

import (
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
)

// validateUserID checks that id is non-empty.
func validateUserID(id string) error {
	if id == "" {
		return fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}
	return nil
}
