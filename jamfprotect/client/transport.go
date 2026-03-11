package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"resty.dev/v3"
)

// Ensure Transport implements GraphQLClient at compile time.
var _ GraphQLClient = (*Transport)(nil)

// Transport is the HTTP transport layer for the Jamf Protect GraphQL API.
// It provides authentication, retry, and logging. Request execution is in request.go.
type Transport struct {
	client             *resty.Client
	baseURL            string
	userAgent          string
	logger             *zap.Logger
	authConfig         *AuthConfig
	tokenManager       *TokenManager
	globalHeaders      map[string]string
	sem                chan struct{}
	requestDelay       time.Duration
	totalRetryDuration time.Duration
}

// NewTransport creates a new Jamf Protect GraphQL transport.
func NewTransport(clientID, clientSecret string, options ...ClientOption) (*Transport, error) {
	if err := ValidateTransportConfig(clientID, clientSecret); err != nil {
		return nil, fmt.Errorf("invalid transport configuration: %w", err)
	}

	settings := &TransportSettings{
		GlobalHeaders: make(map[string]string),
	}
	for _, opt := range options {
		if err := opt(settings); err != nil {
			return nil, fmt.Errorf("applying client option: %w", err)
		}
	}

	// Logger
	logger := settings.Logger
	if logger == nil {
		var err error
		logger, err = zap.NewProduction()
		if err != nil {
			return nil, fmt.Errorf("failed to create logger: %w", err)
		}
	}

	// BaseURL
	baseURL := settings.BaseURL
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}

	// UserAgent
	userAgent := settings.UserAgent
	if userAgent == "" {
		userAgent = fmt.Sprintf("%s/%s", UserAgentBase, Version)
	}

	// Timeouts/retries
	timeout := settings.Timeout
	if timeout == 0 {
		timeout = time.Duration(DefaultTimeout) * time.Second
	}
	retryCount := settings.RetryCount
	if retryCount == 0 {
		retryCount = MaxRetries
	}
	retryWait := settings.RetryWaitTime
	if retryWait == 0 {
		retryWait = time.Duration(RetryWaitTime) * time.Second
	}
	retryMaxWait := settings.RetryMaxWaitTime
	if retryMaxWait == 0 {
		retryMaxWait = time.Duration(RetryMaxWaitTime) * time.Second
	}

	restyClient := resty.New()
	restyClient.SetTimeout(timeout)
	restyClient.SetRetryCount(retryCount)
	restyClient.SetRetryWaitTime(retryWait)
	restyClient.SetRetryMaxWaitTime(retryMaxWait)
	restyClient.SetHeader(HeaderUserAgent, userAgent)
	restyClient.SetHeader(HeaderContentType, ContentTypeJSON)
	restyClient.SetHeader("Accept", AcceptJSON)

	if settings.Debug {
		restyClient.SetDebug(true)
	}
	if settings.InsecureSkipVerify {
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}) //nolint:gosec
	} else if settings.TLSClientConfig != nil {
		restyClient.SetTLSClientConfig(settings.TLSClientConfig)
	}
	if settings.ProxyURL != "" {
		restyClient.SetProxy(settings.ProxyURL)
	}
	if settings.HTTPTransport != nil {
		httpClient := restyClient.Client()
		if httpClient != nil {
			httpClient.Transport = settings.HTTPTransport
		}
	}
	for k, v := range settings.GlobalHeaders {
		restyClient.SetHeader(k, v)
	}

	// Semaphore
	var sem chan struct{}
	if settings.MaxConcurrentRequests > 0 {
		sem = make(chan struct{}, settings.MaxConcurrentRequests)
	}

	transport := &Transport{
		client:             restyClient,
		logger:             logger,
		baseURL:            baseURL,
		globalHeaders:      settings.GlobalHeaders,
		userAgent:          userAgent,
		sem:                sem,
		requestDelay:       settings.MandatoryRequestDelay,
		totalRetryDuration: settings.TotalRetryDuration,
	}

	authConfig := &AuthConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     strings.TrimRight(baseURL, "/") + EndpointToken,
	}
	transport.authConfig = authConfig

	tokenManager, err := SetupAuthentication(restyClient, authConfig, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to setup authentication: %w", err)
	}
	transport.tokenManager = tokenManager

	restyClient.SetBaseURL(baseURL)

	logger.Info("Jamf Protect API client created",
		zap.String("base_url", baseURL),
		zap.String("client_id", clientID))

	return transport, nil
}

// GetHTTPClient returns the underlying resty client.
func (t *Transport) GetHTTPClient() *resty.Client {
	return t.client
}

// GetLogger returns the configured zap logger.
func (t *Transport) GetLogger() *zap.Logger {
	return t.logger
}

// GetTokenManager returns the token manager.
func (t *Transport) GetTokenManager() *TokenManager {
	return t.tokenManager
}

// SetLogger updates the logger at runtime.
func (t *Transport) SetLogger(logger *zap.Logger) {
	if logger != nil {
		t.logger = logger
		t.tokenManager.logger = logger
	}
}

// AccessToken retrieves a valid access token, refreshing if necessary.
func (t *Transport) AccessToken(ctx context.Context) (string, error) {
	return t.tokenManager.GetToken(ctx)
}

// RefreshToken manually refreshes the OAuth2 access token.
func (t *Transport) RefreshToken(ctx context.Context) error {
	_, err := t.tokenManager.RefreshToken(ctx)
	return err
}

// InvalidateToken invalidates the current token, forcing a refresh on next use.
func (t *Transport) InvalidateToken() {
	t.tokenManager.InvalidateToken()
}
