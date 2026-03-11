package jamfprotect

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"go.uber.org/zap"
)

// ClientOption configures the Jamf Protect API transport at construction time.
// It is an alias for client.ClientOption so callers can use jamfprotect.ClientOption
// without importing the client sub-package.
type ClientOption = client.ClientOption

// WithBaseURL sets a custom base URL, overriding the default Jamf Protect cloud URL.
func WithBaseURL(baseURL string) ClientOption {
	return client.WithBaseURL(baseURL)
}

// WithTimeout sets a custom timeout for HTTP requests.
func WithTimeout(timeout time.Duration) ClientOption {
	return client.WithTimeout(timeout)
}

// WithRetryCount sets the number of retries for failed requests.
func WithRetryCount(count int) ClientOption {
	return client.WithRetryCount(count)
}

// WithRetryWaitTime sets the initial wait time between retry attempts.
func WithRetryWaitTime(waitTime time.Duration) ClientOption {
	return client.WithRetryWaitTime(waitTime)
}

// WithRetryMaxWaitTime sets the maximum wait time between retry attempts.
func WithRetryMaxWaitTime(maxWaitTime time.Duration) ClientOption {
	return client.WithRetryMaxWaitTime(maxWaitTime)
}

// WithLogger sets a custom zap logger for the client.
func WithLogger(logger *zap.Logger) ClientOption {
	return client.WithLogger(logger)
}

// WithDebug enables debug mode which logs request and response details.
func WithDebug() ClientOption {
	return client.WithDebug()
}

// WithUserAgent sets a custom user agent string.
func WithUserAgent(userAgent string) ClientOption {
	return client.WithUserAgent(userAgent)
}

// WithGlobalHeader sets a global header that will be included in all requests.
// Per-request headers will override global headers with the same key.
func WithGlobalHeader(key, value string) ClientOption {
	return client.WithGlobalHeader(key, value)
}

// WithGlobalHeaders sets multiple global headers at once.
func WithGlobalHeaders(headers map[string]string) ClientOption {
	return client.WithGlobalHeaders(headers)
}

// WithProxy sets an HTTP proxy for all requests.
// Example: "http://proxy.company.com:8080" or "socks5://127.0.0.1:1080"
func WithProxy(proxyURL string) ClientOption {
	return client.WithProxy(proxyURL)
}

// WithTLSClientConfig sets a custom TLS configuration.
func WithTLSClientConfig(tlsConfig *tls.Config) ClientOption {
	return client.WithTLSClientConfig(tlsConfig)
}

// WithTransport sets a custom HTTP transport (http.RoundTripper).
func WithTransport(transport http.RoundTripper) ClientOption {
	return client.WithTransport(transport)
}

// WithInsecureSkipVerify disables SSL certificate verification.
// WARNING: Only use this for testing. Never in production!
func WithInsecureSkipVerify() ClientOption {
	return client.WithInsecureSkipVerify()
}

// WithMaxConcurrentRequests limits the number of in-flight requests.
// A buffered channel semaphore of size n is used. Zero means unlimited.
func WithMaxConcurrentRequests(n int) ClientOption {
	return client.WithMaxConcurrentRequests(n)
}

// WithMandatoryRequestDelay sets a fixed delay applied after every successful request.
func WithMandatoryRequestDelay(d time.Duration) ClientOption {
	return client.WithMandatoryRequestDelay(d)
}

// WithTotalRetryDuration sets an overall deadline applied to the request context
// when no deadline is already set. Requests exceeding this duration will be cancelled.
func WithTotalRetryDuration(d time.Duration) ClientOption {
	return client.WithTotalRetryDuration(d)
}
