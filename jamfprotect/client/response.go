package client

import (
	"encoding/json"
	"fmt"
	"strings"

	"go.uber.org/zap"
	"resty.dev/v3"
)

// GraphQLResponse represents a GraphQL response payload, including any errors.
type GraphQLResponse struct {
	Data   json.RawMessage `json:"data"`
	Errors []GraphQLError  `json:"errors"`
}

// GraphQLError represents an individual error returned by the GraphQL API.
// Jamf Protect may include errorType (e.g. ArgumentValidationError), data, and errorInfo.
type GraphQLError struct {
	Message    string            `json:"message"`
	Path       []any             `json:"path,omitempty"`
	Data       json.RawMessage   `json:"data,omitempty"`
	ErrorType  string            `json:"errorType,omitempty"`
	ErrorInfo  map[string]any    `json:"errorInfo,omitempty"`
	Locations  []GraphQLLocation `json:"locations,omitempty"`
	Extensions map[string]any    `json:"extensions,omitempty"`
}

// GraphQLLocation represents the line and column of an error in a GraphQL query.
type GraphQLLocation struct {
	Line       int    `json:"line"`
	Column     int    `json:"column"`
	SourceName string `json:"sourceName,omitempty"`
}

// validateResponse validates the HTTP response before processing.
func (t *Transport) validateResponse(resp *resty.Response, method, path string) error {
	bodyLen := len(resp.String())
	if resp.Header().Get("Content-Length") == "0" || bodyLen == 0 {
		t.logger.Debug("Empty response received",
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("status_code", resp.StatusCode()))
		return nil
	}
	if !resp.IsError() && bodyLen > 0 {
		contentType := resp.Header().Get("Content-Type")
		if contentType != "" && !strings.HasPrefix(contentType, "application/json") {
			t.logger.Warn("Unexpected Content-Type in response",
				zap.String("method", method),
				zap.String("path", path),
				zap.String("content_type", contentType),
				zap.String("expected", "application/json"))
			return fmt.Errorf("unexpected response Content-Type from %s %s: got %q, expected application/json",
				method, path, contentType)
		}
	}
	return nil
}
