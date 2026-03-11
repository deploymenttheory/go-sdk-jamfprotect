package jamfprotect

import (
	"crypto/tls"
	"net/http"
	"testing"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestWithBaseURL(t *testing.T) {
	settings := &client.TransportSettings{}
	err := WithBaseURL("https://custom.example.com")(settings)
	require.NoError(t, err)
	assert.Equal(t, "https://custom.example.com", settings.BaseURL)
}

func TestWithTimeout(t *testing.T) {
	settings := &client.TransportSettings{}
	err := WithTimeout(30 * time.Second)(settings)
	require.NoError(t, err)
	assert.Equal(t, 30*time.Second, settings.Timeout)
}

func TestWithRetryCount(t *testing.T) {
	settings := &client.TransportSettings{}
	err := WithRetryCount(5)(settings)
	require.NoError(t, err)
	assert.Equal(t, 5, settings.RetryCount)
}

func TestWithRetryWaitTime(t *testing.T) {
	settings := &client.TransportSettings{}
	err := WithRetryWaitTime(2 * time.Second)(settings)
	require.NoError(t, err)
	assert.Equal(t, 2*time.Second, settings.RetryWaitTime)
}

func TestWithRetryMaxWaitTime(t *testing.T) {
	settings := &client.TransportSettings{}
	err := WithRetryMaxWaitTime(30 * time.Second)(settings)
	require.NoError(t, err)
	assert.Equal(t, 30*time.Second, settings.RetryMaxWaitTime)
}

func TestWithLogger(t *testing.T) {
	logger := zap.NewNop()
	settings := &client.TransportSettings{}
	err := WithLogger(logger)(settings)
	require.NoError(t, err)
	assert.Same(t, logger, settings.Logger)
}

func TestWithDebug(t *testing.T) {
	settings := &client.TransportSettings{}
	err := WithDebug()(settings)
	require.NoError(t, err)
	assert.True(t, settings.Debug)
}

func TestWithUserAgent(t *testing.T) {
	settings := &client.TransportSettings{}
	err := WithUserAgent("CustomAgent/1.0")(settings)
	require.NoError(t, err)
	assert.Equal(t, "CustomAgent/1.0", settings.UserAgent)
}

func TestWithGlobalHeader(t *testing.T) {
	settings := &client.TransportSettings{}
	err := WithGlobalHeader("X-Custom", "value")(settings)
	require.NoError(t, err)
	assert.Equal(t, "value", settings.GlobalHeaders["X-Custom"])
}

func TestWithGlobalHeaders(t *testing.T) {
	settings := &client.TransportSettings{}
	headers := map[string]string{"X-A": "a", "X-B": "b"}
	err := WithGlobalHeaders(headers)(settings)
	require.NoError(t, err)
	assert.Equal(t, "a", settings.GlobalHeaders["X-A"])
	assert.Equal(t, "b", settings.GlobalHeaders["X-B"])
}

func TestWithProxy(t *testing.T) {
	settings := &client.TransportSettings{}
	err := WithProxy("http://proxy.example.com:8080")(settings)
	require.NoError(t, err)
	assert.Equal(t, "http://proxy.example.com:8080", settings.ProxyURL)
}

func TestWithTLSClientConfig(t *testing.T) {
	settings := &client.TransportSettings{}
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}
	err := WithTLSClientConfig(tlsConfig)(settings)
	require.NoError(t, err)
	assert.Same(t, tlsConfig, settings.TLSClientConfig)
}

func TestWithTransport(t *testing.T) {
	settings := &client.TransportSettings{}
	customTransport := http.DefaultTransport
	err := WithTransport(customTransport)(settings)
	require.NoError(t, err)
	assert.Same(t, customTransport, settings.HTTPTransport)
}

func TestWithInsecureSkipVerify(t *testing.T) {
	settings := &client.TransportSettings{}
	err := WithInsecureSkipVerify()(settings)
	require.NoError(t, err)
	assert.True(t, settings.InsecureSkipVerify)
}

func TestWithMaxConcurrentRequests(t *testing.T) {
	settings := &client.TransportSettings{}
	err := WithMaxConcurrentRequests(10)(settings)
	require.NoError(t, err)
	assert.Equal(t, 10, settings.MaxConcurrentRequests)
}

func TestWithMandatoryRequestDelay(t *testing.T) {
	settings := &client.TransportSettings{}
	err := WithMandatoryRequestDelay(100 * time.Millisecond)(settings)
	require.NoError(t, err)
	assert.Equal(t, 100*time.Millisecond, settings.MandatoryRequestDelay)
}

func TestWithTotalRetryDuration(t *testing.T) {
	settings := &client.TransportSettings{}
	err := WithTotalRetryDuration(5 * time.Minute)(settings)
	require.NoError(t, err)
	assert.Equal(t, 5*time.Minute, settings.TotalRetryDuration)
}
