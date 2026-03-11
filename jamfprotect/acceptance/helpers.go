package acceptance

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// SkipIfNotConfigured skips the test when Jamf Protect credentials are not set.
func SkipIfNotConfigured(t *testing.T) {
	t.Helper()
	if !IsConfigured() {
		t.Skip("JAMFPROTECT_CLIENT_ID or JAMFPROTECT_CLIENT_SECRET not set, skipping acceptance test")
	}
}

// RequireClient ensures the shared client is initialised, skipping if
// credentials are absent or initialisation fails.
func RequireClient(t *testing.T) {
	t.Helper()
	SkipIfNotConfigured(t)

	if Client == nil {
		err := InitClient()
		require.NoError(t, err, "Failed to initialise Jamf Protect client")
	}
}

// NewContext creates a context with the configured request timeout.
func NewContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), Config.RequestTimeout)
}

// Cleanup registers a cleanup function, skipped when JAMFPROTECT_SKIP_CLEANUP=true.
func Cleanup(t *testing.T, fn func()) {
	t.Helper()
	if !Config.SkipCleanup {
		t.Cleanup(fn)
	} else if Config.Verbose {
		t.Log("Skipping cleanup (JAMFPROTECT_SKIP_CLEANUP=true)")
	}
}

// LogTestStage logs a named test stage with optional GitHub Actions annotation.
func LogTestStage(t *testing.T, stage, message string, args ...any) {
	t.Helper()
	formatted := message
	if len(args) > 0 {
		formatted = fmt.Sprintf(message, args...)
	}
	if isGitHubActions() {
		fmt.Printf("::notice title=%s::%s\n", stage, formatted)
	}
	if Config.Verbose {
		t.Logf("[%s] %s", stage, formatted)
	}
}

// LogTestSuccess logs a successful step.
func LogTestSuccess(t *testing.T, message string, args ...any) {
	t.Helper()
	formatted := message
	if len(args) > 0 {
		formatted = fmt.Sprintf(message, args...)
	}
	if isGitHubActions() {
		fmt.Printf("::notice title=Success::%s\n", formatted)
	}
	if Config.Verbose {
		t.Logf("OK: %s", formatted)
	}
}

// LogTestWarning logs a non-fatal warning.
func LogTestWarning(t *testing.T, message string, args ...any) {
	t.Helper()
	formatted := message
	if len(args) > 0 {
		formatted = fmt.Sprintf(message, args...)
	}
	if isGitHubActions() {
		fmt.Printf("::warning title=Warning::%s\n", formatted)
	}
	if Config.Verbose {
		t.Logf("WARNING: %s", formatted)
	}
}

// LogCleanupDeleteError logs cleanup delete results.
// A GraphQL "not found" error is treated as expected (resource already deleted).
func LogCleanupDeleteError(t *testing.T, resourceType, id string, err error) {
	t.Helper()
	if err == nil {
		return
	}
	if isNotFoundErr(err) {
		LogTestStage(t, "Cleanup", "%s ID=%s already deleted (not found, expected)", resourceType, id)
		return
	}
	LogTestWarning(t, "Cleanup: failed to delete %s ID=%s: %v", resourceType, id, err)
}

// PollUntil retries fn every interval until it returns true or timeout elapses.
// Used to wait for eventually-consistent API state.
func PollUntil(t *testing.T, timeout, interval time.Duration, fn func() bool) bool {
	t.Helper()
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if fn() {
			return true
		}
		time.Sleep(interval)
	}
	return false
}

// RetryOnNotFound retries an operation when receiving not-found errors,
// with exponential backoff. Returns the last error if all retries fail.
func RetryOnNotFound(t *testing.T, maxRetries int, initialDelay time.Duration, fn func() error) error {
	t.Helper()
	var lastErr error
	delay := initialDelay

	for i := 0; i < maxRetries; i++ {
		lastErr = fn()
		if lastErr == nil {
			return nil
		}

		if !isNotFoundErr(lastErr) {
			return lastErr
		}

		if i < maxRetries-1 {
			if Config.Verbose {
				t.Logf("Resource not found, retry %d/%d: waiting %v before next attempt", i+1, maxRetries, delay)
			}
			time.Sleep(delay)
			delay *= 2
		}
	}

	return lastErr
}

// UniqueName returns a resource name that is unique per test run to avoid
// conflicts with pre-existing data. All test resources are prefixed sdkv2_acc_.
func UniqueName(base string) string {
	return fmt.Sprintf("sdkv2_acc_%s_%d", base, time.Now().UnixMilli())
}

// isNotFoundErr returns true if the error indicates a not-found (404) condition.
func isNotFoundErr(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	return contains(msg, "not found") || contains(msg, "404") || contains(msg, "Not Found")
}

// contains is a case-aware string contains helper.
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr ||
		len(substr) == 0 ||
		indexOf(s, substr) >= 0)
}

// indexOf returns the index of substr in s, or -1.
func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

// isGitHubActions returns true when running inside GitHub Actions.
func isGitHubActions() bool {
	return os.Getenv("GITHUB_ACTIONS") == "true"
}

// ---- env helpers ----

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func getDurationEnv(key string, def time.Duration) time.Duration {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return def
	}
	return d
}

func getBoolEnv(key string, def bool) bool {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		return def
	}
	return b
}
