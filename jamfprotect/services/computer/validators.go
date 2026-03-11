package computer

import (
	"fmt"
	"regexp"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
)

// uuidRegex matches a canonical UUID string (8-4-4-4-12 hex digits).
var uuidRegex = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

// ValidateComputerUUID checks that uuid is non-empty and matches UUID format.
func ValidateComputerUUID(uuid string) error {
	if uuid == "" {
		return fmt.Errorf("%w: uuid is required", client.ErrInvalidInput)
	}
	if !uuidRegex.MatchString(uuid) {
		return fmt.Errorf("%w: uuid must be a valid UUID (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)", client.ErrInvalidInput)
	}
	return nil
}
