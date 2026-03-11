package dataretention

import (
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
)

// ValidateUpdateDataRetentionRequest validates that retention days values are positive.
func ValidateUpdateDataRetentionRequest(req *UpdateDataRetentionRequest) error {
	if req == nil {
		return nil
	}
	if req.DatabaseLogDays <= 0 {
		return fmt.Errorf("%w: databaseLogDays must be a positive integer", client.ErrInvalidInput)
	}
	if req.DatabaseAlertDays <= 0 {
		return fmt.Errorf("%w: databaseAlertDays must be a positive integer", client.ErrInvalidInput)
	}
	if req.ColdAlertDays <= 0 {
		return fmt.Errorf("%w: coldAlertDays must be a positive integer", client.ErrInvalidInput)
	}
	return nil
}
