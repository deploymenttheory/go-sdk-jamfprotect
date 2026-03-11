package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
)

func TestDefaultOTelConfig(t *testing.T) {
	config := DefaultOTelConfig()
	assert.NotNil(t, config)
	assert.Equal(t, "jamfprotect-client", config.ServiceName)
	assert.NotNil(t, config.TracerProvider)
	assert.NotNil(t, config.Propagators)
}

func TestTransport_EnableTracing(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		return nil
	})
	require.NoError(t, err)

	err = tr.EnableTracing(nil)
	assert.NoError(t, err)
}

func TestTransport_EnableTracing_CustomConfig(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	tr, err := NewTransport("test-client", "test-secret", func(s *TransportSettings) error {
		s.BaseURL = srv.URL
		return nil
	})
	require.NoError(t, err)

	config := &OTelConfig{
		TracerProvider: otel.GetTracerProvider(),
		Propagators:    otel.GetTextMapPropagator(),
		ServiceName:    "custom-service",
		SpanNameFormatter: func(operation string, req *http.Request) string {
			return "custom-span"
		},
	}

	err = tr.EnableTracing(config)
	assert.NoError(t, err)
}
