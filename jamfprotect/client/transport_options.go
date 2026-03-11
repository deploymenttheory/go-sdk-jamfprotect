package client

import (
	"crypto/tls"
	"fmt"
	"maps"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// WithBaseURL sets a custom base URL for the API client.
func WithBaseURL(baseURL string) ClientOption {
	return func(s *TransportSettings) error {
		s.BaseURL = baseURL
		return nil
	}
}

// WithTimeout sets a custom timeout for HTTP requests.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(s *TransportSettings) error {
		s.Timeout = timeout
		return nil
	}
}

// WithRetryCount sets the number of retries for failed requests.
func WithRetryCount(count int) ClientOption {
	return func(s *TransportSettings) error {
		s.RetryCount = count
		return nil
	}
}

// WithRetryWaitTime sets the default wait time between retry attempts.
// This is the initial/minimum wait time before the first retry.
func WithRetryWaitTime(waitTime time.Duration) ClientOption {
	return func(s *TransportSettings) error {
		s.RetryWaitTime = waitTime
		return nil
	}
}

// WithRetryMaxWaitTime sets the maximum wait time between retry attempts.
// The wait time increases exponentially with each retry up to this maximum.
func WithRetryMaxWaitTime(maxWaitTime time.Duration) ClientOption {
	return func(s *TransportSettings) error {
		s.RetryMaxWaitTime = maxWaitTime
		return nil
	}
}

// WithLogger sets a custom logger for the client.
func WithLogger(logger *zap.Logger) ClientOption {
	return func(s *TransportSettings) error {
		if logger == nil {
			return fmt.Errorf("logger cannot be nil")
		}
		s.Logger = logger
		return nil
	}
}

// WithDebug enables debug mode which logs request and response details.
func WithDebug() ClientOption {
	return func(s *TransportSettings) error {
		s.Debug = true
		return nil
	}
}

// WithUserAgent sets a custom user agent string.
func WithUserAgent(userAgent string) ClientOption {
	return func(s *TransportSettings) error {
		s.UserAgent = userAgent
		return nil
	}
}

// WithGlobalHeader sets a global header that will be included in all requests.
// Per-request headers will override global headers with the same key.
func WithGlobalHeader(key, value string) ClientOption {
	return func(s *TransportSettings) error {
		if s.GlobalHeaders == nil {
			s.GlobalHeaders = make(map[string]string)
		}
		s.GlobalHeaders[key] = value
		return nil
	}
}

// WithGlobalHeaders sets multiple global headers at once.
func WithGlobalHeaders(headers map[string]string) ClientOption {
	return func(s *TransportSettings) error {
		if s.GlobalHeaders == nil {
			s.GlobalHeaders = make(map[string]string)
		}
		maps.Copy(s.GlobalHeaders, headers)
		return nil
	}
}

// WithProxy sets an HTTP proxy for all requests.
// Example: "http://proxy.company.com:8080" or "socks5://127.0.0.1:1080"
func WithProxy(proxyURL string) ClientOption {
	return func(s *TransportSettings) error {
		s.ProxyURL = proxyURL
		return nil
	}
}

// WithTLSClientConfig sets a custom TLS configuration.
func WithTLSClientConfig(tlsConfig *tls.Config) ClientOption {
	return func(s *TransportSettings) error {
		s.TLSClientConfig = tlsConfig
		return nil
	}
}

// WithTransport sets a custom HTTP transport.
func WithTransport(transport http.RoundTripper) ClientOption {
	return func(s *TransportSettings) error {
		s.HTTPTransport = transport
		return nil
	}
}

// WithInsecureSkipVerify disables SSL certificate verification.
// WARNING: Only use this for testing. Never in production!
func WithInsecureSkipVerify() ClientOption {
	return func(s *TransportSettings) error {
		s.InsecureSkipVerify = true
		return nil
	}
}

// WithMaxConcurrentRequests limits the number of in-flight requests.
// A buffered channel semaphore of size n is used. Zero means unlimited.
func WithMaxConcurrentRequests(n int) ClientOption {
	return func(s *TransportSettings) error {
		s.MaxConcurrentRequests = n
		return nil
	}
}

// WithMandatoryRequestDelay sets a fixed delay applied after every successful request.
func WithMandatoryRequestDelay(d time.Duration) ClientOption {
	return func(s *TransportSettings) error {
		s.MandatoryRequestDelay = d
		return nil
	}
}

// WithTotalRetryDuration sets an overall deadline applied to the request context
// when no deadline is already set. Requests exceeding this duration will be cancelled.
func WithTotalRetryDuration(d time.Duration) ClientOption {
	return func(s *TransportSettings) error {
		s.TotalRetryDuration = d
		return nil
	}
}
