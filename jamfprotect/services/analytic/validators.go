package analytic

import (
	"fmt"
	"regexp"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/validate"
)

// uuidRegex matches a canonical UUID string (8-4-4-4-12 hex digits).
var uuidRegex = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

// Allowed values from provider schema / API enums.

const (
	InputTypeGPFSEvent             = "GPFSEvent"
	InputTypeGPDownloadEvent       = "GPDownloadEvent"
	InputTypeGPProcessEvent        = "GPProcessEvent"
	InputTypeGPScreenshotEvent     = "GPScreenshotEvent"
	InputTypeGPKeylogRegisterEvent = "GPKeylogRegisterEvent"
	InputTypeGPClickEvent          = "GPClickEvent"
	InputTypeGPMRTEvent            = "GPMRTEvent"
	InputTypeGPUSBEvent            = "GPUSBEvent"
	InputTypeGPGatekeeperEvent     = "GPGatekeeperEvent"
)

var allowedInputTypes = []string{
	InputTypeGPFSEvent, InputTypeGPDownloadEvent, InputTypeGPProcessEvent,
	InputTypeGPScreenshotEvent, InputTypeGPKeylogRegisterEvent, InputTypeGPClickEvent,
	InputTypeGPMRTEvent, InputTypeGPUSBEvent, InputTypeGPGatekeeperEvent,
}

const (
	SeverityHigh          = "High"
	SeverityMedium        = "Medium"
	SeverityLow           = "Low"
	SeverityInformational = "Informational"
)

// ValidateInputType validates analytic input type (sensor type) is an allowed enum value.
func ValidateInputType(inputType string) error {
	return validate.OneOf("inputType", inputType, allowedInputTypes...)
}

// ValidateLevel validates analytic level is in allowed range 0-10.
func ValidateLevel(level int) error {
	return validate.IntBetween("level", level, 0, 10)
}

// ValidateSeverity validates analytic severity is an allowed enum value.
func ValidateSeverity(severity string) error {
	return validate.OneOf("severity", severity, SeverityHigh, SeverityMedium, SeverityLow, SeverityInformational)
}

// ValidateCreateAnalyticRequest validates allowed-value constraints on create analytic request.
func ValidateCreateAnalyticRequest(req *CreateAnalyticRequest) error {
	if req == nil {
		return nil
	}
	if err := ValidateInputType(req.InputType); err != nil {
		return err
	}
	if err := ValidateLevel(req.Level); err != nil {
		return err
	}
	if err := ValidateSeverity(req.Severity); err != nil {
		return err
	}
	return nil
}

// ValidateUpdateAnalyticRequest validates allowed-value constraints on update analytic request.
func ValidateUpdateAnalyticRequest(req *UpdateAnalyticRequest) error {
	if req == nil {
		return nil
	}
	if req.InputType != "" {
		if err := ValidateInputType(req.InputType); err != nil {
			return err
		}
	}
	// Level 0 is valid; check only if caller sets a level (e.g. non-zero or explicit 0)
	// Level is int - we always validate if it's in range when present in request
	if err := ValidateLevel(req.Level); err != nil {
		return err
	}
	if req.Severity != nil && *req.Severity != "" {
		if err := ValidateSeverity(*req.Severity); err != nil {
			return err
		}
	}
	return nil
}

// ValidateUpdateInternalAnalyticRequest validates tenant-scoped update request fields.
func ValidateUpdateInternalAnalyticRequest(req *UpdateInternalAnalyticRequest) error {
	if req == nil {
		return nil
	}
	if req.TenantSeverity != "" {
		if err := ValidateSeverity(req.TenantSeverity); err != nil {
			return err
		}
	}
	return nil
}

// ValidateAnalyticID checks that uuid is non-empty and matches UUID format.
func ValidateAnalyticID(uuid string) error {
	if uuid == "" {
		return fmt.Errorf("%w: uuid is required", client.ErrInvalidInput)
	}
	if !uuidRegex.MatchString(uuid) {
		return fmt.Errorf("%w: uuid must be a valid UUID (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)", client.ErrInvalidInput)
	}
	return nil
}
