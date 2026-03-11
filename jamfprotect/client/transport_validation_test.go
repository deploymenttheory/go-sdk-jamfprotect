package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateTransportConfig(t *testing.T) {
	tests := []struct {
		name         string
		clientID     string
		clientSecret string
		wantErr      bool
		errContains  string
	}{
		{"valid", "client", "secret", false, ""},
		{"empty client ID", "", "secret", true, "client ID cannot be empty"},
		{"empty client secret", "client", "", true, "client secret cannot be empty"},
		{"both empty", "", "", true, "client ID cannot be empty"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTransportConfig(tt.clientID, tt.clientSecret)
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errContains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateBaseURL(t *testing.T) {
	tests := []struct {
		name        string
		baseURL     string
		wantErr     bool
		errContains string
	}{
		{"valid https", "https://example.com", false, ""},
		{"valid http", "http://example.com", false, ""},
		{"empty", "", true, "base URL cannot be empty"},
		{"no protocol", "example.com", true, "must start with http:// or https://"},
		{"trailing slash", "https://example.com/", true, "should not end with a trailing slash"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateBaseURL(tt.baseURL)
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errContains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateTimeout(t *testing.T) {
	tests := []struct {
		name        string
		timeout     int
		wantErr     bool
		errContains string
	}{
		{"valid", 30, false, ""},
		{"zero", 0, true, "must be greater than 0"},
		{"negative", -1, true, "must be greater than 0"},
		{"too large", 3601, true, "timeout too large"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTimeout(tt.timeout)
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errContains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateRetryCount(t *testing.T) {
	tests := []struct {
		name        string
		retryCount  int
		wantErr     bool
		errContains string
	}{
		{"valid", 3, false, ""},
		{"zero", 0, false, ""},
		{"negative", -1, true, "retry count cannot be negative"},
		{"too large", 11, true, "retry count too large"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateRetryCount(tt.retryCount)
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errContains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateProxyURL(t *testing.T) {
	tests := []struct {
		name        string
		proxyURL    string
		wantErr     bool
		errContains string
	}{
		{"valid http", "http://proxy.example.com:8080", false, ""},
		{"valid https", "https://proxy.example.com:8080", false, ""},
		{"valid socks5", "socks5://127.0.0.1:1080", false, ""},
		{"empty", "", false, ""},
		{"invalid protocol", "ftp://proxy.example.com", true, "must start with http://, https://, or socks5://"},
		{"no protocol", "proxy.example.com:8080", true, "must start with http://, https://, or socks5://"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateProxyURL(tt.proxyURL)
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errContains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
