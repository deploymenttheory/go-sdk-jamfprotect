package client

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func newMockAuthServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/token":
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(map[string]any{
				"access_token": "test-token",
				"expires_in":   3600,
				"token_type":   "Bearer",
			})
		case "/app":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(map[string]any{
				"data": map[string]any{"result": "success"},
			})
		default:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{}`))
		}
	}))
}

func TestNewTransport(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		return nil
	})
	require.NoError(t, err)
	require.NotNil(t, tr)
	assert.Equal(t, srv.URL, tr.baseURL)
	assert.NotNil(t, tr.GetHTTPClient())
	assert.NotNil(t, tr.GetLogger())
	assert.NotNil(t, tr.GetTokenManager())
}

func TestNewTransport_ValidationErrors(t *testing.T) {
	tests := []struct {
		name         string
		clientID     string
		clientSecret string
		wantErr      string
	}{
		{"empty client ID", "", "secret", "client ID cannot be empty"},
		{"empty client secret", "client", "", "client secret cannot be empty"},
		{"both empty", "", "", "client ID cannot be empty"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewTransport(tt.clientID, tt.clientSecret)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}

func TestNewTransport_CustomLogger(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	logger := zap.NewNop()
	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		s.Logger = logger
		return nil
	})
	require.NoError(t, err)
	assert.Same(t, logger, tr.GetLogger())
}

func TestNewTransport_OptionError(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	_, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		return nil
	}, func(s *TransportSettings) error {
		return fmt.Errorf("option error")
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "option error")
}

func TestNewTransport_Timeout(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	_, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		s.Timeout = 30 * time.Second
		return nil
	})
	require.NoError(t, err)
}

func TestNewTransport_RetrySettings(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	_, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		s.RetryCount = 2
		s.RetryWaitTime = time.Second
		s.RetryMaxWaitTime = 10 * time.Second
		return nil
	})
	require.NoError(t, err)
}

func TestNewTransport_UserAgent_GlobalHeaders(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		s.UserAgent = "custom/1.0"
		s.GlobalHeaders = map[string]string{"X-Custom": "value"}
		return nil
	})
	require.NoError(t, err)
	assert.Equal(t, "custom/1.0", tr.userAgent)
	assert.Equal(t, "value", tr.globalHeaders["X-Custom"])
}

func TestNewTransport_TLSAndProxyOptions(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	_, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		s.TLSClientConfig = &tls.Config{MinVersion: tls.VersionTLS12}
		s.HTTPTransport = http.DefaultTransport
		s.InsecureSkipVerify = true
		return nil
	})
	require.NoError(t, err)
}

func TestNewTransport_ConcurrencyAndDelays(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		s.MaxConcurrentRequests = 2
		s.MandatoryRequestDelay = time.Millisecond
		s.TotalRetryDuration = 10 * time.Second
		return nil
	})
	require.NoError(t, err)
	assert.NotNil(t, tr.sem)
	assert.Greater(t, tr.requestDelay, time.Duration(0))
	assert.Greater(t, tr.totalRetryDuration, time.Duration(0))
}

func TestNewTransport_MaxConcurrentRequests_Zero(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		s.MaxConcurrentRequests = 0
		return nil
	})
	require.NoError(t, err)
	assert.Nil(t, tr.sem)
}

func TestNewTransport_Debug(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	_, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		s.Debug = true
		return nil
	})
	require.NoError(t, err)
}

func TestTransport_NewRequest(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		return nil
	})
	require.NoError(t, err)

	ctx := context.Background()
	builder := tr.NewRequest(ctx)
	assert.NotNil(t, builder)
	assert.Equal(t, ctx, builder.ctx)
	assert.NotNil(t, builder.headers)
}

func TestTransport_ExecuteGraphQL(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		return nil
	})
	require.NoError(t, err)

	var result map[string]any
	resp, err := tr.NewRequest(context.Background()).
		SetQuery("query { test }").
		SetVariables(map[string]any{"id": "123"}).
		SetTarget(&result).
		Post("/app")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "success", result["result"])
}

func TestTransport_ExecuteGraphQL_PathNormalization(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		return nil
	})
	require.NoError(t, err)

	resp, err := tr.NewRequest(context.Background()).
		SetQuery("query { test }").
		Post("app")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestTransport_ExecuteGraphQL_EmptyPath(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		return nil
	})
	require.NoError(t, err)

	_, err = tr.executeGraphQL(context.Background(), "", "query { test }", nil, nil, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "path is required")
}

func TestTransport_ExecuteGraphQL_GraphQLErrors(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		if r.URL.Path == "/app" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(map[string]any{
				"errors": []map[string]any{
					{"message": "GraphQL error occurred"},
				},
			})
			return
		}
	}))
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		return nil
	})
	require.NoError(t, err)

	_, err = tr.executeGraphQL(context.Background(), "/app", "query { test }", nil, nil, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "GraphQL error occurred")
}

func TestTransport_ExecuteGraphQL_NilTarget(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		return nil
	})
	require.NoError(t, err)

	resp, err := tr.executeGraphQL(context.Background(), "/app", "query { test }", nil, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestTransport_ExecuteRequest_ContextCanceled(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		return nil
	})
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err = tr.Post(ctx, "/app", nil, nil, nil)
	require.Error(t, err)
}

func TestTransport_ExecuteRequest_ConcurrencyLimit(t *testing.T) {
	block := make(chan struct{})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		if r.URL.Path == "/block" {
			<-block
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer srv.Close()
	close(block)

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		s.MaxConcurrentRequests = 1
		return nil
	})
	require.NoError(t, err)

	block = make(chan struct{})
	ctx := context.Background()
	go func() { _, _ = tr.Post(ctx, "/block", nil, nil, nil) }()
	time.Sleep(50 * time.Millisecond)

	ctx2, cancel2 := context.WithCancel(context.Background())
	cancel2()
	_, err = tr.Post(ctx2, "/test", nil, nil, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "concurrency limit")
	close(block)
}

func TestTransport_ValidateResponse_EmptyBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		if r.URL.Path == "/empty" {
			w.Header().Set("Content-Length", "0")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		return nil
	})
	require.NoError(t, err)

	resp, err := tr.Post(context.Background(), "/empty", nil, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestTransport_ValidateResponse_WrongContentType(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		if r.URL.Path == "/plain" {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("<html>Not JSON</html>"))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		return nil
	})
	require.NoError(t, err)

	_, err = tr.Post(context.Background(), "/plain", nil, nil, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected response Content-Type")
}

func TestTransport_WithTotalRetryDuration(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		s.TotalRetryDuration = 30 * time.Second
		return nil
	})
	require.NoError(t, err)

	resp, err := tr.Post(context.Background(), "/app", nil, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestTransport_RequestWithMandatoryDelay(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		s.MandatoryRequestDelay = 50 * time.Millisecond
		return nil
	})
	require.NoError(t, err)

	start := time.Now()
	resp, err := tr.Post(context.Background(), "/app", nil, nil, nil)
	elapsed := time.Since(start)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.GreaterOrEqual(t, elapsed, 50*time.Millisecond)
}

func TestTransport_GlobalHeaders(t *testing.T) {
	var receivedHeaders http.Header
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		receivedHeaders = r.Header.Clone()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		s.GlobalHeaders = map[string]string{"X-Global": "global-val"}
		return nil
	})
	require.NoError(t, err)

	resp, err := tr.Post(context.Background(), "/test", nil, map[string]string{"X-Request": "request-val"}, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "global-val", receivedHeaders.Get("X-Global"))
	assert.Equal(t, "request-val", receivedHeaders.Get("X-Request"))
}

func TestTransport_SetLogger(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		return nil
	})
	require.NoError(t, err)

	newLogger := zap.NewNop()
	tr.SetLogger(newLogger)
	assert.Same(t, newLogger, tr.GetLogger())
	assert.Same(t, newLogger, tr.tokenManager.logger)
}

func TestTransport_SetLogger_Nil(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		return nil
	})
	require.NoError(t, err)

	originalLogger := tr.logger
	tr.SetLogger(nil)
	assert.Same(t, originalLogger, tr.logger)
}

func TestTransport_AccessToken(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		return nil
	})
	require.NoError(t, err)

	token, err := tr.AccessToken(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "test-token", token)
}

func TestTransport_RefreshToken(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		return nil
	})
	require.NoError(t, err)

	err = tr.RefreshToken(context.Background())
	require.NoError(t, err)
}

func TestTransport_InvalidateToken(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		return nil
	})
	require.NoError(t, err)

	_, err = tr.AccessToken(context.Background())
	require.NoError(t, err)

	tr.InvalidateToken()

	token, err := tr.AccessToken(context.Background())
	require.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestTransport_ExecuteRequest_ServerError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error":{"message":"internal error"}}`))
	}))
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		return nil
	})
	require.NoError(t, err)

	resp, err := tr.Post(context.Background(), "/error", nil, nil, nil)
	require.Error(t, err)
	assert.Equal(t, 500, resp.StatusCode())
}
