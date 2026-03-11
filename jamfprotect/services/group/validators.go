package group

import (
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
)

// validateGroupID checks that id is non-empty.
func validateGroupID(id string) error {
	if id == "" {
		return fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}
	return nil
}
