package role

import (
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
)

// validateRoleID checks that id is non-empty.
func validateRoleID(id string) error {
	if id == "" {
		return fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}
	return nil
}
