package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestAPIError_Error(t *testing.T) {
	e := &APIError{
		Code:       "ERR",
		Message:    "error message",
		StatusCode: 404,
		Status:     "Not Found",
		Method:     "GET",
		Endpoint:   "/api/test",
	}
	s := e.Error()
	assert.NotEmpty(t, s)
	assert.Contains(t, s, "404")
	assert.Contains(t, s, "error message")
	assert.Contains(t, s, "GET")
	assert.Contains(t, s, "/api/test")
}

func TestAPIError_Error_NoCode(t *testing.T) {
	e := &APIError{
		Message:    "error message",
		StatusCode: 500,
		Status:     "Internal Server Error",
		Method:     "POST",
		Endpoint:   "/api/test",
	}
	s := e.Error()
	assert.NotEmpty(t, s)
	assert.Contains(t, s, "500")
	assert.Contains(t, s, "error message")
}

func TestParseErrorResponse(t *testing.T) {
	logger := zap.NewNop()
	tests := []struct {
		name       string
		body       []byte
		statusCode int
		status     string
		wantMsg    string
	}{
		{"json body", []byte(`{"error":{"code":"ERR","message":"not found"}}`), 404, "Not Found", "not found"},
		{"empty body", []byte(``), 404, "Not Found", "The requested resource was not found."},
		{"plain body", []byte("plain error"), 500, "Internal", "plain error"},
		{"default 400", nil, 400, "Bad Request", "The request is invalid or malformed."},
		{"default 401", nil, 401, "Unauthorized", "Authentication required or token invalid."},
		{"default 403", nil, 403, "Forbidden", "You are not allowed to perform the requested operation."},
		{"default 409", nil, 409, "Conflict", "The resource already exists or conflicts with current state."},
		{"default 422", nil, 422, "Unprocessable", "Validation error."},
		{"default 429", nil, 429, "Too Many", "Rate limit exceeded. Retry after the indicated period."},
		{"default 500", nil, 500, "Internal", "Internal server error."},
		{"default 502", nil, 502, "Bad Gateway", "Bad gateway."},
		{"default 503", nil, 503, "Unavailable", "Service temporarily unavailable. Retry later."},
		{"default 504", nil, 504, "Timeout", "The operation took too long to complete."},
		{"unknown status", nil, 418, "Teapot", "Unknown error."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ParseErrorResponse(tt.body, tt.statusCode, tt.status, "GET", "/api", logger)
			require.Error(t, err)
			apiErr, ok := err.(*APIError)
			require.True(t, ok, "expected *APIError, got %T", err)
			assert.Equal(t, tt.wantMsg, apiErr.Message)
		})
	}
}

func TestIsNotFound(t *testing.T) {
	assert.False(t, IsNotFound(nil))
	assert.False(t, IsNotFound(&APIError{StatusCode: 500}))
	assert.True(t, IsNotFound(&APIError{StatusCode: http.StatusNotFound}))
}

func TestIsUnauthorized(t *testing.T) {
	assert.True(t, IsUnauthorized(&APIError{StatusCode: http.StatusUnauthorized}))
	assert.False(t, IsUnauthorized(&APIError{StatusCode: 200}))
	assert.False(t, IsUnauthorized(nil))
}

func TestIsBadRequest(t *testing.T) {
	assert.True(t, IsBadRequest(&APIError{StatusCode: http.StatusBadRequest}))
	assert.False(t, IsBadRequest(&APIError{StatusCode: 200}))
	assert.False(t, IsBadRequest(nil))
}

func TestIsForbidden(t *testing.T) {
	assert.True(t, IsForbidden(&APIError{StatusCode: http.StatusForbidden}))
	assert.False(t, IsForbidden(&APIError{StatusCode: 200}))
	assert.False(t, IsForbidden(nil))
}

func TestIsConflict(t *testing.T) {
	assert.True(t, IsConflict(&APIError{StatusCode: http.StatusConflict}))
	assert.False(t, IsConflict(&APIError{StatusCode: 200}))
	assert.False(t, IsConflict(nil))
}

func TestIsValidationError(t *testing.T) {
	assert.True(t, IsValidationError(&APIError{StatusCode: http.StatusUnprocessableEntity}))
	assert.False(t, IsValidationError(&APIError{StatusCode: 200}))
	assert.False(t, IsValidationError(nil))
}

func TestIsRateLimited(t *testing.T) {
	assert.True(t, IsRateLimited(&APIError{StatusCode: http.StatusTooManyRequests}))
	assert.False(t, IsRateLimited(&APIError{StatusCode: 200}))
	assert.False(t, IsRateLimited(nil))
}

func TestIsServerError(t *testing.T) {
	assert.True(t, IsServerError(&APIError{StatusCode: 500}))
	assert.True(t, IsServerError(&APIError{StatusCode: 503}))
	assert.False(t, IsServerError(&APIError{StatusCode: 400}))
	assert.False(t, IsServerError(nil))
}

func TestIsTransient(t *testing.T) {
	assert.True(t, IsTransient(&APIError{StatusCode: http.StatusServiceUnavailable}))
	assert.True(t, IsTransient(&APIError{StatusCode: http.StatusGatewayTimeout}))
	assert.True(t, IsTransient(&APIError{Code: ErrorCodeTransient, StatusCode: 400}))
	assert.False(t, IsTransient(&APIError{StatusCode: http.StatusTooManyRequests}))
	assert.False(t, IsTransient(&APIError{StatusCode: 400}))
	assert.False(t, IsTransient(nil))
}

func TestGetErrorCode(t *testing.T) {
	err := &APIError{Code: "NotFound"}
	assert.Equal(t, "NotFound", GetErrorCode(err))

	assert.Equal(t, "", GetErrorCode(nil))
	assert.Equal(t, "", GetErrorCode(&APIError{Code: ""}))
}

func TestIsGraphQL(t *testing.T) {
	assert.True(t, IsGraphQL(&APIError{Code: ErrorCodeGraphQL}))
	assert.False(t, IsGraphQL(&APIError{Code: "OtherCode"}))
	assert.False(t, IsGraphQL(nil))
}
