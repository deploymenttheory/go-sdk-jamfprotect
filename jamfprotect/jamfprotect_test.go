package jamfprotect

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func newMockServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/token" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(map[string]any{
				"access_token": "test-token",
				"expires_in":   3600,
				"token_type":   "Bearer",
			})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
}

func TestNewClient(t *testing.T) {
	srv := newMockServer(t)
	defer srv.Close()

	client, err := NewClient("test-client", "test-secret", WithBaseURL(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	assert.NotNil(t, client.ActionConfig)
	assert.NotNil(t, client.Analytic)
	assert.NotNil(t, client.AnalyticSet)
	assert.NotNil(t, client.ApiClient)
	assert.NotNil(t, client.ChangeManagement)
	assert.NotNil(t, client.Computer)
	assert.NotNil(t, client.DataForwarding)
	assert.NotNil(t, client.DataRetention)
	assert.NotNil(t, client.Downloads)
	assert.NotNil(t, client.ExceptionSet)
	assert.NotNil(t, client.Group)
	assert.NotNil(t, client.IdentityProvider)
	assert.NotNil(t, client.Plan)
	assert.NotNil(t, client.PreventList)
	assert.NotNil(t, client.Role)
	assert.NotNil(t, client.TelemetryV2)
	assert.NotNil(t, client.UnifiedLoggingFilter)
	assert.NotNil(t, client.USBControlSet)
	assert.NotNil(t, client.User)
}

func TestNewClient_ValidationError(t *testing.T) {
	_, err := NewClient("", "secret")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "client ID")

	_, err = NewClient("client", "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "client secret")
}

func TestNewClientFromEnv(t *testing.T) {
	srv := newMockServer(t)
	defer srv.Close()

	os.Setenv("JAMFPROTECT_CLIENT_ID", "test-client")
	os.Setenv("JAMFPROTECT_CLIENT_SECRET", "test-secret")
	os.Setenv("JAMFPROTECT_BASE_URL", srv.URL)
	defer func() {
		os.Unsetenv("JAMFPROTECT_CLIENT_ID")
		os.Unsetenv("JAMFPROTECT_CLIENT_SECRET")
		os.Unsetenv("JAMFPROTECT_BASE_URL")
	}()

	client, err := NewClientFromEnv()
	require.NoError(t, err)
	require.NotNil(t, client)
}

func TestNewClientFromEnv_MissingClientID(t *testing.T) {
	os.Unsetenv("JAMFPROTECT_CLIENT_ID")
	os.Setenv("JAMFPROTECT_CLIENT_SECRET", "secret")
	defer os.Unsetenv("JAMFPROTECT_CLIENT_SECRET")

	_, err := NewClientFromEnv()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "JAMFPROTECT_CLIENT_ID")
}

func TestNewClientFromEnv_MissingClientSecret(t *testing.T) {
	os.Setenv("JAMFPROTECT_CLIENT_ID", "client")
	os.Unsetenv("JAMFPROTECT_CLIENT_SECRET")
	defer os.Unsetenv("JAMFPROTECT_CLIENT_ID")

	_, err := NewClientFromEnv()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "JAMFPROTECT_CLIENT_SECRET")
}

func TestClient_GetLogger(t *testing.T) {
	srv := newMockServer(t)
	defer srv.Close()

	logger := zap.NewNop()
	client, err := NewClient("test-client", "test-secret", WithBaseURL(srv.URL), WithLogger(logger))
	require.NoError(t, err)
	assert.Same(t, logger, client.GetLogger())
}

func TestClient_GetTransport(t *testing.T) {
	srv := newMockServer(t)
	defer srv.Close()

	client, err := NewClient("test-client", "test-secret", WithBaseURL(srv.URL))
	require.NoError(t, err)
	assert.NotNil(t, client.GetTransport())
}

func TestClient_GetTokenManager(t *testing.T) {
	srv := newMockServer(t)
	defer srv.Close()

	client, err := NewClient("test-client", "test-secret", WithBaseURL(srv.URL))
	require.NoError(t, err)
	assert.NotNil(t, client.GetTokenManager())
}

func TestClient_RefreshToken(t *testing.T) {
	srv := newMockServer(t)
	defer srv.Close()

	client, err := NewClient("test-client", "test-secret", WithBaseURL(srv.URL))
	require.NoError(t, err)

	err = client.RefreshToken(context.Background())
	require.NoError(t, err)
}

func TestClient_InvalidateToken(t *testing.T) {
	srv := newMockServer(t)
	defer srv.Close()

	client, err := NewClient("test-client", "test-secret", WithBaseURL(srv.URL))
	require.NoError(t, err)

	client.InvalidateToken()
}
