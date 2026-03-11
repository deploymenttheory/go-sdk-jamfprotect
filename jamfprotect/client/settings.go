package client

import (
	"crypto/tls"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// ClientOption is a function that mutates TransportSettings before Transport construction.
type ClientOption func(*TransportSettings) error

// TransportSettings collects all optional transport configuration.
// Zero values mean "use the built-in default". Options are applied to a
// TransportSettings before the Transport is constructed.
type TransportSettings struct {
	BaseURL               string
	Timeout               time.Duration
	RetryCount            int
	RetryWaitTime         time.Duration
	RetryMaxWaitTime      time.Duration
	Logger                *zap.Logger
	Debug                 bool
	UserAgent             string
	GlobalHeaders         map[string]string
	ProxyURL              string
	TLSClientConfig       *tls.Config
	HTTPTransport         http.RoundTripper
	InsecureSkipVerify    bool
	MaxConcurrentRequests int
	MandatoryRequestDelay time.Duration
	TotalRetryDuration    time.Duration
}
